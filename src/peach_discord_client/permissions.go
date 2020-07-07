package main

// Role represents a discord guild role
type Role struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Color       int    `json:"color"`
	Hoist       bool   `json:"hoist"`
	Position    int    `json:"position"`
	Permissions int    `json:"permissions"`
	Managed     bool   `json:"managed"`
	Mentionable bool   `json:"mentionable"`
}

func (c *Client) hasPermission(channelID string, member GuildMember, perm int) (bool, error) {

	var permissions int

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
