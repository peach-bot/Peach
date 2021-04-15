package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/patrickmn/go-cache"
)

type CreateGuildArgs struct {
	Name                        string     `json:"name"`
	Region                      string     `json:"region,omitempty"`
	Icon                        string     `json:"icon,omitempty"`
	VerificationLevel           int        `json:"verification_level,omitempty"`
	DefaultMessageNotifications int        `json:"default_message_notifications,omitempty"`
	ExplicitContentFilter       int        `json:"explicit_content_filter,omitempty"`
	Roles                       []*Role    `json:"roles,omitempty"`
	Channels                    []*Channel `json:"channels,omitempty"`
	AFKChannelID                string     `json:"afk_channel_id,omitempty"`
	AFKTimeout                  int        `json:"afk_timeout,omitempty"`
	SystemChannelID             string     `json:"system_channel_id,omitempty"`
}

// CreateGuild creates a new guild. If args is nil the name will be used. If args are provided the name has to be part of the args object.
func (c *Client) CreateGuild(args CreateGuildArgs) (*Guild, error) {
	resp, body, err := c.Request(http.MethodPost, EndpointGuilds, args)

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("CreateGuild: %s", err.Error())
		return nil, err
	}

	guild := new(Guild)

	err = json.Unmarshal(body, guild)
	if err != nil {
		return nil, err
	}

	c.GuildCache.Set(guild.ID, *guild, cache.DefaultExpiration)

	return guild, nil
}

// GetGuild retrieves the Guild object for a given ID
func (c *Client) GetGuild(ID string) (g *Guild, err error) {

	cachedGuild, cached := c.GuildCache.Get(ID)

	if cached {
		c.Log.Debugf("GetGuild: found %s in cache", ID)
		guild := cachedGuild.(Guild)
		g = &guild
	} else {
		c.Log.Debugf("GetGuild: could not find %s in cache, retrieving via API", ID)
		g, err = c.getGuild(ID, true)
		return
	}

	return
}

func (c *Client) getGuild(guildID string, withCounts bool) (*Guild, error) {

	query := addURLArg("", "with_counts", strconv.FormatBool(withCounts))

	resp, body, err := c.Request(http.MethodGet, EndpointGuild(guildID)+query, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("getGuild: %s", err.Error())
		return nil, err
	}

	guild := new(Guild)

	err = json.Unmarshal(body, guild)
	if err != nil {
		return nil, err
	}

	c.GuildCache.Set(guildID, *guild, cache.DefaultExpiration)

	return guild, nil
}

func (c *Client) GetGuildPreview(guildID string) (*GuildPreview, error) {
	resp, body, err := c.Request(http.MethodGet, EndpointGuildPreview(guildID), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("GetGuildPreview: %s", err.Error())
		return nil, err
	}

	guild := new(GuildPreview)

	err = json.Unmarshal(body, guild)
	if err != nil {
		return nil, err
	}

	return guild, nil
}

type ModifyGuildArgs struct {
	Name                        string `json:"name,omitempty"`
	Icon                        string `json:"icon,omitempty"`
	Splash                      string `json:"splash,omitempty"`
	OwnerID                     string `json:"owner_id,omitempty"`
	Region                      string `json:"region,omitempty"`
	AFKChannelID                string `json:"afk_channel_id,omitempty"`
	AFKTimeout                  int    `json:"afk_timeout,omitempty"`
	VerificationLevel           int    `json:"verification_level,omitempty"`
	DefaultMessageNotifications int    `json:"default_message_notifications,omitempty"`
	ExplicitContentFilter       int    `json:"explicit_content_filter,omitempty"`
	SystemChannelID             string `json:"system_channel_id,omitempty"`
	RulesChannelID              string `json:"rules_channel_id,omitempty"`
	Banner                      string `json:"banner,omitempty,omitempty"`
	PreferredLanguage           string `json:"preferred_locale,omitempty"`
	PublicUpdatesChannelID      string `json:"public_updates_channel_id,omitempty"`
}

