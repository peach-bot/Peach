package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/patrickmn/go-cache"
)

func (c *Client) getChannel(channelid string) (*Channel, error) {

	// Send Request
	req, err := http.NewRequest("GET", EndpointChannel(channelid), *new(io.Reader))
	if err != nil {
		return nil, err
	}

	req = c.SetDefaultRequestHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Decode Body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}

	data := new(Channel)

	err = json.Unmarshal(bodyBytes, data)
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
	req, err := http.NewRequest("GET", EndpointChannelMessages(channelid)+args, *new(io.Reader))
	if err != nil {
		return nil, err
	}

	req = c.SetDefaultRequestHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Decode Body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}

	data := new([]Message)

	err = json.Unmarshal(bodyBytes, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// DeleteMessage deletes a specific message with a given ID from a specific channel
func (c *Client) DeleteMessage(channelID, messageID string) error {

	// Send Request
	req, err := http.NewRequest("DELETE", EndpointChannelMessage(channelID, messageID), *new(io.Reader))
	if err != nil {
		return err
	}

	req = c.SetDefaultRequestHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("BulkDeleteMessages: unexpected response code. Want: 204 No Content, Got: %s instead", resp.Status)
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

	req, err := http.NewRequest("POST", EndpointChannelMessages(channelid), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req = c.SetDefaultRequestHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Decode Body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}

	sentMessage := new(Message)
	err = json.Unmarshal(bodyBytes, sentMessage)
	if err != nil {
		return nil, err
	}

	return sentMessage, nil
}

// BulkDeleteMessages deletes a lot of messages in a single request, duh
func (c *Client) BulkDeleteMessages(channelid string, messages []string) error {

	bodystruct := struct {
		Messages []string `json:"messages"`
	}{messages}

	// Send Request
	body, err := json.Marshal(bodystruct)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", EndpointChannelMessagesBulkDelete(channelid), bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req = c.SetDefaultRequestHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("BulkDeleteMessages: unexpected response code. Want: 204 No Content, Got: %s instead", resp.Status)
	}

	return nil
}
