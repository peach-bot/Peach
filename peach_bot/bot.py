import asyncio
import pluginhandler

import discord

import _thread as thread
import interfacehandler

class Peach(discord.Client):
    """Main class"""
    def bind(self, bot, log):
        self.bot = bot
        self.log = log

    async def on_ready(self):
        self.pluginhandler = pluginhandler.PluginHandler(self.bot, self.log)
        self.interfacehandler = interfacehandler.InterfaceHandler(self.log, self.bot, self.pluginhandler)
        self.log.info('{0.user} is logged in and online'.format(self.bot))
        self.log.info("Creating tcp connection")
        thread.start_new_thread(self.interfacehandler.tcploop, ())
        self.log.info('Done')

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