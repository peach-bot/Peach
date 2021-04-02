package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/patrickmn/go-cache"
)

func (c *Client) getChannel(channelid string) (*Channel, error) {

	// Send Request
	_, body, err := c.Request("GET", EndpointChannel(channelid), nil)
	if err != nil {
		return nil, err
	}

	data := new(Channel)

	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	c.ChannelCache.Set(channelid, *data, cache.DefaultExpiration)

	return data, nil
}

// GetChannelMessages returns an array containing a channels messages
func (c *Client) GetChannelMessages(channelid string, around string, before string, after string, limit int) (*[]Message, error) {

	var args string
	if around != "" {
		args = addURLArg(args, "around", around)
	}
	if before != "" {
		args = addURLArg(args, "before", before)
	}
	if after != "" {
		args = addURLArg(args, "after", after)
	}
	if limit != 0 || limit > 100 || limit < 1 {
		args = addURLArg(args, "limit", strconv.Itoa(limit))
	}

	// Send Request
	resp, body, err := c.Request("GET", EndpointChannelMessages(channelid)+args, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		c.Log.Debugf("GetChannelMessages: %s", ErrUnexpectedStatus(http.StatusOK, resp.StatusCode).Error())
		return nil, ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
	}

	data := new([]Message)

	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// DeleteMessage deletes a specific message with a given ID from a specific channel
func (c *Client) DeleteMessage(channelID, messageID string) error {

	// Send Request
	resp, _, err := c.Request("DELETE", EndpointChannelMessage(channelID, messageID), *new(io.Reader))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("DeleteMessage: unexpected response code. Want: 204 No Content, Got: %s instead", resp.Status)
	}

	return nil
}

// SendMessage posts a message to a guild text or DM channel.
func (c *Client) SendMessage(channelid string, message NewMessage) (*Message, error) {

	// Send Request
	body, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	resp, body, err := c.Request("POST", EndpointChannelMessages(channelid), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		c.Log.Debugf("SendMessage: %s", ErrUnexpectedStatus(http.StatusCreated, resp.StatusCode).Error())
		return nil, ErrUnexpectedStatus(http.StatusCreated, resp.StatusCode)
	}

	sentMessage := new(Message)
	err = json.Unmarshal(body, sentMessage)
	if err != nil {
		return nil, err
	}

	return sentMessage, nil
}

// BulkDeleteMessages deletes a lot of messages in a single request, duh
func (c *Client) BulkDeleteMessages(channelid string, messages []string) error {

	if len(messages) == 0 {
		return nil
	}

	if len(messages) == 1 {
		return c.DeleteMessage(channelid, messages[0])
	}

	body := struct {
		Messages []string `json:"messages"`
	}{messages}

	resp, _, err := c.Request("POST", EndpointChannelMessagesBulkDelete(channelid), body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		c.Log.Debugf("BulkDeleteMessages: %s", ErrUnexpectedStatus(http.StatusNoContent, resp.StatusCode).Error())
		return ErrUnexpectedStatus(http.StatusNoContent, resp.StatusCode)
	}

	return nil
}
