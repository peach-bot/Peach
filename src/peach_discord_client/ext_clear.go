package main

import (
	"fmt"
	"strconv"
)

func (c *Client) extClearOnMessage(ctx *EventMessageCreate, args []string) error {
	if len(args) > 1 || len(args) == 0 {
		c.SendMessage(ctx.ChannelID, NewMessage{"Please provide an amount of messages to clear. Example: `!clear 10`", false, nil})
		return nil
	}

	amount, err := strconv.Atoi(args[0])
	if err != nil || amount > 50 || amount < 2 {
		c.SendMessage(ctx.ChannelID, NewMessage{"Please provide an amount (Min: 2, Max: 50) of messages to clear. Example: `!clear 10`", false, nil})
		return nil
	}
	messages, err := c.GetChannelMessages(ctx.ChannelID, "", "", "", amount)
	if err != nil {
		return err
	}
	var messageIDs []string
	for _, message := range *messages {
		messageIDs = append(messageIDs, message.ID)
	}

	err = c.BulkDeleteMessages(ctx.ChannelID, messageIDs)
	if err != nil {
		return err
	}

	_, err = c.SendMessage(ctx.ChannelID, NewMessage{fmt.Sprintf("Deleted %s messages for you :slight_smile:", args[0]), false, nil})
	if err != nil {
		return err
	}

	return nil
}
