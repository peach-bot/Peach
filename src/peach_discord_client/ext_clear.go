package main

import (
	"fmt"
	"strconv"
	"time"
)

func (c *Client) extClearOnMessage(ctx *EventMessageCreate, args []string) error {

	// Check if user is allowed to delete messages
	hasPerm, err := c.hasPermission(ctx.ChannelID, ctx.Member, permissionManageMessages)

	if !hasPerm {
		c.SendMessage(ctx.ChannelID, NewMessage{":no_entry: It seems like you do not have the permissions to use this command.", false, nil})
	}

	if len(args) > 1 || len(args) == 0 {
		c.SendMessage(ctx.ChannelID, NewMessage{"Please provide an amount of messages to clear. Example: `!clear 10`", false, nil})
		return nil
	}

	amount, err := strconv.Atoi(args[0])
	if err != nil || amount > 50 || amount < 2 {
		c.SendMessage(ctx.ChannelID, NewMessage{"Please provide an amount (Min: 2, Max: 50) of messages to clear. Example: `!clear 10`", false, nil})
		return nil
	}

	err = ctx.delete(c)
	if err != nil {
		return err
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

	success, err := c.SendMessage(ctx.ChannelID, NewMessage{fmt.Sprintf("Deleted %s messages for you :slight_smile:", args[0]), false, nil})
	if err != nil {
		return err
	}

	time.Sleep(5 * time.Second)

	err = success.delete(c)
	if err != nil {
		return err
	}

	return nil
}
