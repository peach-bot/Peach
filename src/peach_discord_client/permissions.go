package main

import (
	"strconv"
)

// Role represents a discord guild role
type Role struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Color       int     `json:"color"`
	Hoist       bool    `json:"hoist"`
	Position    int     `json:"position"`
	Permissions string  `json:"permissions"`
	Managed     bool    `json:"managed"`
	Mentionable bool    `json:"mentionable"`
	Tags        RoleTag `json:"tags,omitempty"`
}

type RoleTag struct {
	BotID             string `json:"bot_id,omitempty"`
	IntegrationID     string `json:"tags,omitempty"`
	PremiumSubscriber string `json:"premium_subscriber,omitempty"` // Don't use this it's weird!!!!!!!!!
}

func (c *Client) hasPermission(channelID string, author User, member GuildMember, perm int) (bool, error) {

	var permissions int

	member.User = author

	channel, err := c.GetChannel(channelID)
	if err != nil {
		return false, err
	}
	guild, err := c.GetGuild(channel.GuildID)
	if err != nil {
		return false, err
	}

	//
	// Compute base permissions
	//

	// If the member is owner of the guild they have all permissions
	if guild.OwnerID == member.User.ID {
		return true, nil
	}

	for _, role := range guild.Roles {
		intPerm, err := strconv.Atoi(role.Permissions)
		if err != nil {
			c.Log.Errorf("Major fucky wucky in c.hasPermission: %s", err.Error())
		}

		// @everone permissions
		if role.ID == guild.ID {
			permissions |= intPerm
		}

		// role permissions
		for _, roleid := range member.Roles {
			if role.ID == *roleid {
				permissions |= intPerm
			}
		}
	}

	// If the member is administrator they have all permissions
	if (permissions & permissionAdministrator) == permissionAdministrator {
		return true, nil
	}

	//
	// Compute channel overwrites
	//

	// @everyone overwrites
	for _, overwrite := range channel.PermissionOverwrites {
		if overwrite.ID == guild.ID {
			intPerm, err := strconv.Atoi(overwrite.Deny)
			if err != nil {
				c.Log.Errorf("Major fucky wucky in c.hasPermission: %s", err.Error())
			}
			permissions &= ^intPerm
			intPerm, err = strconv.Atoi(overwrite.Allow)
			if err != nil {
				c.Log.Errorf("Major fucky wucky in c.hasPermission: %s", err.Error())
			}
			permissions |= intPerm
		}
	}

	// role overwrites
	for _, overwrite := range channel.PermissionOverwrites {
		for _, roleid := range member.Roles {
			if overwrite.ID == *roleid {
				intPerm, err := strconv.Atoi(overwrite.Deny)
				if err != nil {
					c.Log.Errorf("Major fucky wucky in c.hasPermission: %s", err.Error())
				}
				permissions &= ^intPerm
				intPerm, err = strconv.Atoi(overwrite.Allow)
				if err != nil {
					c.Log.Errorf("Major fucky wucky in c.hasPermission: %s", err.Error())
				}
				permissions |= intPerm
			}
		}
	}

	// member specific overwrites
	for _, overwrite := range channel.PermissionOverwrites {
		if overwrite.ID == member.User.ID {
			intPerm, err := strconv.Atoi(overwrite.Deny)
			if err != nil {
				c.Log.Errorf("Major fucky wucky in c.hasPermission: %s", err.Error())
			}
			permissions &= ^intPerm
			intPerm, err = strconv.Atoi(overwrite.Allow)
			if err != nil {
				c.Log.Errorf("Major fucky wucky in c.hasPermission: %s", err.Error())
			}
			permissions |= intPerm
		}
	}

	// check if member has permission
	if (permissions & perm) == perm {
		return true, nil
	}
	return false, nil
}
