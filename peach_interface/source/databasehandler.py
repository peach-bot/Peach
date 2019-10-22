import psycopg2

class DatabaseHandler:

    def __init__(self, log):
        self.log = log
