"""Assigns the command name to the right function call."""
from commands import cmd_clear, cmd_help

class commandSelector:
    def __init__(self):
        pass

    async def clear(self, message):
        """Clears a given amount of messages from the invoke channel."""
        await cmd_clear.run(message)

    async def help(self, message):
        """Sends help text in invoke channel."""
        await cmd_help.run(message)