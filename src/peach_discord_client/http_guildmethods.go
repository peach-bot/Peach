package main

import (
	"encoding/json"

	"github.com/patrickmn/go-cache"
)

func (c *Client) getGuild(guildid string) (*Guild, error) {

	_, body, err := c.Request("GET", EndpointGuild(guildid), nil)
	if err != nil {
		return nil, err
	}

	data := new(Guild)

	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	c.GuildCache.Set(guildid, *data, cache.DefaultExpiration)

	return data, nil
}
