package database

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"gitlab.ozon.dev/kunata928/telegramBot/internal/logger"
	"go.uber.org/zap"
	"log"
	"time"
)

var ErrNotFound = errors.New("Not found")
var ErrReachLimit = errors.New("Limit was reached")
var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type DB interface {
	AddExpense(ctx context.Context, userID int64, newExpense *Expense) error
	GetClientExpenses(ctx context.Context, userID int64, fromDate time.Time) ([]*Expense, error)
	AddRate(ctx context.Context, rates *Rates, date time.Time) error
	GetRate(ctx context.Context, name string, date time.Time) (decimal.Decimal, error)
	RefreshClientCurrency(ctx context.Context, userID int64, currency string) error
	GetClientCurrency(ctx context.Context, userID int64) (string, error)
	InitClient(ctx context.Context, userID int64)
	RefreshClientLimit(ctx context.Context, userID int64, amount decimal.Decimal) error
}

type Storage struct {
	db *sql.DB
	//db *pgx.Conn
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db}
}

func getDate() string {
	return time.Now().AddDate(0, 0, -time.Now().Day()).Format("2006-01-02")
}

func (s *Storage) AddExpense(ctx context.Context, userID int64, newExpense *Expense) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "AddExpense")
	defer span.Finish()

	var summ decimal.Decimal
	var limit decimal.Decimal
	const querySum = `
		select COALESCE(sum(amount), 0) from expenses where date > $1 and user_id = $2;
	`
	log.Println(getDate())
	if err := s.db.QueryRowContext(ctx, querySum, getDate(), userID).Scan(&summ); err != nil {
		logger.Error("Query err", zap.Error(err))
		return err
	}
	if err := s.db.QueryRowContext(ctx, `select limit_u from users where id=$1;`, userID).Scan(&limit); err != nil {
		logger.Error("Query err", zap.Error(err))
		return err
	}
	if limit.GreaterThan(summ.Add(newExpense.Amount)) || limit.Equal(decimal.NewFromInt(0)) {
		const query = `
		insert into expenses (user_id, date, category, amount)
					values ($1, $2, $3, $4);
		`
		_, err := s.db.ExecContext(ctx, query, userID, newExpense.Date, newExpense.Category, newExpense.Amount)
		if err != nil {
			logger.Error("Exec DB err", zap.Error(err))
			return err
		}
		return nil
	}
	return ErrReachLimit
}

func (s *Storage) GetClientExpenses(ctx context.Context, userID int64, fromDate time.Time) ([]*Expense, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetClientExpenses")
	defer span.Finish()

	const query = `
			select amount, category, date 
			from expenses
			where user_id=$1 AND date>$2;
	`
	rows, err := s.db.QueryContext(ctx, query, userID, fromDate)
	//defer rows.Close()
	if err != nil {
		logger.Error("Query err", zap.Error(err))
		return nil, err
	}

	res := make([]*Expense, 0, 1)
	for rows.Next() {
		var exp Expense
		if err := rows.Scan(&exp.Amount, &exp.Category, &exp.Date); err != nil {
			logger.Error("Scan response err", zap.Error(err))
			return res, err
		}
		res = append(res, &exp)
	}
	return res, nil
}

func (s *Storage) SetRate(ctx context.Context, rates *Rates, date time.Time) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "SetRate")
	defer span.Finish()

	q := psql.
		Insert("rates").
		Columns("code_currency", "amount", "date").
		Suffix("ON CONFLICT (code_currency, date) DO UPDATE SET amount = EXCLUDED.amount")

	for k, v := range *rates {
		q = q.Values(k, v, date.Format("2006-01-02"))
	}

	sql, args, err := q.ToSql()
	if err != nil {
		logger.Error("SQL convert err", zap.Error(err))
		return err
	}

	if _, err := s.db.ExecContext(ctx, sql, args...); err != nil {
		logger.Error("Exec err", zap.Error(err))
		return err
	}

	return nil
}

func (s *Storage) GetRate(ctx context.Context, name string, date time.Time) (decimal.Decimal, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetRate")
	defer span.Finish()

	q := psql.Select("amount").
		From("rates").
		Where("code_currency=? AND date=?", name, date.Format("2006-01-02"))

	sql, args, err := q.ToSql()
	if err != nil {
		logger.Error("SQL convert err", zap.Error(err))
		return decimal.NewFromInt(0), err
	}

	var rate decimal.Decimal
	if err := s.db.QueryRowContext(ctx, sql, args...).Scan(&rate); err != nil {
		logger.Error("Query scan err", zap.Error(err))
		return decimal.NewFromInt(0), ErrNotFound
	}
	return rate, nil
}

func (s *Storage) InitClient(ctx context.Context, userID int64) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InitClient")
	defer span.Finish()

	const query = `
		insert into users (id)
					values ($1);
	`
	_, err := s.db.ExecContext(ctx, query, userID)
	if err != nil {
		logger.Error("Query err", zap.Error(err))
	}
}

func (s *Storage) RefreshClientCurrency(ctx context.Context, userID int64, currency string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RefreshClientCurrency")
	defer span.Finish()

	const query = `
		update users
		set code_currency=$1
		where id = $2
	`
	_, err := s.db.ExecContext(ctx, query, currency, userID)
	if err != nil {
		logger.Error("Exec err", zap.Error(err))
		return err
	}
	return nil
}

func (s *Storage) GetClientCurrency(ctx context.Context, userID int64) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetClientCurrency")
	defer span.Finish()

	const query = `
			select code_currency
			from users
			where id=$1
	`
	var curr string

	err := s.db.QueryRowContext(ctx, query, userID).Scan(&curr)
	if err != nil {
		return "RUB", err
	}

	return curr, nil
}

func (s *Storage) RefreshClientLimit(ctx context.Context, userID int64, amount decimal.Decimal) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RefreshClientLimit")
	defer span.Finish()

	const query = `
		update users
		set limit_u = $1
		where id = $2
	`
	_, err := s.db.ExecContext(ctx, query, amount, userID)
	if err != nil {
		logger.Error("Exec err", zap.Error(err))
		//log.Println("err db")
		return err
	}
	return nil
}
