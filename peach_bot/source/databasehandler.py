import os
import json

import psycopg2

class DatabaseHandler:

    def __init__(self, bot):
        self.bot = bot
        self.dbconn = psycopg2.connect('host={0} user=peach password=peach dbname=peach'.format(os.environ["DBHOST"]))

        self.dbcur = self.dbconn.cursor()

    async def update_servers(self):
        self.bot.log.info("Updating server database")
        botservers = []
        for server in self.bot.guilds:
            botservers.append(server)
        for server in botservers:
            ownerid = server.owner.id
            self.dbcur.execute('INSERT INTO users VALUES ({0}) ON CONFLICT DO NOTHING'.format(ownerid))
            self.dbconn.commit()
            self.dbcur.execute('INSERT INTO servers VALUES ({0.id}, {0.owner_id}) ON CONFLICT DO NOTHING'.format(server))
            self.bot.log.info("Added {0} to server database".format(server.name))
        self.dbconn.commit()
        self.bot.log.info("Updating server database complete")

    async def create_user(self, userid):
        self.dbcur.execute("INSERT INTO users VALUES ({0}) ON CONFLICT DO NOTHING".format(userid))
        self.dbconn.commit()
        user = self.bot.get_user(userid)
        self.bot.log.info("Added {0}#{1} to user database".format(user.name, user.discriminator))

    async def plugin_getuserserverconfig(self, serverid, userid, pluginid, cfgkey):
        self.dbcur.execute("SELECT cfgvalue FROM userserverconfig WHERE serverid = {0} AND userid = {1} AND pluginid = {2} AND cfgkey = {3}".format(serverid, userid, pluginid, cfgkey))
        return json.loads(self.dbcur.fetchall()[0][0])
    
    async def plugin_updateuserserverconfig(self, serverid, userid, pluginid, cfgkey, cfgvalue):
        await self.create_user(userid)
        self.dbcur.execute("INSERT INTO userserverconfig VALUES ({0}, {1}, {2}, {3}, '{4}') ON CONFLICT DO UPDATE SET cfgvalue = '{4}' WHERE serverid = {0} AND userid = {1} AND pluginid = {2} AND cfgkey = {3}".format(serverid, userid, pluginid, cfgkey, json.dumps(cfgvalue)))
        self.dbconn.commit()

    async def plugin_getuserglobalconfig(self, userid, pluginid, cfgkey):
        self.dbcur.execute("SELECT cfgvalue FROM userglobalconfig WHERE userid = {0} AND pluginid = '{1}' AND cfgkey = '{2}'".format(userid, pluginid, cfgkey))
        data = self.dbcur.fetchall()[0][0]
        return data

    async def plugin_updateuserglobalconfig(self, userid, pluginid, cfgkey, cfgvalue):
        await self.create_user(userid)
        self.dbcur.execute("INSERT INTO userglobalconfig VALUES ({0}, '{1}', '{2}', '{3}') ON CONFLICT (userid, pluginid, cfgkey) DO UPDATE SET cfgvalue = '{3}' WHERE userglobalconfig.userid = {0} AND userglobalconfig.pluginid = '{1}' AND userglobalconfig.cfgkey = '{2}'".format(userid, pluginid, cfgkey, json.dumps(cfgvalue)))
        self.dbconn.commit()

    async def plugin_getserverconfig(self, serverid, pluginid, cfgkey):
        self.dbcur.execute("SELECT cfgvalue FROM serverconfig WHERE serverid = {0} AND pluginid = {1} AND cfgkey = {2}".format(serverid, pluginid, cfgkey))
        return json.loads(self.dbcur.fetchall()[0][0])
    
    async def plugin_updateserverconfig(self, serverid, pluginid, cfgkey, cfgvalue):
        await self.create_user(userid)
        self.dbcur.execute("INSERT INTO serverconfig VALUES ({0}, {1}, {2}, '{3}') ON CONFLICT DO UPDATE SET cfgvalue = '{3}' WHERE serverid = {0} AND pluginid = {1} AND cfgkey = {2}".format(serverid, pluginid, cfgkey, json.dumps(cfgvalue)))
        self.dbconn.commit()

    async def plugin_getglobalconfig(self, pluginid, cfgkey):
        self.dbcur.execute("SELECT cfgvalue FROM globalconfig WHERE AND pluginid = '{0}' AND cfgkey = '{1}'".format(pluginid, cfgkey))
        data = self.dbcur.fetchall()[0][0]
        return data

    async def plugin_updateglobalconfig(self, pluginid, cfgkey, cfgvalue):
        self.dbcur.execute("INSERT INTO globalconfig VALUES ('{0}', '{1}', '{2}') ON CONFLICT (pluginid, cfgkey) DO UPDATE SET cfgvalue = '{2}' WHERE globalconfig.pluginid = '{0}' AND globalconfig.cfgkey = '{1}'".format(pluginid, cfgkey, json.dumps(cfgvalue)))
        self.dbconn.commit()