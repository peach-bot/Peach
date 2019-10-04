class EventHandler:
    """Another very boring and repetetive class that calls the event functions from the respectable plugins."""

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

    async def on_resumed(self):
        await self.pluginhandler.on_resumed()

    async def on_typing(self, channel, user, when):
        await self.pluginhandler.on_typing(channel, user, when)
    
    async def on_bulk_message_delete(self, messages):
        await self.pluginhandler.on_bulk_message_delete(messages)
    
    async def on_message_edit(self, before, after):
        await self.pluginhandler.on_message_edit(before, after)
    
    async def on_reaction_add(self, reaction, user):
        await self.pluginhandler.on_reaction_add(before, after)
    
    async def on_reaction_remove(self, reaction, user):
        await self.pluginhandler.on_reaction_remove(before, after)
    
    async def on_guild_channel_create(self, channel):
        await self.pluginhandler.on_guild_channel_create(channel)
    
    async def on_guild_channel_delete(self, channel):
        await self.pluginhandler.on_guild_channel_delete(channel)
    
    async def on_guild_channel_update(self, before, after):
        await self.pluginhandler.on_guild_channel_update(before, after)
    
    async def on_guild_channel_pins_update(self, channel, last_pin):
        await self.pluginhandler.on_guild_channel_pins_update(channel, last_pin)
    
    async def on_guild_integrations_update(self, guild):
        await self.pluginhandler.on_guild_integrations_update(guild)
    
    async def on_webhooks_update(self, channel):
        await self.pluginhandler.on_webhooks_update(channel)
    
    async def on_member_remove(self, member):
        await self.pluginhandler.on_member_remove(member)
    
    async def on_member_update(self, before, after):
        await self.pluginhandler.on_member_update(before, after)
    
    async def on_user_update(self, before, after):
        await self.pluginhandler.on_user_update(before, after)
    
    async def on_guild_join(self, guild):
        await self.pluginhandler.on_guild_join(guild)
    
    async def on_guild_remove(self, guild):
        await self.pluginhandler.on_guild_remove(guild)
    
    async def on_guild_update(self, before, after):
        await self.pluginhandler.on_guild_update(before, after)
    
    async def on_guild_role_create(self, role):
        await self.pluginhandler.on_guild_role_create(role)
    
    async def on_guild_role_delete(self, role):
        await self.pluginhandler.on_guild_role_delete(role)
    
    async def on_guild_role_update(self, before, after):
        await self.pluginhandler.on_guild_role_update(before, after)
    
    async def on_guild_emojis_update(self, guild, before, after):
        await self.pluginhandler.on_guild_emojis_update(guild, before, after)
    
    async def on_guild_available(self, guild):
        await self.pluginhandler.on_guild_available(guild)
    
    async def on_guild_unavailable(self, guild):
        await self.pluginhandler.on_guild_unavailable(guild)
    
    async def on_voice_state_update(self, member, before, after):
        await self.pluginhandler.on_voice_state_update(member, before, after)
    
    async def on_member_ban(self, guild, user):
        await self.pluginhandler.on_member_ban(guild, user)
    
    async def on_member_unban(self, guild, user):
        await self.pluginhandler.on_member_unban(guild, user)