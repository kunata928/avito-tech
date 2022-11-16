package messages

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/shopspring/decimal"
	"strings"
)

func (m *Model) updateLimit(ctx context.Context, msg Message) string {
	span, ctx := opentracing.StartSpanFromContext(ctx, "updateLimit")
	defer span.Finish()

	v := strings.TrimSpace(msg.Text[len("/limit"):])
	if num, err := decimal.NewFromString(v); err == nil {

		if err := m.storage.RefreshClientLimit(ctx, msg.UserID, num); err != nil {
			return internalErrMsg
		}
		return "Successfully set limit!"
	}
	return parseErrMsg
}
