package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type cfgOption struct {
	OptionValue  string `json:"option_value"`
	Type         string `json:"type"`
	Experimental bool   `json:"experimental"`
	Beta         bool   `json:"beta"`
	Hidden       bool   `json:"hidden"`
}

type cfgExtension struct {
	Options map[string]cfgOption `json:"options"`
}

type cfgSettings struct {
	Extensions map[string]cfgExtension `json:"extensions"`
}

func (c *Client) getSetting(guildID, extID, optionID string) string {
	return c.Settings[guildID].Extensions[extID].Options[optionID].OptionValue
}

func (c *Client) getGuildSettings(guildID string) error {
	tempClient := &http.Client{}

	//idk why, where and how but somehow some guildIDs are fucked and have a %s infront of them
	//so we unfuck em
	if strings.HasPrefix(guildID, "%s") {
		guildID = fmt.Sprintf(guildID, "")
	}

	req, err := http.NewRequest("GET", c.ClientCoordinatorURL+"guilds/"+guildID, nil)
	if err != nil {
		return err
	}
	req = setCCRequestHeaders(c, req)
	resp, err := tempClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		c.Log.Errorf("Websocket received unexpected response from client coordinator. Expected Status 200 OK got %s instead", resp.Status)
	}

	settings := cfgSettings{}
	err = json.NewDecoder(resp.Body).Decode(&settings)
	if err != nil {
		return err
	}

	//cache the settings
	c.Settings[guildID] = settings

	return nil
}
