import time, datetime

class Analytics:

    def __init__(self, bot):
        self.bot = bot
        pass

    async def increment_messages(self, channelid: int, serverid: int):
        """Logs a message :3"""

        if datetime.datetime.fromtimestamp(time.time()).minute < 30:
            timestamp = int(datetime.datetime.fromtimestamp(time.time()).replace(minute=0, second=0, microsecond=0).timestamp())
        else:
            timestamp = int(datetime.datetime.fromtimestamp(time.time()).replace(minute=30, second=0, microsecond=0).timestamp())

        await self.bot.db.increment_messages(channelid, timestamp, serverid)