func (c *Client) ModifyGuild(guildID string, args ModifyGuildArgs) (*Guild, error) {

	resp, body, err := c.Request(http.MethodPatch, EndpointGuild(guildID), args)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("ModifyGuild: %s", err.Error())
		return nil, err
	}

	guild := new(Guild)

	err = json.Unmarshal(body, guild)
	if err != nil {
		return nil, err
	}

	c.GuildCache.Set(guild.ID, *guild, cache.DefaultExpiration)

	return guild, nil
}

func (c *Client) DeleteGuild(guildID string, withCounts bool) error {

	resp, _, err := c.Request(http.MethodDelete, EndpointGuild(guildID), nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		err = ErrUnexpectedStatus(http.StatusNoContent, resp.StatusCode)
		c.Log.Debugf("DeleteGuild: %s", err.Error())
		return err
	}

	c.GuildCache.Delete(guildID)

	return nil
}

func (c *Client) GetGuildChannels(guildID string) (*[]Channel, error) {
	resp, body, err := c.Request(http.MethodGet, EndpointGuildChannels(guildID), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("GetGuildChannels: %s", err.Error())
		return nil, err
	}

	channels := new([]Channel)

	err = json.Unmarshal(body, channels)
	if err != nil {
		return nil, err
	}

	for _, channel := range *channels {
		c.ChannelCache.Set(channel.ID, channel, cache.DefaultExpiration)
	}

	return channels, nil
}

type CreateGuildChannelArgs struct {
	Name                 string       `json:"name"`
	Type                 int          `json:"type,omitempty"`
	Topic                string       `json:"topic,omitempty"`
	Bitrate              int          `json:"bitrate,omitempty"`
	UserLimit            int          `json:"user_limit,omitempty"`
	SlowModeRateLimit    int          `json:"rate_limit_per_user,omitempty"`
	Position             int          `json:"position,omitempty"`
	PermissionOverwrites []*Overwrite `json:"permission_overwrites,omitempty"`
	ParentId             string       `json:"parent_id,omitempty"`
	NSFW                 bool         `json:"nsfw,omitempty"`
}

func (c *Client) CreateGuildChannel(guildID string, args CreateGuildArgs) (*Channel, error) {
	resp, body, err := c.Request(http.MethodPost, EndpointGuildChannels(guildID), args)

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("CreateGuildChannel: %s", err.Error())
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

type ModifyGuildChannelPositionsArgs struct {
	ChannelID string `json:"id"`
	Position  int    `json:"position,omitempty"`
}

func (c *Client) ModifyGuildChannelPositions(guildID string, args []ModifyGuildChannelPositionsArgs) error {

	resp, _, err := c.Request(http.MethodPatch, EndpointGuildChannels(guildID), args)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		err = ErrUnexpectedStatus(http.StatusNoContent, resp.StatusCode)
		c.Log.Debugf("ModifyGuildChannelPositions: %s", err.Error())
		return err
	}
	return nil
}

func (c *Client) GetGuildMember(guildID string, userID string) (*GuildMember, error) {
	resp, body, err := c.Request(http.MethodGet, EndpointGuildMember(guildID, userID), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("GetGuildMember: %s", err.Error())
		return nil, err
	}

	member := new(GuildMember)

	err = json.Unmarshal(body, member)
	if err != nil {
		return nil, err
	}

	return member, nil
}

func (c *Client) GetGuildMembers(guildID string, limit int, afterid string) (*[]GuildMember, error) {

	var query string
	if limit > 0 {
		query = addURLArg(query, "limit", fmt.Sprint(limit))
	}
	if afterid != "" {
		query = addURLArg(query, "after", afterid)
	}

	resp, body, err := c.Request(http.MethodGet, EndpointGuildMembers(guildID)+query, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("GetGuildMembers: %s", err.Error())
		return nil, err
	}

	members := new([]GuildMember)

	err = json.Unmarshal(body, members)
	if err != nil {
		return nil, err
	}

	return members, nil
}

type ModifyGuildMemberArgs struct {
	Nick      string    `json:"nick,omitempty"`
	Roles     []*string `json:"roles,omitempty"`
	Mute      bool      `json:"mute,omitempty"`
	Deaf      bool      `json:"deaf,omitempty"`
	ChannelID string    `json:"channel_id,omitempty"`
}

