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
