package main

import (
	"fmt"
	"strconv"
	"time"
)

func (c *Client) extClearOnMessage(ctx *EventMessageCreate, args []string) error {

	// Check if user is allowed to delete messages
	hasPerm, err := c.hasPermission(ctx.ChannelID, ctx.Author, ctx.Member, permissionManageMessages)
	if err != nil {
		return err
	}

	if !hasPerm {
		err = c.handleNoPermission(ctx.Message)
		return err
	}

	if len(args) > 1 || len(args) == 0 {
		c.SendMessage(ctx.ChannelID, NewMessage{"Please provide an amount of messages to clear. Example: `!clear 10`", false, nil})
		return nil
	}

	amount, err := strconv.Atoi(args[0])
	if err != nil || amount > 100 || amount < 1 {
		c.SendMessage(ctx.ChannelID, NewMessage{"Please provide an amount of messages to clear (Max 100). Example: `!clear 10`", false, nil})
		return nil
	}

	err = ctx.Delete(c)
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

	pluralS := ""
	if amount > 1 {
		pluralS = "s"
	}

	err = c.BulkDeleteMessages(ctx.ChannelID, messageIDs)
	if err != nil {
		return err
	}

	success, err := c.SendMessage(ctx.ChannelID, NewMessage{fmt.Sprintf("Deleted %s message%s for you :slight_smile:", args[0], pluralS), false, nil})
	if err != nil {
		return err
	}

	time.Sleep(5 * time.Second)

	err = success.Delete(c)
	if err != nil {
		return err
	}

	return nil
}
