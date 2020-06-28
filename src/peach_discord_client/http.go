package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

// SetDefaultRequestHeaders adds authorization and content type to request header
func (c *Client) SetDefaultRequestHeaders(req *http.Request) *http.Request {
	req.Header.Add("Authorization", fmt.Sprintf("Bot %s", c.TOKEN))
	req.Header.Add("Content-Type", "application/json")
	return req
}

func addURLArg(query string, key string, value string) string {
	if query == "" {
		return fmt.Sprintf("?%s=%s", key, value)
	}
	return fmt.Sprintf("%s&%s=%s", query, key, value)
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

// GetChannelMessages posts a message to a guild text or DM channel.
func (c *Client) GetChannelMessages(channelid string, around string, before string, after string, limit int) (*[]Message, error) {

	var urlargs string
	if around != "" {
		urlargs = addURLArg(urlargs, "around", around)
	}
	if before != "" {
		urlargs = addURLArg(urlargs, "before", before)
	}
	if after != "" {
		urlargs = addURLArg(urlargs, "after", after)
	}
	if limit != 0 {
		urlargs = addURLArg(urlargs, "limit", strconv.Itoa(limit))
	}

	// Send Request
	req, err := http.NewRequest("GET", EndpointChannelMessages(channelid)+queryargs, *new(io.Reader))
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
