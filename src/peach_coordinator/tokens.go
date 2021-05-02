package main

import "context"

var tokens []string

func (c *Coordinator) gettokens() {
	rows, err := c.DB.DBConn.Query(context.Background(), QueryTokens)
	if err != nil {
		c.log.Fatal(err)
	}
	defer rows.Close()
	tokens = []string{}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			c.log.Fatal(err)
		}
		token := values[0].(string)
		tokens = append(tokens, token)
	}
}
