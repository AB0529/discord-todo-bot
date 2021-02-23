package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var LastMessage = make(chan *discordgo.Message)
var LastReaction = make(chan *discordgo.MessageReaction)

// MessageCreate the function which handles message events
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore bots
	if m.Author.Bot {
		return
	}
	// Ignore no content messages
	if m.Content == "" {
		return
	}

	// Create new context on each message1
	msg := strings.Split(strings.ToLower(m.Message.Content)[1:], " ")

	if len(msg) <= 0 {
		// Send message to LastMessage channel
		go func() { LastMessage <- m.Message }()
		return
	}

	c := msg[0]
	// Find the command with the matching name alias and run it
	cmd, ok := Commands[c]
	if !ok {
		go func() { LastMessage <- m.Message }()
		return
	}
	// Make sure message starts with prefix
	if string(m.Message.Content[0]) != Config.Prefix {
		return
	}
	ctx := &Context{
		Session: s,
		Msg:     m,
		Command: cmd,
		LastMessage: LastMessage,
		LastReaction: LastReaction,
	}
	cmd.Handler(ctx)
}

// MessageReactionAdd the function which handles message reaction events
func MessageReactionAdd(_ *discordgo.Session, m *discordgo.MessageReactionAdd) {
	go func() { LastReaction <- m.MessageReaction }()
}

// Ready the function which handles when the bot is ready
func Ready(s *discordgo.Session, e *discordgo.Ready) {
	fmt.Printf("[%s] - in as %s%s with prefix: \"%s\"\n", Purple.Sprint("BOT"), Yellow.Sprint(e.User.Username), Yellow.Sprint("#"+e.User.Discriminator), Green.Sprint(Config.Prefix))
	err := s.UpdateGameStatus(0, "with yo daddy")
	Die(err)
}

