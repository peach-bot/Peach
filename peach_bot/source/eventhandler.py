class EventHandler:
    
    async def on_message(self, message):
        self.log.info("Received message: {0}#{1}@{2} --> {3}".format(message.author.name, message.author.discriminator, message.guild.name, message.content))

        #ignore messages sent by the bot
        if message.author == self.user:
            return
        #filter for manual page invokes
        if message.content.startswith('!man'):
            await self.pluginhandler.man(message)

        #try to run a command in message starts with prefix
        elif message.content.startswith('!'):
            await self.pluginhandler.runcommand(message)

        await self.pluginhandler.on_message(message)


    async def on_member_join(self, member):
        # Welcome message
        await member.guild.system_channel.send('{0.mention} felt cute.'.format(member))
        self.log.info('{0.mention} joined {0.guild.name}'.format(member))

        await self.pluginhandler.on_member_join(member)

    async def on_message_delete(self, message):
        await self.pluginhandler.on_message_delete(message)

    async def on_connect(self):
        await self.pluginhandler.on_connect()
    
    async def on_disconnect(self):
        await self.pluginhandler.on_disconnect()

    async def on_message_delete(self, message):
        await self.pluginhandler.on_message_delete(message)
    

    async def on_message_delete(self, message):
        await self.pluginhandler.on_message_delete(message)
    

    async def on_message_delete(self, message):
        await self.pluginhandler.on_message_delete(message)
    