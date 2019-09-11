import asyncio
import socket
import pluginhandler

import discord

import _thread as thread
from interface import tcpresponse


class Peach(discord.Client):
    """Main class"""
    def bind(self, bot, log):
        self.bot = bot
        self.log = log

    def tcploop(self):
        self.PORT = 42069
        self.HOST = '127.0.0.1'
        self.responder = tcpresponse(self.log)
        with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
            s.connect((self.HOST, self.PORT))
            s.sendall(b'-auth bot')
            self.log.info("tcploop is listening")
            while True:
                data = s.recv(1024)
                self.log.info(data.decode("utf-8"))

    async def on_ready(self):
        self.pluginhandler = pluginhandler.PluginHandler()
        self.log.info('{0.user} is logged in and online'.format(self.bot))
        self.log.info("Creating tcp connection")
        thread.start_new_thread(self.tcploop, ())
        self.log.info('Done')

    async def on_message(self, message):
        if message.author == self.bot.user:
            return

        if message.content.startswith('!'):
            self.pluginhandler.runcommand(message)

    async def on_member_join(self, member):
        # Welcome message
        await member.guild.system_channel.send('{0.mention} felt cute.'.format(member))
        self.log.info('{0.mention} joined the server.'.format(member))
