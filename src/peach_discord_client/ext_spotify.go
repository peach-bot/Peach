package main

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/hako/durafmt"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

type extSpotify struct {
	Bot           *Client
	URLRegex      *regexp.Regexp
	ClientID      string
	ClientSecret  string
	SpotifyClient spotify.Client
}

func (e *extSpotify) Setup(clientid, clientsecret string, bot *Client) error {
	e.ClientID = clientid
	e.ClientSecret = clientsecret
	e.Bot = bot
	e.URLRegex = regexp.MustCompile(`https:\/\/open\.spotify\.com\/(album|track|artist)\/(\w{22})`)

	config := &clientcredentials.Config{
		ClientID:     clientid,
		ClientSecret: clientsecret,
		TokenURL:     spotify.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		return err
	}
	e.SpotifyClient = spotify.Authenticator{}.NewClient(token)

	return nil
}

func (e *extSpotify) OnMessage(ctx *Message) error {

	s := e.URLRegex.FindStringSubmatch(ctx.Content)
	if len(s) == 0 {
		return nil
	}

	spotifytype := s[1]
	spotifyid := spotify.ID(s[2])

	switch spotifytype {
	case "album":
		_, err := e.SpotifyClient.GetAlbum(spotifyid)
		if err != nil {
			return err
		}
	case "track":
		track, err := e.SpotifyClient.GetTrack(spotifyid)
		if err != nil {
			return err
		}

		_, err = e.Bot.SendMessage(ctx.ChannelID, NewMessage{
			Embed: Embed{
				Author: EmbedAuthor{
					Name:    "Spotify",
					IconURL: "https://assets.ifttt.com/images/channels/51464135/icons/large.png",
				},
				Thumbnail:   EmbedThumbnail{URL: track.Album.Images[0].URL},
				Color:       1947988,
				Title:       track.Name,
				URL:         track.ExternalURLs["spotify"],
				Description: "by " + track.Artists[0].Name,
				Fields: []*EmbedField{
					{
						Name:   "Release Date",
						Value:  track.Album.ReleaseDate,
						Inline: true,
					},
					{
						Name:   "Length",
						Value:  durafmt.Parse(time.Duration(track.Duration * int(time.Millisecond))).LimitFirstN(2).String(),
						Inline: true,
					},
					{
						Name:   "Album",
						Value:  fmt.Sprintf("[%s](%s)", track.Album.Name, track.Album.ExternalURLs["spotify"]),
						Inline: false,
					},
				},
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}
