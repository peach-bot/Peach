package main

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
