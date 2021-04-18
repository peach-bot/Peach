package main

import "time"

var aliasMap = map[string]string{
	"clear":   "clear",
	"c":       "clear",
	"about":   "about",
	"version": "about",
	"help":    "about",
}

func (c *Client) runOnMessage(invoke string, args []string, ctx *EventMessageCreate) error {

	var err error

	command := aliasMap[invoke]

	switch command {
	case "clear":
		err = c.extClearOnMessage(ctx, args)
	case "about":
		err = c.extAboutOnMessage(ctx)
	default:
		err = nil
	}
	return err
}

func (c *Client) handleNoPermission(m *Message) error {
	sorry, err := c.SendMessage(m.ChannelID, NewMessage{":no_entry: It seems like you do not have the permissions to use this command.", false, nil})
	if err != nil {
		return err
	}

	time.Sleep(5 * time.Second)

	err = sorry.Delete(c)
	if err != nil {
		return err
	}
	err = m.Delete(c)
	if err != nil {
		return err
	}
	return nil
}
