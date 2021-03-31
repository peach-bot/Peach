package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func (c *Client) getUserGuilds() (*[]Guild, error) {

	// Send Request
	req, err := http.NewRequest("GET", EndpointUserGuilds("@me"), *new(io.Reader))
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
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		c.Log.Debug(string(bodyBytes))
		return nil, errors.New(fmt.Sprintf("Got unexpected response from Discord API: Want: 200 OK Got: %s", resp.Status))
	}

	var data []Guild

	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
