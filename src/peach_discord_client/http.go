package main

import (
	"fmt"
	"net/http"
)

// SetDefaultRequestHeaders adds authorization and content type to request header
func (c *Client) SetDefaultRequestHeaders(req *http.Request) *http.Request {
	req.Header.Add("Authorization", fmt.Sprintf("Bot %s", c.TOKEN))
	req.Header.Add("Content-Type", "application/json")
	return req
}

func addURLArg(query string, key string, value string) string {
	if query == "" {
		return fmt.Sprintf("?%s=%s", key, value)
	}
	return fmt.Sprintf("%s&%s=%s", query, key, value)
}
