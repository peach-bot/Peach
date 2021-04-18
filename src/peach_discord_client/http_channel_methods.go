package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
	resp, body, err := c.Request(http.MethodGet, EndpointChannel(channelID), nil)
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
func (c *Client) GetChannelMessages(channelID, around, before, after string, limit int) (*[]Message, error) {

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
	resp, body, err := c.Request(http.MethodGet, EndpointChannelMessages(channelID)+args, nil)
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

func (c *Client) GetChannelMessage(channelID, messageID string) (*Message, error) {

	// Send Request
	resp, body, err := c.Request(http.MethodGet, EndpointChannelMessage(channelID, messageID), nil)
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

// SendMessage posts a message to a guild text or DM channel.
func (c *Client) SendMessage(channelID string, message NewMessage) (*Message, error) {

	resp, body, err := c.Request(http.MethodPost, EndpointChannelMessages(channelID), message)
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

func (c *Client) CrosspostMessage(channelID, messageID string) (*Message, error) {

	resp, body, err := c.Request(http.MethodPost, EndpointChannelCrosspostMessage(channelID, messageID), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("CrosspostMessage: %s", err.Error())
		return nil, err
	}

	message := new(Message)
	err = json.Unmarshal(body, message)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (c *Client) CreateReaction(channelID, messageID, emoji string, customEmoji Emoji) error {

	if emoji == "" {
		emoji = customEmoji.Name + ":" + customEmoji.ID
	}

	resp, _, err := c.Request(http.MethodPost, EndpointMessageReaction(channelID, messageID, url.QueryEscape(emoji), "@me"), nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		err = ErrUnexpectedStatus(http.StatusNoContent, resp.StatusCode)
		c.Log.Debugf("CreateReaction: %s", err.Error())
		return err
	}

	return nil
}

func (c *Client) DeleteBotReaction(channelID, messageID, emoji string, customEmoji Emoji) error {

	if emoji == "" {
		emoji = customEmoji.Name + ":" + customEmoji.ID
	}

	resp, _, err := c.Request(http.MethodDelete, EndpointMessageReaction(channelID, messageID, url.QueryEscape(emoji), "@me"), nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		err = ErrUnexpectedStatus(http.StatusNoContent, resp.StatusCode)
		c.Log.Debugf("DeleteBotReaction: %s", err.Error())
		return err
	}

	return nil
}

func (c *Client) DeleteUserReaction(channelID, messageID, emoji string, customEmoji Emoji, userID string) error {

	if emoji == "" {
		emoji = customEmoji.Name + ":" + customEmoji.ID
	}

	resp, _, err := c.Request(http.MethodDelete, EndpointMessageReaction(channelID, messageID, url.QueryEscape(emoji), userID), nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		err = ErrUnexpectedStatus(http.StatusNoContent, resp.StatusCode)
		c.Log.Debugf("DeleteUserReaction: %s", err.Error())
		return err
	}

	return nil
}

func (c *Client) GetReactions(channelID, messageID, emoji string, customEmoji Emoji) (*[]User, error) {

	if emoji == "" {
		emoji = customEmoji.Name + ":" + customEmoji.ID
	}

	// Send Request
	resp, body, err := c.Request(http.MethodGet, EndpointMessageReactions(channelID, messageID, url.QueryEscape(emoji)), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("GetReactions: %s", err.Error())
		return nil, err
	}

	users := new([]User)

	err = json.Unmarshal(body, users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (c *Client) DeleteAllReactions(channelID, messageID string) error {

	// Send Request
	resp, _, err := c.Request(http.MethodDelete, EndpointMessageReactionsAll(channelID, messageID), nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("DeleteAllReactions: %s", err.Error())
		return err
	}

	return nil
}

func (c *Client) DeleteEmojiReactions(channelID, messageID, emoji string, customEmoji Emoji) error {

	if emoji == "" {
		emoji = customEmoji.Name + ":" + customEmoji.ID
	}

	// Send Request
	resp, _, err := c.Request(http.MethodDelete, EndpointMessageReactions(channelID, messageID, url.QueryEscape(emoji)), nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("DeleteEmojiReactions: %s", err.Error())
		return err
	}

	return nil
}

type EditMessageArgs struct {
	Content         string          `json:"content,omitempty"`
	Embed           Embed           `json:"embed,omitempty"`
	Flags           int             `json:"flags,omitempty"`
	AllowedMentions AllowedMentions `json:"allowed_mentions,omitempty"`
}

// SendMessage posts a message to a guild text or DM channel.
func (c *Client) EditMessage(channelID, messageID string, args EditMessageArgs) (*Message, error) {

	resp, body, err := c.Request(http.MethodPatch, EndpointChannelMessage(channelID, messageID), args)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("EditMessage: %s", err.Error())
		return nil, err
	}

	sentMessage := new(Message)
	err = json.Unmarshal(body, sentMessage)
	if err != nil {
		return nil, err
	}

	return sentMessage, nil
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

	resp, _, err := c.Request(http.MethodPost, EndpointChannelMessagesBulkDelete(channelID), body)
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

type EditChannelPermissionsArgs struct {
	Allow string `json:"allow"`
	Deny  string `json:"deny"`
	Type  string `json:"type"`
}

func (c *Client) EditChannelPermissions(channelID, overwriteID string, args EditChannelPermissionsArgs) error {

	resp, _, err := c.Request(http.MethodPut, EndpointChannelPermission(channelID, overwriteID), args)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		err = ErrUnexpectedStatus(http.StatusNoContent, resp.StatusCode)
		c.Log.Debugf("EditChannelPermissions: %s", err.Error())
		return err
	}

	return nil
}

func (c *Client) GetChannelInvites(channelID string) (*[]Invite, error) {

	resp, body, err := c.Request(http.MethodGet, EndpointChannelInvites(channelID), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("GetChannelInvites: %s", err.Error())
		return nil, err
	}

	invites := new([]Invite)

	err = json.Unmarshal(body, invites)
	if err != nil {
		return nil, err
	}

	return invites, nil
}

type CreateChannelInviteArgs struct {
	MaxAge              int    `json:"max_age,omitempty"`
	MaxUses             int    `json:"max_uses,omitempty"`
	Temporary           bool   `json:"temporary,omitempty"`
	Unique              bool   `json:"unique,omitempty"`
	TargetType          int    `json:"target_type,omitempty"`
	TargetUserID        string `json:"target_user_id,omitempty"`
	TargetApplicationID string `json:"target_application_id,omitempty"`
}

func (c *Client) CreateChannelInvite(channelID string, args CreateChannelInviteArgs) (*Invite, error) {

	resp, body, err := c.Request(http.MethodPost, EndpointChannelInvites(channelID), args)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("CreateChannelInvite: %s", err.Error())
		return nil, err
	}

	invite := new(Invite)

	err = json.Unmarshal(body, invite)
	if err != nil {
		return nil, err
	}

	return invite, nil
}

func (c *Client) DeleteChannelPermission(channelID, overwriteID string) error {

	resp, _, err := c.Request(http.MethodDelete, EndpointChannelPermission(channelID, overwriteID), nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		err = ErrUnexpectedStatus(http.StatusNoContent, resp.StatusCode)
		c.Log.Debugf("DeleteChannelPermission: %s", err.Error())
		return err
	}

	return nil
}

func (c *Client) FollowNewsChannel(channelID, targetChannelID string) error {

	args := struct {
		WebhookChannelID string `json:"webhook_channel_id"`
	}{
		WebhookChannelID: targetChannelID,
	}

	resp, _, err := c.Request(http.MethodPost, EndpointChannelFollowers(channelID), args)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("FollowNewsChannel: %s", err.Error())
		return err
	}

	return nil
}

func (c *Client) TriggerTypingIndicator(channelID string) error {

	resp, _, err := c.Request(http.MethodPost, EndpointChannelTyping(channelID), nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		err = ErrUnexpectedStatus(http.StatusNoContent, resp.StatusCode)
		c.Log.Debugf("TriggerTypingIndicator: %s", err.Error())
		return err
	}

	return nil
}

func (c *Client) GetPinnedMessages(channelID string) (*[]Message, error) {

	// Send Request
	resp, body, err := c.Request(http.MethodGet, EndpointChannelMessagesPins(channelID), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("GetPinnedMessages: %s", err.Error())
		return nil, err
	}

	messages := new([]Message)

	err = json.Unmarshal(body, messages)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (c *Client) PinMessage(channelID, messageID string) error {

	resp, _, err := c.Request(http.MethodPut, EndpointChannelMessagePin(channelID, messageID), nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		err = ErrUnexpectedStatus(http.StatusNoContent, resp.StatusCode)
		c.Log.Debugf("PinMessage: %s", err.Error())
		return err
	}

	return nil
}

func (c *Client) DeletePinnedMessage(channelID, messageID string) error {

	resp, _, err := c.Request(http.MethodDelete, EndpointChannelMessagePin(channelID, messageID), nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		err = ErrUnexpectedStatus(http.StatusNoContent, resp.StatusCode)
		c.Log.Debugf("DeletePinnedMessage: %s", err.Error())
		return err
	}

	return nil
}
