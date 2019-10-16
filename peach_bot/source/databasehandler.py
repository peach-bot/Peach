import os

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
                self.dbcur.execute('INSERT INTO servers VALUES ({0.id}, {0.owner_id})'.format(self.bot.get_guild(serverid)))
        self.dbconn.commit()
        self.bot.log.info("Updating server database complete")
