import pyscopg2

class DatabaseHandler:

    async def __init__(bot):
        self.bot = bot