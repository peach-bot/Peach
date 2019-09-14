import asyncio
import pluginhandler

import discord
import json

import _thread as thread
import interfacehandler

class Peach(discord.Client):
    """Main class"""
    async def loadconfig(self):
        with open("config.json") as f:
            config = json.load(f)

    async def updatepresence(self, presence):
        #await self.bot.change_presence(status=discord.Status.online, activity=discord.Streaming(name="Code", url="https://www.twitch.tv/jul_is_lazy", details="Coding"))
        await self.bot.change_presence(status=discord.Status.online, activity=discord.Game(name="with eggplants", details="all day long"))

    def bind(self, bot, log):
        self.bot = bot
        self.log = log

    async def on_ready(self):
        self.pluginhandler = pluginhandler.PluginHandler(self.bot, self.log)
        self.interfacehandler = interfacehandler.InterfaceHandler(self.log, self.bot, self.pluginhandler)
        self.log.info('{0.user} is logged in and online'.format(self.bot))
        self.log.info("Creating tcp connection")
        thread.start_new_thread(self.interfacehandler.tcploop, ())
        self.log.info("Loading config")
        await self.updatepresence("idle")
        self.log.info('Startup complete!')

    async def on_message(self, message):
        if message.author == self.bot.user:
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
        self.bot.logout()

        quit()