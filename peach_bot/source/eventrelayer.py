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