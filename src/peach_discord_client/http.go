package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// SetDefaultRequestHeaders adds authorization and content type to request header
func (c *Client) SetDefaultRequestHeaders(req *http.Request) *http.Request {
	req.Header.Add("Authorization", fmt.Sprintf("Bot %s", c.TOKEN))
	req.Header.Add("Content-Type", "application/json")
	return req
}

// SendMessage posts a message to a guild text or DM channel.
func (c *Client) SendMessage(channelid string, message NewMessage) error {

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", EndpointChannelMessages(channelid), bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req = c.SetDefaultRequestHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}
	bodyString := string(bodyBytes)

	c.Log.Debug(resp.StatusCode, bodyString)

	return nil
}
