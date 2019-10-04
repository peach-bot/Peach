class EventRelayer:
    """A very boring and repetetive class that calls the event functions from the respectable plugins."""    

    async def on_ready(self):
        try:
            for plugin in self.eventlink["on_ready"]:
                plugin.on_ready(self.bot)
        except KeyError:
            #no plugins for on_ready
            pass

    async def on_message(self, message):
        try:
            for plugin in self.eventlink["on_message"]:
                plugin.on_message(self.bot, message)
        except KeyError:
            #no plugins for on_message
            pass

    async def on_member_join(self, member):
        try:
            for plugin in self.eventlink["on_member_join"]:
                plugin.on_member_join(self.bot, member)
        except KeyError:
            #no plugins for on_member_join
            pass

    async def on_message_delete(self, message):
        try:
            for plugin in self.eventlink["on_message_delete"]:
                plugin.on_message_delete(self.bot, message)
        except KeyError:
            #no plugins for on_message_delete
            pass

    async def on_connect(self):
        try:
            for plugin in self.eventlink["on_connect"]:
                plugin.on_connect(self.bot)
        except KeyError:
            #no plugins for on_connect
            pass

    async def on_disconnect(self):
        try:
            for plugin in self.eventlink["on_disconnect"]:
                plugin.on_disconnect(self.bot)
        except KeyError:
            #no plugins for on_disconnect
            pass

    async def on_resumed(self):
        try:
            for plugin in self.eventlink["on_resumed"]:
                plugin.on_resumed(self.bot)
        except KeyError:
            #no plugins for on_resumed
            pass

    async def on_typing(self, channel, user, when):
        try:
            for plugin in self.eventlink["on_typing"]:
                plugin.on_typing(self.bot, channel, user, when)
        except KeyError:
            #no plugins for on_typing
            pass

    async def on_bulk_message_delete(self, messages):
        try:
            for plugin in self.eventlink["on_bulk_message_delete"]:
                plugin.on_bulk_message_delete(self.bot, messages)
        except KeyError:
            #no plugins for on_bulk_message_delete
            pass

    async def on_message_edit(self, before, after):
        try:
            for plugin in self.eventlink["on_message_edit"]:
                plugin.on_message_edit(self.bot, before, after)
        except KeyError:
            #no plugins for on_message_edit
            pass

    async def on_reaction_add(self, reaction, user):
        try:
            for plugin in self.eventlink["on_reaction_add"]:
                plugin.on_reaction_add(self.bot, reaction, user)
        except KeyError:
            #no plugins for on_reaction_add
            pass

    async def on_reaction_remove(self, reaction, user):
        try:
            for plugin in self.eventlink["on_reaction_remove"]:
                plugin.on_reaction_remove(self.bot, reaction, user)
        except KeyError:
            #no plugins for on_reaction_remove
            pass

    async def on_reaction_clear(self, message, reactions):
        try:
            for plugin in self.eventlink["on_reaction_clear"]:
                plugin.on_reaction_clear(self.bot, message, reactions)
        except KeyError:
            #no plugins for on_reaction_clear
            pass

    async def on_guild_channel_create(self, channel):
        try:
            for plugin in self.eventlink["on_guild_channel_create"]:
                plugin.on_guild_channel_create(self.bot, channel)
        except KeyError:
            #no plugins for on_guild_channel_create
            pass

    async def on_guild_channel_delete(self, channel):
        try:
            for plugin in self.eventlink["on_guild_channel_delete"]:
                plugin.on_guild_channel_delete(self.bot, channel)
        except KeyError:
            #no plugins for on_guild_channel_delete
            pass

    async def on_guild_channel_update(self, before, after):
        try:
            for plugin in self.eventlink["on_guild_channel_update"]:
                plugin.on_guild_channel_update(self.bot, before, after)
        except KeyError:
            #no plugins for on_guild_channel_update
            pass

    async def on_guild_channel_pins_update(self, channel, last_pin):
        try:
            for plugin in self.eventlink["on_guild_channel_pins_update"]:
                plugin.on_guild_channel_pins_update(self.bot, channel, last_pin)
        except KeyError:
            #no plugins for on_guild_channel_pins_update
            pass

    async def on_guild_integrations_update(self, guild):
        try:
            for plugin in self.eventlink["on_guild_integrations_update"]:
                plugin.on_guild_integrations_update(self.bot, guild)
        except KeyError:
            #no plugins for on_guild_integrations_update
            pass

    async def on_webhooks_update(self, channel):
        try:
            for plugin in self.eventlink["on_webhooks_update"]:
                plugin.on_webhooks_update(self.bot, channel)
        except KeyError:
            #no plugins for on_webhooks_update
            pass

    async def on_member_remove(self, member):
        try:
            for plugin in self.eventlink["on_member_remove"]:
                plugin.on_member_remove(self.bot, member)
        except KeyError:
            #no plugins for on_member_remove
            pass

    async def on_member_update(self, before, after):
        try:
            for plugin in self.eventlink["on_member_update"]:
                plugin.on_member_update(self.bot, before, after)
        except KeyError:
            #no plugins for on_member_update
            pass

    async def on_user_update(self, before, after):
        try:
            for plugin in self.eventlink["on_user_update"]:
                plugin.on_user_update(self.bot, before, after)
        except KeyError:
            #no plugins for on_user_update
            pass

    async def on_guild_join(self, guild):
        try:
            for plugin in self.eventlink["on_guild_join"]:
                plugin.on_guild_join(self.bot, guild)
        except KeyError:
            #no plugins for on_guild_join
            pass

    async def on_guild_remove(self, guild):
        try:
            for plugin in self.eventlink["on_guild_remove"]:
                plugin.on_guild_remove(self.bot, guild)
        except KeyError:
            #no plugins for on_guild_remove
            pass

    async def on_guild_update(self, before, after):
        try:
            for plugin in self.eventlink["on_guild_update"]:
                plugin.on_guild_update(self.bot, before, after)
        except KeyError:
            #no plugins for on_guild_update
            pass

    async def on_guild_role_create(self, role):
        try:
            for plugin in self.eventlink["on_guild_role_create"]:
                plugin.on_guild_role_create(self.bot, role)
        except KeyError:
            #no plugins for on_guild_role_create
            pass

    async def on_guild_role_delete(self, role):
        try:
            for plugin in self.eventlink["on_guild_role_delete"]:
                plugin.on_guild_role_delete(self.bot, role)
        except KeyError:
            #no plugins for on_guild_role_delete
            pass

    async def on_guild_role_update(self, before, after):
        try:
            for plugin in self.eventlink["on_guild_role_update"]:
                plugin.on_guild_role_update(self.bot, before, after)
        except KeyError:
            #no plugins for on_guild_role_update
            pass

    async def on_guild_emojis_update(self, guild, before, after):
        try:
            for plugin in self.eventlink["on_guild_emojis_update"]:
                plugin.on_guild_emojis_update(self.bot, guild, before, after)
        except KeyError:
            #no plugins for on_guild_emojis_update
            pass

    async def on_guild_available(self, guild):
        try:
            for plugin in self.eventlink["on_guild_available"]:
                plugin.on_guild_available(self.bot, guild)
        except KeyError:
            #no plugins for on_guild_available
            pass

    async def on_guild_unavailable(self, guild):
        try:
            for plugin in self.eventlink["on_guild_unavailable"]:
                plugin.on_guild_unavailable(self.bot, guild)
        except KeyError:
            #no plugins for on_guild_unavailable
            pass

    async def on_voice_state_update(self, member, before, after):
        try:
            for plugin in self.eventlink["on_voice_state_update"]:
                plugin.on_voice_state_update(self.bot, member, before, after)
        except KeyError:
            #no plugins for on_voice_state_update
            pass

    async def on_member_ban(self, guild, user):
        try:
            for plugin in self.eventlink["on_member_ban"]:
                plugin.on_member_ban(self.bot, guild, user)
        except KeyError:
            #no plugins for on_member_ban
            pass

    async def on_member_unban(self, guild, user):
        try:
            for plugin in self.eventlink["on_member_unban"]:
                plugin.on_member_unban(self.bot, guild, user)
        except KeyError:
            #no plugins for on_member_unban
            pass