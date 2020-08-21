package main

import (
	"fmt"
	hbot "github.com/whyrusleeping/hellabot"
	log "gopkg.in/inconshreveable/log15.v2"
	"os"
	"regexp"
)

func main() {
	// TODO make it configurable
	r, _ := regexp.Compile("hey")
	var heyGuysTrigger = hbot.Trigger{
		Condition: func(bot *hbot.Bot, m *hbot.Message) bool {
			return m.Command == "PRIVMSG" && r.MatchString(m.Content)
		},
		Action: func(irc *hbot.Bot, m *hbot.Message) bool {
			// TODO make it configurable
			irc.Reply(m, "Hello")
			return false
		},
	}
	nick, ok := os.LookupEnv("NICK")
	if !ok {
		nick = "fawkes"
	}
	server, ok := os.LookupEnv("SERVER")
	if !ok {
		server = "chat.freenode.net:6667"
	}

	// TODO complete the port

	options := func(bot *hbot.Bot) {
		// TODO guess SSL based on the port
		bot.SSL = false
		//		bot.SASL = true
		//ot.Password = *password
	}
	// TODO handle CHANNELS
	// TODO handle password / SASL
	channels := func(bot *hbot.Bot) {
		bot.Channels = []string{"#test-misc-bot"}
	}

	irc, err := hbot.NewBot(server, nick, options, channels)
	if err != nil {
		panic(err)
	}
	irc.Logger.SetHandler(log.StdoutHandler)

	irc.AddTrigger(heyGuysTrigger)
	irc.Run()
	fmt.Println("Bot shutting down.")

}
