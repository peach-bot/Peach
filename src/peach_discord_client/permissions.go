package main

import "time"

// Role represents a discord guild role
type Role struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Color       int     `json:"color"`
	Hoist       bool    `json:"hoist"`
	Position    int     `json:"position"`
	Permissions int     `json:"permissions"`
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
	guild, err := c.getGuild(channel.GuildID)
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

		// @everone permissions
		if role.ID == guild.ID {
			permissions |= role.Permissions
		}

		// role permissions
		for _, roleid := range member.Roles {
			if role.ID == *roleid {
				permissions |= role.Permissions
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
			permissions &= ^overwrite.Deny
			permissions |= overwrite.Allow
		}
	}

	// role overwrites
	for _, overwrite := range channel.PermissionOverwrites {
		for _, roleid := range member.Roles {
			if overwrite.ID == *roleid {
				permissions &= ^overwrite.Deny
				permissions |= overwrite.Allow
			}
		}
	}

	// member specific overwrites
	for _, overwrite := range channel.PermissionOverwrites {
		if overwrite.ID == member.User.ID {
			permissions &= ^overwrite.Deny
			permissions |= overwrite.Allow
		}
	}

	// check if member has permission
	if (permissions & perm) == perm {
		return true, nil
	}
	return false, nil
}

func (c *Client) handleNoPermission(m *Message) error {
	sorry, err := c.SendMessage(m.ChannelID, NewMessage{":no_entry: It seems like you do not have the permissions to use this command.", false, nil})
	if err != nil {
		return err
	}

	time.Sleep(5 * time.Second)

	err = sorry.delete(c)
	if err != nil {
		return err
	}
	err = m.delete(c)
	if err != nil {
		return err
	}
	return nil
}
