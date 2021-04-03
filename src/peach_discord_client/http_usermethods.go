package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetUserGuilds() (*[]Guild, error) {

	// Send Request
	resp, body, err := c.Request("GET", EndpointUserGuilds("@me"), *new(io.Reader))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("GetUserGuilds: %s", err.Error())
		return nil, err
	}

	var data []Guild

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
