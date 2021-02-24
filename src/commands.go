package main

import (
	"fmt"
	"strings"
	"time"
)

// Ping command which returns a message
func Ping(ctx *Context) {
	m := ctx.NewEmbed("Pinging....")
	ts, _ := m.Timestamp.Parse()
	now := time.Now()
	ctx.EditEmbed(m, fmt.Sprintf("🏓 | **Ping my mom**\n\n💗 | **Heartbeat**: `%1.fms`\n ⏱️| **Message Delay**: `%1.fms`",
		float64(ctx.Session.HeartbeatLatency().Milliseconds()),
		float64(now.Sub(ts).Milliseconds())))
}

// Test command used for testing
func Test(ctx *Context) {
	flags, err := ctx.FindCommandFlag()
	if err != nil {
		ctx.SendErr(err)
		return
	}

	for _, flag := range flags {
		ctx.Send(fmt.Sprintf("Name: %s\nValue: %s", flag.Name, flag.Value))
	}

}

// Todo todo list command
func Todo(ctx *Context) {
	flags, err := ctx.FindCommandFlag()
	if err != nil {
		ctx.SendErr(err)
		return
	}
	if flags[0].RequiresValue && flags[0].Value == "" {
		ctx.SendErr("no value provided for the flag " + flags[0].Name)
		return
	}

	//db = *NewDB()

	switch flags[0].Name {
	// Add todo item to DB
	case "add":
		//var item *ListItem
		args := flags[0].Value

		// End keywords:
		// second, minute, hour, day, week, month, year
		// --td add do stuff at 11:14am X
		// --td add do stuff in 3 days
		// --td add do stuff at 11:14am in 3 days

		// Get parsed time from Python script

		ctx.NewEmbed(fmt.Sprintf("✅ | Got it, **added** to your todo list!\n```css\n%s\n```", item.Name))

	default:
		ctx.SendCommandHelp()
	}

}
