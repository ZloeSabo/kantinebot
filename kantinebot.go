package kantinebot

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

type BotProduct struct {
	Name string
	// Description() string
	// Price() string
	// KCal() string
}

type BotApiBridge interface {
	Can(ask string) bool
	Menu() *[]BotProduct
}

type Bot struct {
	apikey  string
	debug   bool
	log     *log.Logger
	bridges *[]BotApiBridge
}

type OptionB func(*Bot)

func OptionApiBridge(bridge BotApiBridge) func(*Bot) {
	return func(b *Bot) {
		*b.bridges = append(*b.bridges, bridge)
	}
}

//NewBot returns new Bot
func NewBot(apikey string, options ...OptionB) *Bot {
	b := &Bot{
		apikey:  apikey,
		debug:   false,
		log:     log.New(os.Stderr, "zloesabo/kantinebot", log.LstdFlags|log.Lshortfile),
		bridges: &[]BotApiBridge{},
	}

	for _, opt := range options {
		opt(b)
	}

	return b
}

func (bot *Bot) Run() {
	logger := bot.log

	api := slack.New(
		bot.apikey,
		slack.OptionDebug(bot.debug),
		slack.OptionLog(logger),
	)

	auth, err := api.AuthTest()
	if err != nil {
		logger.Panicln(err)
		return
	}

	logger.Output(2, fmt.Sprintf("ID: %s User: %s", auth.UserID, auth.User))

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		logger.Print("Event Received: ")
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			// Ignore hello

		case *slack.ConnectedEvent:
			logger.Println("Infos:", ev.Info)
			logger.Println("Connection counter:", ev.ConnectionCount)
			// Replace C2147483705 with your Channel ID
			// rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", "C2147483705"))

		case *slack.MessageEvent:
			logger.Printf("Message: %v\n", ev)
			// fmt.Printf("Message: %v\n", ev.)
			channelID := ev.Channel
			for _, bridge := range *bot.bridges {
				if bridge.Can(ev.Msg.Text) {
					products := bridge.Menu()
					message := toMessage(products)
					rtm.SendMessage(rtm.NewOutgoingMessage(message, channelID))

					break
				}
			}

		case *slack.PresenceChangeEvent:
			logger.Printf("Presence Change: %v\n", ev)

		case *slack.LatencyReport:
			logger.Printf("Current latency: %v\n", ev.Value)

		case *slack.RTMError:
			logger.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			logger.Printf("Invalid credentials")
			return

		default:

			// Ignore other events..
			// fmt.Printf("Unexpected: %v\n", msg.Data)
		}
	}
}

func toMessage(products *[]BotProduct) string {
	res := strings.Builder{}

	for _, product := range *products {
		res.WriteString(fmt.Sprintln(product.Name))
	}

	return res.String()
}
