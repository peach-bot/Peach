package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	ErrUnexpectedStatus = func(want, have int) error {
		return fmt.Errorf("Unexpected response status. Want: %s, Have: %s.", fmt.Sprint(want)+" "+http.StatusText(want), fmt.Sprint(have)+" "+http.StatusText(have))
	}
)

func addURLArg(query string, key string, value string) string {
	if query == "" {
		return fmt.Sprintf("?%s=%s", key, value)
	}
	return fmt.Sprintf("%s&%s=%s", query, key, value)
}

func (c *Client) request(method string, endpointURL string, routeid string, body []byte, attempt int) (*http.Response, []byte, error) {
	c.Log.Debugf("Sending %s request to %s. Attempt #%d", method, endpointURL, attempt)
	route := c.Ratelimiter.PrepareRoute(routeid)

	req, err := http.NewRequest(method, endpointURL, bytes.NewBuffer(body))
	if err != nil {
		route.Release(nil)
		return nil, nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bot %s", c.TOKEN))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		route.Release(nil)
		return nil, nil, err
	}
	defer resp.Body.Close()

	err = route.Release(resp.Header)
	if err != nil {
		return nil, nil, err
	}

	respbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return resp, respbody, nil
}

func (c *Client) Request(method string, endpointURL string, body interface{}) (*http.Response, []byte, error) {
	routeid := strings.SplitN(endpointURL, "?", 2)[0]

	var bodybytes []byte
	var err error
	if body != nil {
		bodybytes, err = json.Marshal(body)
		if err != nil {
			return nil, nil, err
		}
	}

	resp, respbody, err := c.request(method, endpointURL, routeid, bodybytes, 1)
	return resp, respbody, err
}
