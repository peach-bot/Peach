import discord
import asyncio
from commands import commandhandler

class Cuddler(discord.Client):
    def bind(self, bot, log):
        self.bot = bot
        self.log = log

    async def on_ready(self):
        self.CommandSelector = commandhandler.commandSelector()
        self.log.info('{0.user} is logged in and online.'.format(self.bot))
        self.log.info("Creating rabbit task")
        #self.receiver_task = self.loop.create_task(self.rabbitreceiver())

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