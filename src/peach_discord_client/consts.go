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
	_                                       // 6 is not a opcode
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
