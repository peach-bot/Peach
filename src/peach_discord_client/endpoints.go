package main

import "strconv"

// APIVersion duh
var APIVersion = "6"

// Known Discord API Endpoints.
var (
	EndpointStatus     = "https://status.discord.com/api/v2/"
	EndpointSm         = EndpointStatus + "scheduled-maintenances/"
	EndpointSmActive   = EndpointSm + "active.json"
	EndpointSmUpcoming = EndpointSm + "upcoming.json"

	EndpointDiscord    = "https://discord.com/"
	EndpointAPI        = EndpointDiscord + "api/v" + APIVersion + "/"
	EndpointGuilds     = EndpointAPI + "guilds/"
	EndpointChannels   = EndpointAPI + "channels/"
	EndpointUsers      = EndpointAPI + "users/"
	EndpointGateway    = EndpointAPI + "gateway"
	EndpointGatewayBot = EndpointGateway + "/bot"
	EndpointWebhooks   = EndpointAPI + "webhooks/"

	EndpointCDN             = "https://cdn.discord.com/"
	EndpointCDNAttachments  = EndpointCDN + "attachments/"
	EndpointCDNAvatars      = EndpointCDN + "avatars/"
	EndpointCDNIcons        = EndpointCDN + "icons/"
	EndpointCDNSplashes     = EndpointCDN + "splashes/"
	EndpointCDNChannelIcons = EndpointCDN + "channel-icons/"
	EndpointCDNBanners      = EndpointCDN + "banners/"

	EndpointAuth           = EndpointAPI + "auth/"
	EndpointLogin          = EndpointAuth + "login"
	EndpointLogout         = EndpointAuth + "logout"
	EndpointVerify         = EndpointAuth + "verify"
	EndpointVerifyResend   = EndpointAuth + "verify/resend"
	EndpointForgotPassword = EndpointAuth + "forgot"
	EndpointResetPassword  = EndpointAuth + "reset"
	EndpointRegister       = EndpointAuth + "register"

	EndpointVoice        = EndpointAPI + "/voice/"
	EndpointVoiceRegions = EndpointVoice + "regions"
	EndpointVoiceIce     = EndpointVoice + "ice"

	EndpointTutorial           = EndpointAPI + "tutorial/"
	EndpointTutorialIndicators = EndpointTutorial + "indicators"

	EndpointTrack        = EndpointAPI + "track"
	EndpointSso          = EndpointAPI + "sso"
	EndpointReport       = EndpointAPI + "report"
	EndpointIntegrations = EndpointAPI + "integrations"

	EndpointUser       = func(useroleID string) string { return EndpointUsers + useroleID }
	EndpointUserAvatar = func(useroleID, avataroleID string) string {
		return EndpointCDNAvatars + useroleID + "/" + avataroleID + ".png"
	}
	EndpointUserAvatarAnimated = func(useroleID, avataroleID string) string {
		return EndpointCDNAvatars + useroleID + "/" + avataroleID + ".gif"
	}
	EndpointDefaultUserAvatar = func(uDiscriminator string) string {
		uDiscriminatorInt, _ := strconv.Atoi(uDiscriminator)
		return EndpointCDN + "embed/avatars/" + strconv.Itoa(uDiscriminatorInt%5) + ".png"
	}
	EndpointUserSettings      = func(useroleID string) string { return EndpointUsers + useroleID + "/settings" }
	EndpointUserGuilds        = func(useroleID string) string { return EndpointUsers + useroleID + "/guilds" }
	EndpointUserGuild         = func(useroleID, guildID string) string { return EndpointUsers + useroleID + "/guilds/" + guildID }
	EndpointUserGuildSettings = func(useroleID, guildID string) string {
		return EndpointUsers + useroleID + "/guilds/" + guildID + "/settings"
	}
	EndpointUserChannels    = func(useroleID string) string { return EndpointUsers + useroleID + "/channels" }
	EndpointUserDevices     = func(useroleID string) string { return EndpointUsers + useroleID + "/devices" }
	EndpointUserConnections = func(useroleID string) string { return EndpointUsers + useroleID + "/connections" }
	EndpointUserNotes       = func(useroleID string) string { return EndpointUsers + "@me/notes/" + useroleID }

	EndpointGuild           = func(guildID string) string { return EndpointGuilds + guildID }
	EndpointGuildChannels   = func(guildID string) string { return EndpointGuilds + guildID + "/channels" }
	EndpointGuildMembers    = func(guildID string) string { return EndpointGuilds + guildID + "/members" }
	EndpointGuildMember     = func(guildID, useroleID string) string { return EndpointGuilds + guildID + "/members/" + useroleID }
	EndpointGuildMemberRole = func(guildID, useroleID, roleID string) string {
		return EndpointGuilds + guildID + "/members/" + useroleID + "/roles/" + roleID
	}
	EndpointGuildBans         = func(guildID string) string { return EndpointGuilds + guildID + "/bans" }
	EndpointGuildBan          = func(guildID, useroleID string) string { return EndpointGuilds + guildID + "/bans/" + useroleID }
	EndpointGuildIntegrations = func(guildID string) string { return EndpointGuilds + guildID + "/integrations" }
	EndpointGuildIntegration  = func(guildID, integrationID string) string {
		return EndpointGuilds + guildID + "/integrations/" + integrationID
	}
	EndpointGuildIntegrationSync = func(guildID, integrationID string) string {
		return EndpointGuilds + guildID + "/integrations/" + integrationID + "/sync"
	}
	EndpointGuildRoles        = func(guildID string) string { return EndpointGuilds + guildID + "/roles" }
	EndpointGuildRole         = func(guildID, roleID string) string { return EndpointGuilds + guildID + "/roles/" + roleID }
	EndpointGuildInvites      = func(guildID string) string { return EndpointGuilds + guildID + "/invites" }
	EndpointGuildEmbed        = func(guildID string) string { return EndpointGuilds + guildID + "/embed" }
	EndpointGuildPrune        = func(guildID string) string { return EndpointGuilds + guildID + "/prune" }
	EndpointGuildIcon         = func(guildID, hash string) string { return EndpointCDNIcons + guildID + "/" + hash + ".png" }
	EndpointGuildIconAnimated = func(guildID, hash string) string { return EndpointCDNIcons + guildID + "/" + hash + ".gif" }
	EndpointGuildSplash       = func(guildID, hash string) string { return EndpointCDNSplashes + guildID + "/" + hash + ".png" }
	EndpointGuildWebhooks     = func(guildID string) string { return EndpointGuilds + guildID + "/webhooks" }
	EndpointGuildAuditLogs    = func(guildID string) string { return EndpointGuilds + guildID + "/audit-logs" }
	EndpointGuildEmojis       = func(guildID string) string { return EndpointGuilds + guildID + "/emojis" }
	EndpointGuildEmoji        = func(guildID, emojiID string) string { return EndpointGuilds + guildID + "/emojis/" + emojiID }
	EndpointGuildBanner       = func(guildID, hash string) string { return EndpointCDNBanners + guildID + "/" + hash + ".png" }

	EndpointChannel            = func(channelID string) string { return EndpointChannels + channelID }
	EndpointChannelPermissions = func(channelID string) string { return EndpointChannels + channelID + "/permissions" }
	EndpointChannelPermission  = func(channelID, overwriteID string) string {
		return EndpointChannels + channelID + "/permissions/" + overwriteID
	}
	EndpointChannelInvites  = func(channelID string) string { return EndpointChannels + channelID + "/invites" }
	EndpointChannelTyping   = func(channelID string) string { return EndpointChannels + channelID + "/typing" }
	EndpointChannelMessages = func(channelID string) string { return EndpointChannels + channelID + "/messages" }
	EndpointChannelMessage  = func(channelID, messagemojiID string) string {
		return EndpointChannels + channelID + "/messages/" + messagemojiID
	}
	EndpointChannelMessageAck = func(channelID, messagemojiID string) string {
		return EndpointChannels + channelID + "/messages/" + messagemojiID + "/ack"
	}
	EndpointChannelMessagesBulkDelete = func(channelID string) string { return EndpointChannel(channelID) + "/messages/bulk-delete" }
	EndpointChannelMessagesPins       = func(channelID string) string { return EndpointChannel(channelID) + "/pins" }
	EndpointChannelMessagePin         = func(channelID, messagemojiID string) string {
		return EndpointChannel(channelID) + "/pins/" + messagemojiID
	}

	EndpointGroupIcon = func(channelID, hash string) string { return EndpointCDNChannelIcons + channelID + "/" + hash + ".png" }

	EndpointChannelWebhooks = func(channelID string) string { return EndpointChannel(channelID) + "/webhooks" }
	EndpointWebhook         = func(webhookID string) string { return EndpointWebhooks + webhookID }
	EndpointWebhookToken    = func(webhookID, token string) string { return EndpointWebhooks + webhookID + "/" + token }

	EndpointMessageReactionsAll = func(channelID, messagemojiID string) string {
		return EndpointChannelMessage(channelID, messagemojiID) + "/reactions"
	}
	EndpointMessageReactions = func(channelID, messagemojiID, emojiID string) string {
		return EndpointChannelMessage(channelID, messagemojiID) + "/reactions/" + emojiID
	}
	EndpointMessageReaction = func(channelID, messagemojiID, emojiID, useroleID string) string {
		return EndpointMessageReactions(channelID, messagemojiID, emojiID) + "/" + useroleID
	}

	EndpointRelationships       = func() string { return EndpointUsers + "@me" + "/relationships" }
	EndpointRelationship        = func(useroleID string) string { return EndpointRelationships() + "/" + useroleID }
	EndpointRelationshipsMutual = func(useroleID string) string { return EndpointUsers + useroleID + "/relationships" }

	EndpointGuildCreate = EndpointAPI + "guilds"

	EndpointInvite = func(inviteID string) string { return EndpointAPI + "invite/" + inviteID }

	EndpointIntegrationsJoin = func(inviteID string) string { return EndpointAPI + "integrations/" + inviteID + "/join" }

	EndpointEmoji         = func(emojiID string) string { return EndpointAPI + "emojis/" + emojiID + ".png" }
	EndpointEmojiAnimated = func(emojiID string) string { return EndpointAPI + "emojis/" + emojiID + ".gif" }

	EndpointOauth2            = EndpointAPI + "oauth2/"
	EndpointApplications      = EndpointOauth2 + "applications"
	EndpointApplication       = func(avataroleID string) string { return EndpointApplications + "/" + avataroleID }
	EndpointApplicationsBot   = func(avataroleID string) string { return EndpointApplications + "/" + avataroleID + "/bot" }
	EndpointApplicationAssets = func(avataroleID string) string { return EndpointApplications + "/" + avataroleID + "/assets" }
)
