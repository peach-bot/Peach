import psycopg2

class DatabaseHandler:

    async def __init__(self, bot):
        self.bot = bot