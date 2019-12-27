import psycopg2
import os
import json

class DatabaseHandler:

    def __init__(self, log):
        self.log = log
        self.dbconn = psycopg2.connect("host={0} port={1} user={2} password={3} dbname={4} application_name='peach - bot'".format(os.environ['DBHOST'], os.environ['DBPORT'], os.environ['DBUSER'], os.environ['DBPASSWORD'], os.environ['DBNAME'],))

        self.dbcur = self.dbconn.cursor()

    def fetch_settings(self, serverid):
        """Returns an epic dictionary containing a server's settings."""
        rawsettings = self.plugin_serverconfig_get(serverid)
        settings = {
            "plugins": {}
        }
        for setting in rawsettings:
            if setting[0] not in settings["plugins"]:
                settings["plugins"][setting[0]] = {}
            settings["plugins"][setting[0]][setting[1]] = {
                "type": setting[3]["type"],
                "value": setting[3]["value"],
            }
        return settings

    def plugin_serverconfig_get(self, serverid):
        """Grabs a server's settings from the database."""
        self.dbcur.execute("SELECT defaults.pluginid, defaults.cfgkey, serverconfig.serverid, CASE WHEN serverconfig.cfgvalue IS NULL THEN defaults.cfgvalue ELSE serverconfig.cfgvalue END AS result FROM defaults LEFT JOIN (SELECT * FROM serverconfig WHERE serverid = {}) serverconfig ON defaults.pluginid = serverconfig.pluginid AND defaults.cfgkey = serverconfig.cfgkey".format(serverid))
        data = self.dbcur.fetchall()
        return data
