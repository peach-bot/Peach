package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/patrickmn/go-cache"
)

func (c *Client) getGuild(guildid string) (*Guild, error) {

	// Send Request
	req, err := http.NewRequest("GET", EndpointGuild(guildid), *new(io.Reader))
	if err != nil {
		return nil, err
	}

	req = c.SetDefaultRequestHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Decode Body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}

	data := new(Guild)

	err = json.Unmarshal(bodyBytes, data)
	if err != nil {
		return nil, err
	}

	c.GuildCache.Set(guildid, *data, cache.DefaultExpiration)

	return data, nil
}
