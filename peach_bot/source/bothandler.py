import asyncio
import json

import discord

import _thread as thread
from source import interfacehandler, pluginhandler, databasehandler


class Peach(discord.Client):
    """Bot class"""

    async def updatepresence(message: str):
        """Updates discord rich presence to display something super funny"""
        self.log.info("Updated Rich Presence")
        await self.change_presence(status=discord.Status.online, activity=discord.Game(name=message, details="all day long"))

    def bind(self, log):
        """Binds a logger to the bot class"""
        self.log = log

    async def on_ready(self):
        self.log.info('{0.user} is logged in and online'.format(self))
        #load plugins
        self.pluginhandler = pluginhandler.PluginHandler(self, self.log)
        #establish connection to interface
        self.interfacehandler = interfacehandler.InterfaceHandler(self.log, self, self.pluginhandler)
        thread.start_new_thread(self.interfacehandler.tcploop, ())
        #load database connection
        self.db = databasehandler.DatabaseHandler(self)
        #update rich presence
        await self.updatepresence("with eggplants")
        self.log.info('Startup complete!')

    async def on_message(self, message):
        self.log.info("Received message: {0}#{1}@{2} -->".format(message.author.name, message.author.discriminator, message.guild.name, message.content))

        #ignore messages sent by the bot
        if message.author == self.user:
            return
        #filter for manual page invokes
        if message.content.startswith('!man'):
            await self.pluginhandler.man(message)

        #try to run a command in message starts with prefix
        elif message.content.startswith('!'):
            await self.pluginhandler.runcommand(message)

        #add to analytics


    async def on_member_join(self, member):
        # Welcome message
        await member.guild.system_channel.send('{0.mention} felt cute.'.format(member))
        self.log.info('{0.mention} joined {0.guild.name}'.format(member))

    async def shutdown(self):
        self.logout()

        quit()
