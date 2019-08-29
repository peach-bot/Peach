import discord
import socket
import asyncio
from commands import commandhandler

class Peach(discord.Client):
    def bind(self, bot, log):
        self.bot = bot
        self.log = log

    async def tcploop(self):
        self.PORT = 42069
        self.HOST = '127.0.0.1'
        with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
            s.connect((self.HOST, self.PORT))
            s.sendall(b'Hello, world')
            data = s.recv(1024)

    async def on_ready(self):
        self.CommandSelector = commandhandler.commandSelector()
        self.log.info('{0.user} is logged in and online.'.format(self.bot))
        self.log.info("Creating tcp connection")
        self.connection_task = self.loop.create_task(self.tcploop())

    async def on_message(self, message):
        if message.author == self.bot.user:
            return

        if message.content.startswith('!'):
            command = message.content.split()[0][1:]
            await getattr(self.CommandSelector, command)(message)

    async def on_member_join(self, member):
        # Welcome message
        await member.guild.system_channel.send('{0.mention} felt cute.'.format(member))
        self.log.info('{0.mention} joined the server.'.format(member))