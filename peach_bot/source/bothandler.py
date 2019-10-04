import asyncio
import json

import discord

import _thread as thread
from source import interfacehandler, pluginhandler, eventhandler


class Peach(discord.Client, eventhandler.EventHandler):
    """Bot class"""

    async def updatepresence(self, message: str):
        """Updates discord rich presence to display something super funny"""
        self.log.info("Updated Rich Presence")
        await self.change_presence(status=discord.Status.online, activity=discord.Game(name=message, details="all day long"))

    def bind(self, log):
        """Binds a logger to the bot class"""
        self.log = log

    async def on_ready(self):
        self.pluginhandler = pluginhandler.PluginHandler(self, self.log)
        self.interfacehandler = interfacehandler.InterfaceHandler(self.log, self, self.pluginhandler)
        self.log.info('{0.user} is logged in and online'.format(self))
        self.log.info("Creating tcp connection")
        thread.start_new_thread(self.interfacehandler.tcploop, ())
        await self.updatepresence("with eggplants")
        self.log.info('Startup complete!')
        await self.pluginhandler.on_ready()

    async def shutdown(self):
        self.logout()

        quit()