func (c *Client) ModifyGuildMember(guildID string, userID string, args ModifyGuildMemberArgs) (*GuildMember, error) {
	resp, body, err := c.Request(http.MethodPatch, EndpointGuildChannels(guildID), args)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("ModifyGuildMember: %s", err.Error())
		return nil, err
	}

	member := new(GuildMember)

	err = json.Unmarshal(body, member)
	if err != nil {
		return nil, err
	}

	return member, nil
}

func (c *Client) ModifyCurrentUserNick(guildID string, nick string) (*string, error) {

	args := struct {
		Nick string `json:"nick,omitempty"`
	}{nick}

	resp, body, err := c.Request(http.MethodPatch, EndpointGuildChannels(guildID), args)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("ModifyCurrentUserNick: %s", err.Error())
		return nil, err
	}

	respbody := new(struct {
		Nick string `json:"nick,omitempty"`
	})

	err = json.Unmarshal(body, respbody)
	if err != nil {
		return nil, err
	}
	newnick := new(string)
	*newnick = respbody.Nick

	return newnick, nil
}

func (c *Client) AddGuildMemberRole(guildID string, userID string, roleID string) error {

	resp, _, err := c.Request(http.MethodPut, EndpointGuildMemberRole(guildID, userID, roleID), nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		err = ErrUnexpectedStatus(http.StatusNoContent, resp.StatusCode)
		c.Log.Debugf("AddGuildMemberRole: %s", err.Error())
		return err
	}
	return nil
}

func (c *Client) RemoveGuildMemberRole(guildID string, userID string, roleID string) error {

	resp, _, err := c.Request(http.MethodDelete, EndpointGuildMemberRole(guildID, userID, roleID), nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		err = ErrUnexpectedStatus(http.StatusNoContent, resp.StatusCode)
		c.Log.Debugf("RemoveGuildMemberRole: %s", err.Error())
		return err
	}
	return nil
}

func (c *Client) RemoveGuildMember(guildID string, userID string) error {

	resp, _, err := c.Request(http.MethodDelete, EndpointGuildMember(guildID, userID), nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		err = ErrUnexpectedStatus(http.StatusNoContent, resp.StatusCode)
		c.Log.Debugf("RemoveGuildMemberRole: %s", err.Error())
		return err
	}
	return nil
}

func (c *Client) GetGuildBans(guildID string) (*[]Ban, error) {
	resp, body, err := c.Request(http.MethodGet, EndpointGuildBans(guildID), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("GetGuildBans: %s", err.Error())
		return nil, err
	}

	bans := new([]Ban)

	err = json.Unmarshal(body, bans)
	if err != nil {
		return nil, err
	}

	return bans, nil
}

func (c *Client) GetGuildBan(guildID string, userID string) (*Ban, error) {
	resp, body, err := c.Request(http.MethodGet, EndpointGuildBan(guildID, userID), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("GetGuildBan: %s", err.Error())
		return nil, err
	}

	ban := new(Ban)

	err = json.Unmarshal(body, ban)
	if err != nil {
		return nil, err
	}

	return ban, nil
}

type CreateGuildBanArgs struct {
	DeleteMessageDays int    `json:"delete_message_days,omitempty"`
	Reason            string `json:"reason,omitempty"`
}

func (c *Client) CreateGuildBan(guildID string, userID string, args CreateGuildBanArgs) error {

	resp, _, err := c.Request(http.MethodPut, EndpointGuildBan(guildID, userID), args)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		err = ErrUnexpectedStatus(http.StatusNoContent, resp.StatusCode)
		c.Log.Debugf("CreateGuildBan: %s", err.Error())
		return err
	}
	return nil
}

func (c *Client) RemoveGuildBan(guildID string, userID string) error {

	resp, _, err := c.Request(http.MethodDelete, EndpointGuildBan(guildID, userID), nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		err = ErrUnexpectedStatus(http.StatusNoContent, resp.StatusCode)
		c.Log.Debugf("RemoveGuildBan: %s", err.Error())
		return err
	}
	return nil
}

