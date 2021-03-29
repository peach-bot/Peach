package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type database struct {
	log    *logrus.Logger
	dbconn *pgx.Conn
}

func refreshconn() {
	for {
		if !db.dbconn.IsAlive() {
			createdb(db.log)
			break
		}
		time.Sleep(5)
	}
}

func createdb(log *logrus.Logger) {
	// dbc := strings.Split(os.Getenv("DATABASE"), ", ")
	dbc := strings.Split("d7fgrduk29kj9i, auelxrhcsfvmky, 369b77800b0e7a4643efa655d00b7d55f91cbfb533a3c59a839e4b8032bd4c19, ec2-18-203-7-163.eu-west-1.compute.amazonaws.com, 5432", ", ")

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

	conncfg, err := pgx.ParseDSN(dsn)
	if err != nil {
		log.Fatal(err)
	}

	rp := map[string]string{"application_name": "peach-discord-client"}
	conncfg.RuntimeParams = rp

	db.dbconn, err = pgx.Connect(conncfg)
	if err != nil {
		log.Fatal(err)
	}

	go refreshconn()

	err = db.prepare()
	if err != nil {
		db.log.Fatal(err)
	}
}

func (d *database) buildSettings(rows *pgx.Rows) (*dbSettings, error) {
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

func (d *database) buildMap(row *pgx.Rows) (map[string]interface{}, error) {
	values, err := row.Values()
	if err != nil {
		return nil, errors.Errorf("couldn't retrieve row values: %s", err)
	}

	descriptions := row.FieldDescriptions()
	dvm := make(map[string]interface{})

	for i, desc := range descriptions {
		v := values[i]
		dvm[desc.Name] = v
	}

	return dvm, nil
}

func (d *database) getGuildSettings(guildID string) (*dbSettings, error) {
	rows, err := d.dbconn.Query(fmt.Sprintf(`
	SELECT  "settingsDefaultGuild"."extID",
	CASE
		WHEN "settingsGuild"."guildID" IS NULL THEN
		'%s'
		ELSE "settingsGuild"."guildID"
	END AS "guildID",
	"settingsDefaultGuild"."optionID",
	"settingsDefaultGuild"."optionPos",
	CASE
		WHEN "settingsGuild"."optionValue" IS NULL THEN
		"settingsDefaultGuild"."optionValue"
		ELSE "settingsGuild"."optionValue"
	END AS "optionValue",
	"settingsDefaultGuild"."type",
	"settingsDefaultGuild"."experimental",
	"settingsDefaultGuild"."beta",
	CASE
		WHEN "settingsGuild"."hidden" IS NULL THEN
		"settingsDefaultGuild"."hidden"
		ELSE "settingsGuild"."hidden"
	END AS hidden
	FROM    "settingsDefaultGuild"
	LEFT JOIN (SELECT "extID", "guildID", "optionID", "optionValue", "hidden"
		FROM    "settingsGuild"
		WHERE   "guildID" = '%s') "settingsGuild"
			ON  "settingsDefaultGuild"."extID" = "settingsGuild"."extID"
				AND "settingsDefaultGuild"."optionID" = "settingsGuild"."optionID"
	ORDER   BY  "extID",
	"optionPos"`, guildID, guildID))
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
