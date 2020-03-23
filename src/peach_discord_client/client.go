package main

// Run starts various background routines and starts listeners
func (c *Client) Run() error {
	err := c.CreateWebsocket()
	if err != nil {
		return err
	}

	return nil
}

// CreateClient creates a new discord client
func CreateClient() (c *Client, err error) {

	c = &Client{Sequence: new(int64)}

	// Parse shard coordinator for gateway url and shardID
	c.GatewayURL = "wss://gateway.discord.gg/"
	c.GatewayURL = c.GatewayURL + "?v=" + APIVersion + "&encoding=json"

	return
}
