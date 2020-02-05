import os
import re
from ast import literal_eval
import json

import psycopg2

RE_FREE_DOUBLE_QUOTES = re.compile(r"""(?<!\\)(?P<slashes>\\\\)*(?=\")""")

def _sanitize_params(*params):
    """
    Sanitizes all parameters passed in as keywords, returns tuple of strings
    that are safe to be fed into database query.
    """
    out = []
    for param in params:
        param = param.replace("\\", "\\\\")
        param = param.replace('"', '\\"')
        param = param.replace("'", "\\'")
        param = RE_FREE_DOUBLE_QUOTES.sub(r"\g<slashes>\\", param)
        out.append('"' + param + '"')
    return tuple(out)

class DatabaseHandler:
    def __init__(self, bot):
        self.bot = bot
        self.dbconn = psycopg2.connect("""
            host={} port={} user={} password={} dbname={} 
            application_name='peach - bot'""".format(
                os.environ['DBHOST'], os.environ['DBPORT'], os.environ['DBUSER'],
                os.environ['DBPASSWORD'], os.environ['DBNAME'],))
        self.dbcur = self.dbconn.cursor()

    async def update_servers(self):
        self.bot.log.info("Updating server database")
        botservers = []
        for server in self.bot.guilds:
            botservers.append(server)
        for server in botservers:
            self.dbcur.execute(
                "INSERT INTO users VALUES ({0}) ON CONFLICT"
                "DO NOTHING".format(*_sanitize_params(server.owner.id)))
            self.dbconn.commit()
            self.dbcur.execute(
                "INSERT INTO servers VALUES ({0}, {1}, {2}) ON "
                "CONFLICT (id) DO UPDATE SET name = {2} WHERE "
                "servers.id = {0}".format(
                    *_sanitize_params(server.id, server.owner.id, server.name)
                )
            )
        self.dbconn.commit()
        self.bot.log.info("Updating server database complete")

    async def create_user(self, userid):
        """Adds a userid to the database."""
        self.dbcur.execute(
            "INSERT INTO users VALUES ({}) ON "
            "CONFLICT DO NOTHING".format(*_sanitize_params(userid)))
        self.dbconn.commit()
        user = self.bot.get_user(userid)
        self.bot.log.info("Added {0}#{1} to user database".format(
                user.name, user.discriminator
            )
        )

    async def plugin_userserverconfig_get(self, serverid, userid,
            pluginid, cfgkey):
        """Retrieves a config value for a given key."""
        self.dbcur.execute(
            "SELECT cfgvalue FROM userserverconfig WHERE serverid = {} "
            "AND userid = {} AND pluginid = {} AND cfgkey = {}".format(
                *_sanitize_params(serverid, userid, pluginid, cfgkey)
            )
        )
        return json.loads(self.dbcur.fetchall()[0][0])

    async def plugin_userserverconfig_update(self, serverid, userid, pluginid,
            cfgkey, cfgvalue):
        """Updates config key."""
        await self.create_user(userid)
        self.dbcur.execute(
            "INSERT INTO userserverconfig VALUES ({0}, {1}, {2}, {3}, {4}) "
            "ON CONFLICT DO UPDATE SET cfgvalue = {4} WHERE serverid = {0} "
            "AND userid = {1} AND pluginid = {2} AND cfgkey = {3}".format(
                *_sanitize_params(serverid, userid, pluginid,
                    cfgkey, json.dumps(cfgvalue)
                )
            )
        )
        self.dbconn.commit()

    async def plugin_userglobalconfig_get(self, userid, pluginid, cfgkey):
        """Retrieves a config value for a given key."""
        self.dbcur.execute(
            "SELECT cfgvalue FROM userglobalconfig WHERE "
            "userid = {} AND pluginid = {} AND cfgkey = {}".format(
                *_sanitize_params(userid, pluginid, cfgkey)
            )
        )
        data = self.dbcur.fetchall()[0][0]
        return data

    async def plugin_userglobalconfig_update(self, userid, pluginid, cfgkey,
            cfgvalue):
        """Updates config key."""
        await self.create_user(userid)
        self.dbcur.execute(
            "INSERT INTO userglobalconfig VALUES "
            "({0}, {1}, {2}, {3}) ON CONFLICT (userid, pluginid, cfgkey) "
            "DO UPDATE SET cfgvalue = {3} WHERE userglobalconfig.userid = {0} "
            "AND userglobalconfig.pluginid = {1} AND "
            "userglobalconfig.cfgkey = {2}".format(
                *_sanitize_params(userid, pluginid, cfgkey,
                    json.dumps(cfgvalue)
                )
             )
        )
        self.dbconn.commit()

    async def plugin_serverconfig_get(self, serverid, pluginid, cfgkey):
        """Retrieves a config value for a given key."""
        self.dbcur.execute(
            "SELECT cfgvalue FROM serverconfig WHERE "
            "serverid = {0} AND pluginid = {1} AND cfgkey = {2}".format(
                *_sanitize_params(serverid, pluginid, cfgkey)
            )
        )
        value = await self.solvevalue(self.dbcur.fetchall()[0][0])
        return value
    
    async def plugin_defaults_update(self, pluginid, cfgkey, cfgvalue,
            position):
        """Adds new config keys without overwriting already existing ones."""
        self.dbcur.execute(
            "INSERT INTO defaults VALUES ({0}, {1}, {2}, {3}) ON CONFLICT "
            "(pluginid, cfgkey) DO UPDATE SET cfgvalue = {2}, "
            "position = {3} WHERE defaults.pluginid = {0} AND "
            "defaults.cfgkey = {1}".format(
                *_sanitize_params(
                    pluginid, cfgkey, json.dumps(cfgvalue), position
                )
            )
        )
        self.dbconn.commit()

    async def plugin_serverconfig_update(self, serverid, pluginid,
            cfgkey, cfgvalue):
        """Updates config key."""
        self.dbcur.execute(
            "INSERT INTO serverconfig VALUES ({0}, {1}, {2}, {3}) ON CONFLICT "
            "DO UPDATE SET cfgvalue = {3} WHERE serverid = {0} AND "
            "pluginid = {1} AND cfgkey = {2}".format(
                    *_sanitize_params(
                        serverid, pluginid, cfgkey, json.dumps(cfgvalue)
                )
            )
        )
        self.dbconn.commit()

    async def plugin_globalconfig_get(self, pluginid, cfgkey):
        """Retrieves a config value for a given key."""
        self.dbcur.execute(
            "SELECT cfgvalue FROM globalconfig WHERE AND pluginid = {} "
            "AND cfgkey = {}".format(
                *_sanitize_params(pluginid, cfgkey)
            )
        )
        data = self.dbcur.fetchall()[0][0]
        return data

    async def plugin_globalconfig_update(self, pluginid, cfgkey,
            cfgvalue):
        """Updates config key."""
        self.dbcur.execute(
            "INSERT INTO globalconfig VALUES ({0}, {1}, {2}) ON CONFLICT "
            "(pluginid, cfgkey) DO UPDATE SET cfgvalue = {2} WHERE "
            "globalconfig.pluginid = {0} AND globalconfig.cfgkey = {1}".format(
                    *_sanitize_params(pluginid, cfgkey, json.dumps(cfgvalue)
                )
            )
        )
        self.dbconn.commit()

    async def query(self, query):
        """
        Runs a custom database query. POST ONLY! Does not return anything.
        To fetch data use query_return(query).
        """
        self.bot.log.info("Running custom post query: {}".format(query))
        self.dbcur.execute(query)
        self.dbconn.commit()

    async def query_return(self, query):
        """
        Runs a custom database query. FETCH ONLY! Returns data.
        To post data use query(query).
        """
        self.bot.log.info("Running cutom fetch query: {}".format(query))
        self.dbcur.execute(query)
        return self.dbcur.fetchall()

    async def solvevalue(self, data):
        """Returns the value stored inside a database set."""
        if data["type"] == "bool":
            if data["value"] == "true":
                return True
            elif data["value"] == "false":
                return False
            else:
                self.bot.log.error("Could not solve database value, "
                    "boolean by type but not boolean by value.")

    async def increment_messages(self, channelid: int,
            timestamp: int, serverid: int):
        """Increments the message counter in the database."""
        self.dbcur.execute(
            "INSERT INTO channelstats VALUES ({0}, {1}, {2}, 1) ON CONFLICT "
            "(channelid, unixtimestamp) DO UPDATE SET "
            "messages = channelstats.messages + 1 WHERE "
            "channelstats.unixtimestamp = {1} AND "
            "channelstats.serverid = {2}".format(
                *_sanitize_params(channelid, timestamp, serverid)
            )
        )
        self.dbconn.commit()
