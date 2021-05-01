package main

import (
	"context"
	"strings"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type database struct {
	log    *logrus.Logger
	dbconn *pgx.Conn
}

func refreshconn(dbcstring string) {
	for {
		if db.dbconn.IsClosed() {
			createdb(db.log, dbcstring)
			break
		}
		time.Sleep(5)
	}
}

func createdb(log *logrus.Logger, dbcstring string) {
	// dbc := strings.Split(os.Getenv("DATABASE"), ", ")
	dbc := strings.Split(dbcstring, ", ")

	db = database{log, nil}

	dsn := ""
	dbname := dbc[0]
	dsn = dsn + "dbname=" + dbname
	user := dbc[1]
	dsn = dsn + " user=" + user
	password := dbc[2]
	dsn = dsn + " password=" + password
	host := dbc[3]
	dsn = dsn + " host=" + host
	port := dbc[4]
	dsn = dsn + " port=" + port
	dsn = dsn + " sslmode=require"

	conncfg, err := pgx.ParseConfig(dsn)
	if err != nil {
		log.Fatal(err)
	}

	rp := map[string]string{"application_name": "peach-discord-client"}
	conncfg.RuntimeParams = rp

	db.dbconn, err = pgx.ConnectConfig(context.Background(), conncfg)
	if err != nil {
		log.Fatal(err)
	}

	go refreshconn(dbcstring)

	err = db.prepare()
	if err != nil {
		db.log.Fatal(err)
	}
}

func (d *database) buildSettings(rows pgx.Rows) (*dbSettings, error) {
	settings := dbSettings{Extensions: map[string]dbExtension{}}
	for rows.Next() {
		values, err := d.buildMap(rows)
		if err != nil {
			return nil, err
		}

		optn := dbOption{
			OptionValue:  values["optionValue"].(string),
			Type:         values["type"].(string),
			Experimental: values["experimental"].(bool),
			Beta:         values["beta"].(bool),
			Hidden:       values["hidden"].(bool),
		}

		if ext, ok := settings.Extensions[values["extID"].(string)]; ok {
			ext.Options[values["optionID"].(string)] = optn
		} else {
			newExtension := dbExtension{Options: map[string]dbOption{values["optionID"].(string): optn}}
			settings.Extensions[values["extID"].(string)] = newExtension
		}
	}
	return &settings, nil
}

func (d *database) buildMap(row pgx.Rows) (map[string]interface{}, error) {
	values, err := row.Values()
	if err != nil {
		return nil, errors.Errorf("couldn't retrieve row values: %s", err)
	}

	descriptions := row.FieldDescriptions()
	dvm := make(map[string]interface{})

	for i, desc := range descriptions {
		v := values[i]
		dvm[string(desc.Name)] = v
	}

	return dvm, nil
}

func (d *database) getGuildSettings(guildID string) (*dbSettings, error) {
	rows, err := d.dbconn.Query(context.Background(), QueryGuildSettings(guildID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	settings, err := d.buildSettings(rows)
	if err != nil {
		return nil, errors.Errorf("building settings failed: %s", err)
	}

	return settings, nil
}

func (d *database) getUserSettings(userID string) (*dbSettings, error) {
	rows, err := d.dbconn.Query(context.Background(), QueryUserSettings(userID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	settings, err := d.buildSettings(rows)
	if err != nil {
		return nil, errors.Errorf("building settings failed: %s", err)
	}

	return settings, nil
}
