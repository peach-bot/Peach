package main

import (
	"time"

	"github.com/hako/durafmt"
)

func (c *Client) extAboutOnMessage(ctx *EventMessageCreate) error {
	m := NewMessage{
		Embed: Embed{
			Description: "Need help with something? Join the [support server](https://discord.gg/HfrjV3ybEs)!",
			Fields: []*EmbedField{
				{
					Name:   "Version",
					Value:  VERSION,
					Inline: true,
				},
				{
					Name:   "Uptime",
					Value:  durafmt.Parse(time.Now().Sub(c.Starttime)).LimitFirstN(2).String(),
					Inline: true,
				},
			},
			Color: 0xf78c80,
			Author: EmbedAuthor{
				Name:    "About Peach",
				IconURL: "https://cdn.discordapp.com/avatars/608717006132346918/85a667edc36a4a0679ab5473aa109aea.png",
			},
		},
	}
	_, err := c.SendMessage(ctx.ChannelID, m)
	if err != nil {
		return err
	}

	err = ctx.delete(c)
	if err != nil {
		return err
	}

	return nil
}