func (c *Client) GetGuildRoles(guildID string) (*[]Role, error) {
	resp, body, err := c.Request(http.MethodGet, EndpointGuildChannels(guildID), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("GetGuildChannels: %s", err.Error())
		return nil, err
	}

	roles := new([]Role)

	err = json.Unmarshal(body, roles)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

type CreateGuildRoleArgs struct {
	Name        string `json:"name,omitempty"`
	Permissions string `json:"permissions,omitempty"`
	Color       int    `json:"color,omitempty"`
	Hoist       bool   `json:"hoist,omitempty"`
	Mentionable bool   `json:"mentionable,omitempty"`
}

func (c *Client) CreateGuildRole(guildID string, args CreateGuildRoleArgs) (*Role, error) {
	resp, body, err := c.Request(http.MethodPost, EndpointGuildRoles(guildID), args)

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("CreateGuildRole: %s", err.Error())
		return nil, err
	}

	role := new(Role)

	err = json.Unmarshal(body, role)
	if err != nil {
		return nil, err
	}

	return role, nil
}

type ModifyGuildRolePositionsArgs struct {
	RoleID   string `json:"id"`
	Position int    `json:"position,omitempty"`
}

func (c *Client) ModifyGuildRolePositions(guildID string, args []ModifyGuildRolePositionsArgs) (*[]Role, error) {

	resp, body, err := c.Request(http.MethodPatch, EndpointGuildRoles(guildID), args)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("ModifyGuildChannelPositions: %s", err.Error())
		return nil, err
	}

	roles := new([]Role)

	err = json.Unmarshal(body, roles)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

type ModifyGuildRoleArgs struct {
	Name        string `json:"name,omitempty"`
	Permissions string `json:"permissions,omitempty"`
	Color       int    `json:"color,omitempty"`
	Hoist       bool   `json:"hoist,omitempty"`
	Mentionable bool   `json:"mentionable,omitempty"`
}

func (c *Client) ModifyGuildRole(guildID string, roleID string, args ModifyGuildRoleArgs) (*Role, error) {
	resp, body, err := c.Request(http.MethodPatch, EndpointGuildRole(guildID, roleID), args)

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("ModifyGuildRole: %s", err.Error())
		return nil, err
	}

	role := new(Role)

	err = json.Unmarshal(body, role)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (c *Client) DeleteGuildRole(guildID string, roleID string) error {
	resp, _, err := c.Request(http.MethodPatch, EndpointGuildRole(guildID, roleID), nil)

	if resp.StatusCode != http.StatusNoContent {
		err = ErrUnexpectedStatus(http.StatusNoContent, resp.StatusCode)
		c.Log.Debugf("DeleteGuildRole: %s", err.Error())
		return err
	}

	return nil
}

func (c *Client) GetGuildPruneCount(guildID string, days int, includeRoles []string) (*int, error) {

	query := addURLArg("", "days", fmt.Sprint(days))
	query = addURLArg(query, "include_roles", strings.Join(includeRoles, ","))

	resp, body, err := c.Request(http.MethodGet, EndpointGuildPrune(guildID)+query, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("GetGuildPruneCount: %s", err.Error())
		return nil, err
	}

	respbody := new(struct {
		Amount int `json:"pruned,omitempty"`
	})

	err = json.Unmarshal(body, respbody)
	if err != nil {
		return nil, err
	}
	amount := new(int)
	*amount = respbody.Amount

	return amount, nil
}

type GuildPruneArgs struct {
	Days         int
	ReturnAmount bool
	IncludeRoles []string
}

func (c *Client) GuildPrune(guildID string, args GuildPruneArgs) (*int, error) {
	resp, body, err := c.Request(http.MethodPost, EndpointGuildPrune(guildID), args)

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("GuildPrune: %s", err.Error())
		return nil, err
	}

	if !args.ReturnAmount {
		return nil, nil
	}

	respbody := new(struct {
		Amount int `json:"pruned,omitempty"`
	})

	err = json.Unmarshal(body, respbody)
	if err != nil {
		return nil, err
	}
	amount := new(int)
	*amount = respbody.Amount

	return amount, nil
}

func (c *Client) GetGuildVoiceRegions(guildID string) (*[]VoiceRegion, error) {

	resp, body, err := c.Request(http.MethodGet, EndpointGuildVoiceRegions(guildID), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("GetGuildChannels: %s", err.Error())
		return nil, err
	}

	regions := new([]VoiceRegion)

	err = json.Unmarshal(body, regions)
	if err != nil {
		return nil, err
	}

	return regions, nil
}

