package main

import (
	"os"

	"github.com/zloesabo/kantinebot"
	"github.com/zloesabo/kantinebot/apibridge"
	"github.com/zloesabo/kantinebot/wandel"
)

func main() {
	client := wandel.NewClient(
		wandel.OptionDebug(false),
		wandel.OptionAuthorization(os.Getenv("AUTHORIZATION")),
	)

	bridge := apibridge.NewWandelBridge(client)

	bot := kantinebot.NewBot(
		os.Getenv("SLACK_KEY"),
		kantinebot.OptionApiBridge(bridge),
	)
	bot.Run()
}
