import asyncio
import json

import discord

import _thread as thread
import interfacehandler
import pluginhandler


class Peach(discord.Client):
    """Main class"""

    async def updatepresence(self):
        self.log.info("Updated Rich Presence")
        await self.change_presence(status=discord.Status.online, activity=discord.Game(name="with eggplants", details="all day long"))

    def bind(self, log):
        self.log = log

    async def on_ready(self):
        self.pluginhandler = pluginhandler.PluginHandler(self, self.log)
        self.interfacehandler = interfacehandler.InterfaceHandler(self.log, self, self.pluginhandler)
        self.log.info('{0.user} is logged in and online'.format(self))
        self.log.info("Creating tcp connection")
        thread.start_new_thread(self.interfacehandler.tcploop, ())
        await self.updatepresence()
        self.log.info('Startup complete!')

    async def on_message(self, message):
        if message.author == self.user:
            return

        if message.content.startswith('!man'):
            await self.pluginhandler.man(message)

        elif message.content.startswith('!'):
            await self.pluginhandler.runcommand(message)


    async def on_member_join(self, member):
        # Welcome message
        await member.guild.system_channel.send('{0.mention} felt cute.'.format(member))
        self.log.info('{0.mention} joined the server.'.format(member))

    async def shutdown(self):
        self.logout()

        quit()
