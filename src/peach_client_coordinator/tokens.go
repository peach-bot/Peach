package main

import "context"

var tokens []string

func (c *clientCoordinator) gettokens() {
	rows, err := db.dbconn.Query(context.Background(), "SELECT token FROM tokens ORDER BY priority ASC")
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
