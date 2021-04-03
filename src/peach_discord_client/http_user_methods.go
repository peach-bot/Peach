package main

import (
	"encoding/json"
	"net/http"
)

func (c *Client) GetUserGuilds() (*[]Guild, error) {

	// Send Request
	resp, body, err := c.Request("GET", EndpointUserGuilds("@me"), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("GetUserGuilds: %s", err.Error())
		return nil, err
	}

	guild := new([]Guild)

	err = json.Unmarshal(body, guild)
	if err != nil {
		return nil, err
	}

	return guild, nil
}
