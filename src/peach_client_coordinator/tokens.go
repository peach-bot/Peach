package main

var tokens []string

func (c *clientCoordinator) gettokens() {
	tokenrows, err := db.dbconn.Query("SELECT * FROM tokens ORDER BY priority ASC")
	if err != nil {
		c.log.Error(err)
	}
	tokens = []string{}
	for {
		outofrange := tokenrows.Next()
		if !outofrange {
			break
		}
		values, err := tokenrows.Values()
		if err != nil {
			c.log.Error(err)
		}
		token := values[0].(string)
		tokens = append(tokens, token)
	}
}
