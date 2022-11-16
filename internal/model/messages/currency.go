package messages

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"strings"
)

var parseErrCurr = "Couldn't understand currency.\nTry these: /currency usd"

func (m *Model) changeCurrency(ctx context.Context, msg Message) string {
	span, ctx := opentracing.StartSpanFromContext(ctx, "changeCurrency")
	defer span.Finish()

	curr := strings.ToUpper(strings.TrimSpace(msg.Text[len("/currency "):]))
	switch curr {
	case "RUB", "USD", "EUR", "CNY":
		if err := m.storage.RefreshClientCurrency(ctx, msg.UserID, curr); err == nil {
			return fmt.Sprintf("Successfuly change currency to %s!", curr)
		}
	default:
		Logger.Debug("Client msg parse currency err")
		return parseErrCurr
	}
	return internalErrMsg
}
