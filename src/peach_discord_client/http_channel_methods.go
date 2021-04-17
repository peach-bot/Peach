package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/patrickmn/go-cache"
)

// GetChannel retrieves the Channel object for a given ID
func (c *Client) GetChannel(ID string) (ch *Channel, err error) {

	cachedChannel, cached := c.ChannelCache.Get(ID)

	if cached {
		c.Log.Debugf("GetChannel: found %s in cache", ID)
		channel := cachedChannel.(Channel)
		ch = &channel
	} else {
		c.Log.Debugf("GetChannel: could not find %s in cache, retrieving via API", ID)
		ch, err = c.getChannel(ID)
		return
	}

	return
}

func (c *Client) getChannel(channelID string) (*Channel, error) {

	// Send Request
	resp, body, err := c.Request("GET", EndpointChannel(channelID), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("getChannel: %s", err.Error())
		return nil, err
	}

	guild := new(Channel)

	err = json.Unmarshal(body, guild)
	if err != nil {
		return nil, err
	}

	c.ChannelCache.Set(channelID, *guild, cache.DefaultExpiration)

	return guild, nil
}

type ModifyChannelArgs struct {
	Name                 string       `json:"string,omitempty"`
	Type                 int          `json:"type,omitempty"`
	Position             int          `json:"position,omitempty"`
	Topic                string       `json:"topic,omitempty"`
	NSFW                 bool         `json:"nsfw,omitempty"`
	RateLimitPerUser     int          `json:"rate_limit_per_user,omitempty"`
	Bitrate              int          `json:"bitrate,omitempty"`
	UserLimit            int          `json:"user_limit,omitempty"`
	PermissionOverwrites []*Overwrite `json:"permission_overwrites,omitempty"`
	ParentID             string       `json:"partent_id,omitempty"`
	RTCRegion            string       `json:"rtc_region,omitempty"`
	VideoQualityMode     int          `json:"video_quality_mode,omitempty"`
}

func (c *Client) ModifyChannel(guildID string, args ModifyChannelArgs) (*Channel, error) {

	resp, body, err := c.Request(http.MethodPatch, EndpointGuild(guildID), args)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("ModifyGuild: %s", err.Error())
		return nil, err
	}

	channel := new(Channel)

	err = json.Unmarshal(body, channel)
	if err != nil {
		return nil, err
	}

	c.ChannelCache.Set(channel.ID, *channel, cache.DefaultExpiration)

	return channel, nil
}

func (c *Client) DeleteChannel(channelID string, withCounts bool) (*Channel, error) {

	resp, body, err := c.Request(http.MethodDelete, EndpointGuild(channelID), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusNoContent {
		err = ErrUnexpectedStatus(http.StatusNoContent, resp.StatusCode)
		c.Log.Debugf("DeleteChannel: %s", err.Error())
		return nil, err
	}
	data := new(Channel)

	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	c.ChannelCache.Delete(channelID)

	return data, nil
}

// GetChannelMessages returns an array containing a channels messages
func (c *Client) GetChannelMessages(channelID string, around string, before string, after string, limit int) (*[]Message, error) {

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
	resp, body, err := c.Request("GET", EndpointChannelMessages(channelID)+args, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("GetChannelMessages: %s", err.Error())
		return nil, err
	}

	data := new([]Message)

	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) GetChannelMessage(channelID string, messageID string) (*Message, error) {

	// Send Request
	resp, body, err := c.Request("GET", EndpointChannelMessage(channelID, messageID), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("GetChannelMessage: %s", err.Error())
		return nil, err
	}

	message := new(Message)

	err = json.Unmarshal(body, message)
	if err != nil {
		return nil, err
	}

	return message, nil
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
func (c *Client) SendMessage(channelID string, message NewMessage) (*Message, error) {

	resp, body, err := c.Request("POST", EndpointChannelMessages(channelID), message)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("SendMessage: %s", err.Error())
		return nil, err
	}

	sentMessage := new(Message)
	err = json.Unmarshal(body, sentMessage)
	if err != nil {
		return nil, err
	}

	return sentMessage, nil
}

// BulkDeleteMessages deletes a lot of messages in a single request, duh
func (c *Client) BulkDeleteMessages(channelID string, messages []string) error {

	if len(messages) == 0 {
		return nil
	}

	if len(messages) == 1 {
		return c.DeleteMessage(channelID, messages[0])
	}

	body := struct {
		Messages []string `json:"messages"`
	}{messages}

	c.Log.Debug(body)

	resp, _, err := c.Request("POST", EndpointChannelMessagesBulkDelete(channelID), body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		err = ErrUnexpectedStatus(http.StatusNoContent, resp.StatusCode)
		c.Log.Debugf("BulkDeleteMessages: %s", err.Error())
		return err
	}

	return nil
}
