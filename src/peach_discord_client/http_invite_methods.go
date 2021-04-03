package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetInvite(inviteID string, withCounts bool) (*Invite, error) {

	query := addURLArg("", "with_counts", fmt.Sprintf("%t", withCounts))

	resp, body, err := c.Request(http.MethodGet, EndpointInvite(inviteID)+query, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("GetGuildInvites: %s", err.Error())
		return nil, err
	}

	invite := new(Invite)

	err = json.Unmarshal(body, invite)
	if err != nil {
		return nil, err
	}

	return invite, nil
}

func (c *Client) DeleteInvite(inviteID string) error {

	resp, _, err := c.Request(http.MethodGet, EndpointInvite(inviteID), nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("GetGuildInvites: %s", err.Error())
		return err
	}

	return nil
}
