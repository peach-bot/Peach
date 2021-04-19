package main

import "time"

type Extensions struct {
	Bot        *Client
	Spotify    extSpotify
	Moderation extModeration
	Internal   extInternal
	AliasMap   map[string]string
}

func (x *Extensions) setup(bot *Client, spotifyid, spotifysecret string) error {

	x.Bot = bot
	// configure command aliases
	x.AliasMap = map[string]string{
		"clear":   "clear",
		"c":       "clear",
		"about":   "about",
		"version": "about",
		"help":    "about",
	}

	// run extension setups
	x.Spotify.Setup(spotifyid, spotifysecret, x.Bot)
	x.Moderation.Setup(x.Bot)
	x.Internal.Setup(x.Bot)

	return nil
}

func (x *Extensions) runCommand(invoke string, args []string, ctx *EventMessageCreate) error {

	var err error

	command := x.AliasMap[invoke]

	switch command {
	case "clear":
		err = x.Moderation.Clear(ctx, args)
	case "about":
		err = x.Internal.About(ctx)
	default:
		err = nil
	}
	return err
}

func (x *Extensions) handleNoPermission(m *Message) error {
	sorry, err := x.Bot.SendMessage(m.ChannelID, NewMessage{":no_entry: It seems like you do not have the permissions to use this command.", false, nil})
	if err != nil {
		return err
	}

	time.Sleep(5 * time.Second)

	err = sorry.Delete(x.Bot)
	if err != nil {
		return err
	}
	err = m.Delete(x.Bot)
	if err != nil {
		return err
	}
	return nil
}