func (c *Client) GetGuildInvites(guildID string) (*[]Invite, error) {

	resp, body, err := c.Request(http.MethodGet, EndpointGuildInvites(guildID), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("GetGuildInvites: %s", err.Error())
		return nil, err
	}

	invites := new([]Invite)

	err = json.Unmarshal(body, invites)
	if err != nil {
		return nil, err
	}

	return invites, nil
}

func (c *Client) GetGuildIntegrations(guildID string) (*[]Integration, error) {

	resp, body, err := c.Request(http.MethodGet, EndpointGuildIntegrations(guildID), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("GetGuildIntegrations: %s", err.Error())
		return nil, err
	}

	integrations := new([]Integration)

	err = json.Unmarshal(body, integrations)
	if err != nil {
		return nil, err
	}

	return integrations, nil
}

type CreateGuildIntegrationArgs struct {
	Type          string `json:"type"`
	IntegrationID string `json:"id"`
}

func (c *Client) CreateGuildIntegration(guildID string, args CreateGuildIntegrationArgs) error {

	resp, _, err := c.Request(http.MethodPost, EndpointGuildIntegrations(guildID), args)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		err = ErrUnexpectedStatus(http.StatusNoContent, resp.StatusCode)
		c.Log.Debugf("CreateGuildIntegration: %s", err.Error())
		return err
	}
	return nil
}

type ModifyGuildIntegrationArgs struct {
	EnableEmoticons   bool `json:"enable_emoticons,omitempty"`
	ExpireBehavior    int  `json:"expire_behavior,omitempty"`
	ExpireGracePeriod int  `json:"expire_grace_period,omitempty"`
}

func (c *Client) ModifyGuildIntegration(guildID string, integrationID string, args ModifyGuildIntegrationArgs) error {

	resp, _, err := c.Request(http.MethodPatch, EndpointGuildIntegration(guildID, integrationID), args)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		err = ErrUnexpectedStatus(http.StatusNoContent, resp.StatusCode)
		c.Log.Debugf("ModifyGuildIntegrations: %s", err.Error())
		return err
	}
	return nil
}

func (c *Client) DeleteGuildIntegration(guildID string, integrationID string) error {

	resp, _, err := c.Request(http.MethodDelete, EndpointGuildIntegration(guildID, integrationID), nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		err = ErrUnexpectedStatus(http.StatusNoContent, resp.StatusCode)
		c.Log.Debugf("DeleteGuildIntegration: %s", err.Error())
		return err
	}
	return nil
}

func (c *Client) SyncGuildIntegration(guildID string, integrationID string) error {

	resp, _, err := c.Request(http.MethodPost, EndpointGuildIntegrationSync(guildID, integrationID), nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		err = ErrUnexpectedStatus(http.StatusNoContent, resp.StatusCode)
		c.Log.Debugf("SyncGuildIntegration: %s", err.Error())
		return err
	}
	return nil
}

func (c *Client) GetGuildWidgetSettings(guildID string) (*Widget, error) {

	resp, body, err := c.Request(http.MethodGet, EndpointGuildWidget(guildID), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("GetGuildWidgetSettings: %s", err.Error())
		return nil, err
	}

	widget := new(Widget)

	err = json.Unmarshal(body, widget)
	if err != nil {
		return nil, err
	}

	return widget, nil
}

func (c *Client) ModifyGuildWidget(guildID string, args Widget) (*Widget, error) {

	resp, body, err := c.Request(http.MethodPatch, EndpointGuildWidget(guildID), args)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("ModifyGuildWidget: %s", err.Error())
		return nil, err
	}

	widget := new(Widget)

	err = json.Unmarshal(body, widget)
	if err != nil {
		return nil, err
	}

	return widget, nil
}

func (c *Client) GetGuildVanityURL(guildID string) (*Invite, error) {

	resp, body, err := c.Request(http.MethodGet, EndpointGuildVanityURL(guildID), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = ErrUnexpectedStatus(http.StatusOK, resp.StatusCode)
		c.Log.Debugf("GetGuildVanityURL: %s", err.Error())
		return nil, err
	}

	invite := new(Invite)

	err = json.Unmarshal(body, invite)
	if err != nil {
		return nil, err
	}

	return invite, nil
}
