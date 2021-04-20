package main

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/hako/durafmt"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type extSpotify struct {
	Bot           *Client
	URLRegex      *regexp.Regexp
	ClientID      string
	ClientSecret  string
	SpotifyClient spotify.Client
	Responses     map[string]string
	sync.Mutex
}

func (e *extSpotify) Setup(clientid, clientsecret string, bot *Client) error {

	// constructor
	e.ClientID = clientid
	e.ClientSecret = clientsecret
	e.Bot = bot
	e.Responses = make(map[string]string)
	e.URLRegex = regexp.MustCompile(`https:\/\/open\.spotify\.com\/(album|track|artist)\/(\w{22})`)

	go e.RefreshToken()

	return nil
}

func (e *extSpotify) RefreshToken() {
	ticker := time.NewTicker(time.Hour)
	for {
		// Authorize spotify client
		config := &clientcredentials.Config{
			ClientID:     e.ClientID,
			ClientSecret: e.ClientSecret,
			TokenURL:     spotify.TokenURL,
		}
		token, err := config.Token(context.Background())
		if err != nil {
			e.Bot.Log.Fatal(err)
		}
		e.SpotifyClient = spotify.Authenticator{}.NewClient(token)
		e.Bot.Log.Debug("Refreshed Spotify token")

		// Loop
		select {
		case <-ticker.C:
		case <-e.Bot.Connected:
		}
	}
}

func (e *extSpotify) OnMessage(ctx *Message) error {

	// See if message contains spotify url
	s := e.URLRegex.FindStringSubmatch(ctx.Content)
	if len(s) == 0 {
		return nil
	}

	// Suppress discord embed
	flags := new(int)
	*flags = MessageFlagSuppressEmbeds
	_, err := e.Bot.EditMessage(ctx.ChannelID, ctx.ID, EditMessageArgs{
		Flags:           flags,
		Content:         "",
		Embed:           nil,
		AllowedMentions: nil,
	})

	// extract information
	spotifytype := s[1]
	spotifyid := spotify.ID(s[2])

	var msg *Message = nil

	// send message with embed for type and id
	switch spotifytype {
	case "album":
		album, err := e.SpotifyClient.GetAlbum(spotifyid)
		if err != nil {
			return err
		}

		var playtime int
		for _, track := range album.Tracks.Tracks {
			playtime += track.Duration
		}

		msg, err = e.Bot.SendMessage(ctx.ChannelID, NewMessage{
			Embed: Embed{
				Author: EmbedAuthor{
					Name:    "Spotify",
					IconURL: "https://assets.ifttt.com/images/channels/51464135/icons/large.png",
				},
				Thumbnail:   EmbedThumbnail{URL: album.Images[0].URL},
				Color:       1947988,
				Title:       album.Name,
				URL:         album.ExternalURLs["spotify"],
				Description: "by " + album.Artists[0].Name,
				Fields: []*EmbedField{
					{
						Name:   "Release Date",
						Value:  album.ReleaseDate,
						Inline: true,
					},
					{
						Name:   "Tracks",
						Value:  fmt.Sprint(album.Tracks.Total),
						Inline: true,
					},
					{
						Name:   "Length",
						Value:  durafmt.Parse(time.Duration(playtime * int(time.Millisecond))).LimitFirstN(2).String(),
						Inline: false,
					},
				},
			},
		})
		if err != nil {
			return err
		}

	case "track":
		track, err := e.SpotifyClient.GetTrack(spotifyid)
		if err != nil {
			return err
		}

		msg, err = e.Bot.SendMessage(ctx.ChannelID, NewMessage{
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

	case "artist":
		artist, err := e.SpotifyClient.GetArtist(spotifyid)
		if err != nil {
			return err
		}

		p := message.NewPrinter(language.English)

		var genreLabel string = "Genre"
		var genre string = "None"
		if len(artist.Genres) > 0 {
			if len(artist.Genres) >= 3 {
				genreLabel = "Genres"
				genre = fmt.Sprintf("%s, %s, %s", artist.Genres[0], artist.Genres[1], artist.Genres[2])
			} else {
				genre = fmt.Sprintf("%s", artist.Genres[0])
			}
		}

		popularity := artist.Popularity / 20

		msg, err = e.Bot.SendMessage(ctx.ChannelID, NewMessage{
			Embed: Embed{
				Author: EmbedAuthor{
					Name:    "Spotify",
					IconURL: "https://assets.ifttt.com/images/channels/51464135/icons/large.png",
				},
				Thumbnail: EmbedThumbnail{URL: artist.Images[0].URL},
				Color:     1947988,
				Title:     artist.Name,
				URL:       artist.ExternalURLs["spotify"],
				Fields: []*EmbedField{
					{
						Name:   "Followers",
						Value:  p.Sprint(artist.Followers.Count),
						Inline: true,
					},
					{
						Name:   "Popularity",
						Value:  strings.Repeat("★", popularity) + strings.Repeat("☆", 5-popularity),
						Inline: true,
					},
					{
						Name:   genreLabel,
						Value:  genre,
						Inline: false,
					},
				},
			},
		})
		if err != nil {
			return err
		}
	}

	if msg == nil {
		return nil
	}

	// Add delete reaction
	err = e.Bot.CreateReaction(msg.ChannelID, msg.ID, "🗑", nil)
	if err != nil {
		return err
	}

	// Add response to cache
	e.Lock()
	e.Responses[msg.ID] = ctx.Author.ID
	e.Unlock()

	return nil
}

func (e *extSpotify) OnReact(ctx *EventMessageReactionAdd) {
	e.Bot.Log.Debug(ctx.Emoji.Name)
	if userID, ok := e.Responses[ctx.MessageID]; ok {
		if ctx.UserID == userID && ctx.Emoji.Name == "🗑" {
			e.Bot.DeleteMessage(ctx.ChannelID, ctx.MessageID)
			e.Lock()
			delete(e.Responses, ctx.MessageID)
			e.Unlock()
		}
	}

}
