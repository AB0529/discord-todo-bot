package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Commands all the commands for the bot
var Commands = make(map[string]*Command)

func main() {
	// Load config
	Config = NewConfig("../config.yml")
	// Setup Discord
	dg, _ := discordgo.New("Bot " + Config.Token)
	// Register events
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages|discordgo.IntentsGuildMessageReactions|discordgo.IntentsGuildEmojis|discordgo.IntentsDirectMessages)
	dg.AddHandler(Ready)
	dg.AddHandler(MessageCreate)
	dg.AddHandler(MessageReactionAdd)

	// Register commands
	RegisterCommands([]*Command{
		{
			Name:    "ping",
			Aliases: []string{"pong"},
			Example: []string{Config.Prefix + "ping"},
			Desc:    []string{"Generic Ping-Pong command"},
			Handler: Ping,
		},
		{
			Name:    "test",
			Aliases: []string{},
			Example: []string{Config.Prefix + "test <flag> <value>"},
			Desc:    []string{"Command used for testing"},
			Handler: Test,
			Flags: []*Flag{ { Name: "add", RequiresValue: true }, { Name: "list" } },
		},
		{
			Name:    "todo",
			Aliases: []string{"td" },
			Example: []string{Config.Prefix + "todo <flag> <value>", Config.Prefix + "todo add Do dishes at 3:14pm"},
			Desc:    []string{"Controls the user's todo list", "Flag 'add': adds an item to your list", "Flag 'list': lists your list", "Flag 'remove': removes an item to your list"},
			Handler: Todo,
			Flags: []*Flag{ { Name: "add", RequiresValue: true }, { Name: "list" }, { Name: "li" }, { Name: "remove" }, { Name: "rm" } },
		},
	})

	// Open a websocket connection to Discord and begin listening.
	err := dg.Open()
	if err != nil {
		Die("could not creating session")
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	err = dg.Close()
	Die(err)
}
