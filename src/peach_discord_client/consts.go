package main

//go:generate stringer -type opcode,closecode,activitytype -output consts_generated.go

type opcode int

// Gateway opcodes, denote payload type, see https://discordapp.com/developers/docs/topics/opcodes-and-status-codes#gateway-opcodes
const (
	opcodeDispatch            opcode = iota // Receive      | An event was dispatched.
	opcodeHeartbeat                         // Send/Receive | Fired periodically by the client to keep the connection alive.
	opcodeIdentify                          // Send         | Starts a new session during the initial handshake.
	opcodePresenceUpdate                    // Send         | Update the client's presence.
	opcodeVoiceStateUpdate                  // Send         | Used to join/leave or move between voice channels.
	_                                       // 5 is not a opcode
	opcodeResume                            // Send         | Resume a previous session that was disconnected.
	opcodeReconnect                         // Receive      | You must reconnect with a new session immediately.
	opcodeRequestGuildMembers               // Send         | Request information about offline guild members in a large guild.
	opcodeInvalidSession                    // Receive      | The session has been invalidated. You should reconnect and identify/resume accordingly.
	opcodeHello                             // Receive      | Sent immediately after connecting, contains the heartbeat_interval to use.
	opcodeHeartbeatACK                      // Receive      | Sent in response to receiving a heartbeat to acknowledge that it has been received.
)

type closecode int

// Gateway Close Event Codes, denote reason for gateway closure, see https://discordapp.com/developers/docs/topics/opcodes-and-status-codes#gateway-opcodes
const (
	closecodeUnknownError         closecode = iota + 4000 // Not sure what went wrong. Try reconnecting.
	closecodeUnknownOpCode                                // Sent invalid opcode or invalid payload for opcode.
	closecodeDecodeError                                  // Sent invalid payload.
	closecodeNotAuthenticated                             // Sent payload prior to identifying.
	closecodeAuthenticationFailed                         // Account token in identify payload is incorrect.
	closecodeAlreadyAuthenticated                         // Sent more than one identify payload.
	_                                                     // 4006 in not a closecode
	closecodeInvalidSquence                               // Sent invalid sequence when resuming.
	closecodeRateLimited                                  // Sending payloads to quickly.
	closecodeSessionTimedOut                              // Session timed out. Reconnect or start new session.
	closecodeInvalidShard                                 // Sent invalid shard in identify payload.
	closecodeShardingRequired                             // Sharding required because bot is in too many guilds.
	closecodeInvalidAPIVersion                            // Sent an invalid gateway version.
	closecodeInvalidIntents                               // Sent invalid gateway intent.
	closecodeDisallowedIntents                            // Sent intent the account isn't eligible for.
)

const closecodeReconnect closecode = 1001

type activitytype int

const (
	activitytypeGame activitytype = iota
	activitytypeStreaming
	activitytypeListening
	_
	activitytypeCustom
)

type messagetype int

const (
	messagetypeDefault messagetype = iota
	messagetypeRecipientAdd
	messagetypeRecipientRemove
	messagetypeCall
	messagetypeChannelNameChange
	messagetypeChannelIconChange
	messagetypeChannelPinnedMessage
	messagetypeGuildMemberJoin
	messagetypeNitroBoost
	messagetypeNitroBoostTier1
	messagetypeNitroBoostTier2
	messagetypeNitroBoostTier3
	messagetypeChannelFollowAdd
	_
	messagetypeGuildDiscoveryDisqualified
	messagetypeGuildDiscoveryRequalified
)

const (
	permissionCreateInstantInvite = 0x00000001
	permissionKickMembers         = 0x00000002
	permissionBanMembers          = 0x00000004
	permissionAdministrator       = 0x00000008
	permissionManageChannels      = 0x00000010
	permissionManageGuild         = 0x00000020
	permissionAddReactions        = 0x00000040
	permissionViewAuditLog        = 0x00000080
	permissionPrioritySpeaker     = 0x00000100
	permissionStream              = 0x00000200
	permissionViewChannel         = 0x00000400
	permissionSendMessages        = 0x00000800
	permissionSendTtsMessages     = 0x00001000
	permissionManageMessages      = 0x00002000
	permissionEmbedLinks          = 0x00004000
	permissionAttachFiles         = 0x00008000
	permissionReadMessageHistory  = 0x00010000
	permissionMentionEveryone     = 0x00020000
	permissionUseExternalEmojis   = 0x00040000
	permissionViewGuildInsights   = 0x00080000
	permissionConnect             = 0x00100000
	permissionSpeak               = 0x00200000
	permissionMuteMembers         = 0x00400000
	permissionDeafenMembers       = 0x00800000
	permissionMoveMembers         = 0x01000000
	permissionUseVad              = 0x02000000
	permissionChangeNickname      = 0x04000000
	permissionManageNicknames     = 0x08000000
	permissionManageRoles         = 0x10000000
	permissionManageWebhooks      = 0x20000000
	permissionManageEmojis        = 0x40000000
)

const (
	intentGuilds uint16 = 1 << iota
	intentGuildMembers
	intentGuildBans
	intentGuildEmojis
	intentGuildIntegrations
	intentGuildWebhooks
	intentGuildInvites
	intentGuildVoiceStates
	intentGuildPresences
	intentGuildMessages
	intentGuildMessageReactions
	intentGuildMessageTyping
	intentDirectMessages
	intentDirectMessageReactions
	intentDirectMessageTyping
)
