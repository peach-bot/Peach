import asyncio
import concurrent.futures
import json
import logging
import os
import time

import discord
from aiohttp import client_exceptions
from bot import Peach

if __name__ == "__main__":

    bot = Peach()

    logging.basicConfig(format='%(asctime)s - %(levelname)s: %(message)s', level=logging.INFO, datefmt='%Y-%m-%d %H:%M:%S')
    log = logging.getLogger('peach-logger')
    allowedloggers = ['peach-logger']
    for loggers in logging.Logger.manager.loggerDict:
        if loggers not in allowedloggers:
            logging.getLogger(loggers).disabled = True
        else:
            pass
    
    with open("config.json") as config:
        config = json.load(config)

    bot.log = log
    log.info('Starting bot')

    
    try:
        bot.run(config['TOKEN'])
    except client_exceptions.ClientConnectorError:
        log.error("No connection to discordapp.com available.")
