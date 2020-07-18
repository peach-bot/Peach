package main

import "github.com/jackc/pgx"

type database struct {
	dbconn *pgx.Conn
}
