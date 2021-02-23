package main

import (
	"fmt"
	"time"
)

// Ping command which returns a message
func Ping(ctx *Context) {
	m := ctx.NewEmbed("Pinging....")
	ts, _ := m.Timestamp.Parse()
	now := time.Now()
	ctx.EditEmbed(m, fmt.Sprintf("🏓 | **Pong my ping**\n\n💗 | **Heartbeat**: `%1.fms`\n ⏱️| **Message Delay**: `%1.fms`",
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
