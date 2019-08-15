import json
import asyncio
import logging
import os
from commands import commandhandler

import discord
import pika
from aiohttp import client_exceptions

bot = discord.Client()

def callback(ch, method, properties, body):
    log.info(" [x] Received %r" % body)


def rabbitreceiver():
    channel.start_consuming()

async def main(auth):
    print("before")
    await asyncio.gather(
        bot.run(json.load(auth)['TOKEN']),
        asyncio.run(rabbitreceiver()),
    )
    print("after")

@bot.event
async def on_ready():
    global CommandSelector
    CommandSelector = commandhandler.commandSelector()
    log.info('{0.user} is logged in and online.'.format(bot))

@bot.event
async def on_message(message):
    if message.author == bot.user:
        return

    if message.content.startswith('!'):
        command = message.content.split()[0][1:]
        await getattr(CommandSelector, command)(message)

@bot.event
async def on_member_join(member):
    # Welcome message
    await member.guild.system_channel.send('{0.mention} felt cute.'.format(member))
    log.info('{0.mention} joined the server.'.format(member))

if __name__ == "__main__":

    logging.basicConfig(format='%(asctime)s - %(levelname)s: %(message)s', level=logging.INFO, datefmt='%Y-%m-%d %H:%M:%S')
    log = logging.getLogger('cuddler-logger')
    allowedloggers = ['cuddler-logger']
    for loggers in logging.Logger.manager.loggerDict:
        if loggers not in allowedloggers:
            logging.getLogger(loggers).disabled = True
        else:
            pass

    log.info('Starting bot')

    connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
    channel = connection.channel()
    channel.queue_declare(queue='interface')
    channel.basic_consume(queue='interface',
                      auto_ack=True,
                      on_message_callback=callback)

    with open("auth.json") as auth:
        try:
            asyncio.run(main(auth))
        except client_exceptions.ClientConnectorError:
            log.error("No connection to discordapp.com available.")
