import discord
import asyncio
import pika
from commands import commandhandler

class Cuddler(discord.Client):
    def bind(self, bot, log):
        self.bot = bot
        self.log = log

    def callback(self, ch, method, properties, body):
        print("less")
        self.log.info(" [x] Received %r" % body)

    async def on_ready(self):
        self.CommandSelector = commandhandler.commandSelector()
        self.log.info('{0.user} is logged in and online.'.format(self.bot))
        self.log.info("Creating rabbit task")
        self.receiver_task = self.loop.create_task(self.rabbitreceiver())

    async def on_message(self, message):
        if message.author == self.bot.user:
            return

        if message.content.startswith('!'):
            command = message.content.split()[0][1:]
            await getattr(self.CommandSelector, command)(message)

    async def rabbitreceiver(self):
        connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
        channel = connection.channel()
        await self.wait_until_ready()
        self.log.info("rabbit: starting consumption")
        while True:
            queue_state = channel.queue_declare(queue='interface', durable=True, passive=True)
            queue_empty = queue_state.method.message_count == 0
            if not queue_empty:
                print("yeah")
                method, properties, body = channel.basic_get(queue='interface', callback=callback, auto_ack=False)
                self.callback(channel, method, properties, body)
            await asyncio.sleep(2)
            print("loop")

    async def on_member_join(self, member):
        # Welcome message
        await member.guild.system_channel.send('{0.mention} felt cute.'.format(member))
        self.log.info('{0.mention} joined the server.'.format(member))