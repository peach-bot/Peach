import os
import json

import psycopg2


class DatabaseHandler:

    def __init__(self, bot):
        self.bot = bot
        self.dbconn = psycopg2.connect('host={0} user={1} password={2} dbname=peach'.format(os.environ["DBHOST"], os.environ["DBUSER"], os.environ["DBPASSWORD"]))

        self.dbcur = self.dbconn.cursor()

    async def update_servers(self):
        self.bot.log.info("Updating server database")
        self.dbcur.execute('SELECT * FROM servers')
        results = self.dbcur.fetchall()
        botserverids = []
        dbserverids= [] 
        for server in self.bot.guilds:
            botserverids.append(server.id)
        for server in results:
            dbserverids.append(server[0])
        for serverid in botserverids:
            if serverid not in dbserverids:
                server = self.bot.get_guild(serverid)
                self.dbcur.execute('INSERT INTO servers VALUES ({0.id}, {0.owner_id})'.format(self.bot.get_guild(serverid)))
                self.bot.log.info("Added {0} to server database".format(server.name))
        self.dbconn.commit()
        self.bot.log.info("Updating server database complete")

    async def create_user(self, userid):
        self.dbcur.execute("INSERT INTO users VALUES({0})".format(userid))
        self.dbconn.commit()
        self.bot.log.info("Added {0}#{1} to user database".format(message.author.name, message.author.discriminator))

    async def plugin_getuser(self, userid, pluginname):
        self.dbcur.execute("SELECT {0} FROM users WHERE id = {1}".format("plugin_"+pluginname, userid))
        return json.loads(self.dbcur.fetchall()[0][0])
    
    async def plugin_updateuser(self, userid, pluginname, newdata):
        self.dbcur.execute("SELECT * FROM users WHERE id = {0}".format(userid))
        if self.dbcur.fetchall() == []:
            await self.create_user(userid)
        self.dbcur.execute("UPDATE users SET {0} = '{1}' WHERE id = {2}".format("plugin_"+pluginname, json.dumps(newdata), userid))
        self.dbconn.commit()

    async def plugin_getserver(self, serverid, pluginname):
        self.dbcur.execute("SELECT {0} FROM servers WHERE id = {1}".format("plugin_"+pluginname, serverid))
        return self.dbcur.fetchall()
    
    async def plugin_updateserver(self, serverid, pluginname, newdata):
        self.dbcur.execute("UPDATE servers SET {0} = '{1}' WHERE id = {2}".format("plugin_"+pluginname, json.dumps(newdata), userid))
        self.dbconn.commit()