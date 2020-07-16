package main

import (
	"os"
	"strings"
)

var tokens []string

func (c *clientCoordinator) gettokens() {
	tokens = strings.Split(os.Getenv("BOTTOKEN"), ", ")
}
