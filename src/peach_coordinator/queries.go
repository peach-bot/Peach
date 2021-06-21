package main

var (
	QueryGuildSettings = func(guildID string) string {
		return `
		SELECT "settingsDefaultGuild"."extID",
				CASE
				WHEN "settingsGuild"."guildID" IS NULL THEN '` + guildID + `'
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
		FROM   "settingsDefaultGuild"
				LEFT JOIN (SELECT "extID",
								"guildID",
								"optionID",
								"optionValue",
								"hidden"
						FROM   "settingsGuild"
						WHERE  "guildID" = '` + guildID + `') "settingsGuild"
					ON "settingsDefaultGuild"."extID" = "settingsGuild"."extID"
						AND "settingsDefaultGuild"."optionID" =
							"settingsGuild"."optionID"
		ORDER  BY "extID",
				"optionPos"`
	}
	QueryUserSettings = func(userID string) string {
		return `
		SELECT "settingsDefaultUser"."extID",
			CASE
				WHEN "settingsUser"."userID" IS NULL THEN '` + userID + `'
				ELSE "settingsUser"."userID"
			END AS "userID",
			"settingsDefaultUser"."optionID",
			"settingsDefaultUser"."optionPos",
			CASE
				WHEN "settingsUser"."optionValue" IS NULL THEN
				"settingsDefaultUser"."optionValue"
				ELSE "settingsUser"."optionValue"
			END AS "optionValue",
			"settingsDefaultUser"."type",
			"settingsDefaultUser"."experimental",
			"settingsDefaultUser"."beta",
			CASE
				WHEN "settingsUser"."hidden" IS NULL THEN
				"settingsDefaultUser"."hidden"
				ELSE "settingsUser"."hidden"
			END AS hidden
		FROM   "settingsDefaultUser"
			LEFT JOIN (SELECT "extID",
								"userID",
								"optionID",
								"optionValue",
								"hidden"
						FROM   "settingsUser"
						WHERE  "userID" = '` + userID + `') "settingsUser"
					ON "settingsDefaultUser"."extID" = "settingsUser"."extID"
						AND "settingsDefaultUser"."optionID" =
							"settingsUser"."optionID"
		ORDER  BY "extID",
				"optionPos" `
	}
	QueryTokens = "SELECT token FROM tokens ORDER BY priority ASC"
)
