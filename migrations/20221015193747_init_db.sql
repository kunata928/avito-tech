-- +goose Up
-- +goose StatementBegin

create table users
(
    id              integer primary key,
    limit_u         decimal NOT NULL default 0,
    code_currency   VARCHAR(3) default 'RUB',
    CONSTRAINT not_negative_limit CHECK (limit_u >= 0)
);

create table rates
(
    id               serial primary key,
    code_currency    VARCHAR(3),
    amount           decimal,
    date             date
);

create UNIQUE index rates_code_date_idx on rates(code_currency, date);

create table expenses
(
    id          SERIAL primary key,
    user_id     integer REFERENCES users (id),
    date        date,
    category    varchar(250),
    amount      decimal
);


-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin

drop index rates_code_date_idx;
drop table users cascade ;
drop table rates cascade ;
drop table expenses cascade ;

-- +goose StatementEnd
