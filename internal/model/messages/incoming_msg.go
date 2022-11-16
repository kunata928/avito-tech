package messages

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"gitlab.ozon.dev/kunata928/telegramBot/internal/metrics"
	"time"
)

var startMsg = "Hi! I respond to collect expenses:\n/add amount category date\n" +
	"/report week/month/year\n" +
	"and change /currency USD/CNY/EUR/RUB\n" +
	"Ex: </add 146.5 transport 28/09/2022>\n" +
	"Ex: </report month>\n" +
	"Ex: </currency USD>\n" +
	"Ex: </limit 120000>\n" +
	"Notice: Report will show you amount of expenses by category. Currency will change your default RUB on chosen." +
	"Limit will add set or refresh your limit per calendar month\n" +
	"Try these!\n"

var defaultMsg = "I don't know this command.\n" +
	"Type /help to see more information"

var helpMsg = "Ex: </add 146 transport 28/09/2022>\n" +
	"Ex: </report month>\n" +
	"Ex: </currency USD>\n" +
	"Ex: </limit 120000>"

var parseErrMsg = "I couldn't understand your message :(\nTry again or /help"

var internalErrMsg = "Some internal error occurred :("

func (m *Model) IncomingMessage(ctx context.Context, msg Message) error {
	// ----- metrics
	startTime := time.Now()
	defer func() {
		duration := time.Since(startTime)
		metrics.SummaryResponseTime.Observe(duration.Seconds())

		metrics.HistogramResponseTime.
			WithLabelValues(defineCommand(msg.Command)).Observe(duration.Seconds())

	}()

	// ----- span
	span, ctx := opentracing.StartSpanFromContext(
		ctx,
		defineCommand(msg.Command), //ext.RPCServerOption(incomingTrace),
	)
	defer span.Finish()

	textMsg := ""
	switch msg.Command {
	case "start":
		metrics.CommandStartTotal.Inc()
		m.storage.InitClient(ctx, msg.UserID)
		textMsg = startMsg
	case "help":
		metrics.CommandHelpTotal.Inc()
		textMsg = helpMsg
	case "add":
		metrics.CommandAddTotal.Inc()
		textMsg = m.addExpense(ctx, msg)
	case "report":
		metrics.CommandReportTotal.Inc()
		textMsg = m.reportExpenses(ctx, msg)
	case "currency":
		metrics.CommandCurrencyTotal.Inc()
		textMsg = m.changeCurrency(ctx, msg)
	case "limit":
		metrics.CommandLimitTotal.Inc()
		textMsg = m.updateLimit(ctx, msg)
	default:
		metrics.CommandDefaultTotal.Inc()
		textMsg = defaultMsg
	}
	return m.client.SendMessage(textMsg, msg.UserID)

}

func defineCommand(str string) string {
	commands := map[string]struct{}{
		"start":    struct{}{},
		"help":     struct{}{},
		"limit":    struct{}{},
		"add":      struct{}{},
		"report":   struct{}{},
		"currency": struct{}{},
	}

	if _, ok := commands[str]; ok {
		return str
	}
	return "default"
}
