"""This file serves as a template for building chat command plugins."""
import discord

def manifest():
    """Receive the plugins manifest."""
    manifest = {
            "commands": {
                "example": {
                    "function": "example",
                    "permreqs": [],
                    "deleteinvoke": True,
                    "invoke": "example",
                    "aliases": ["xmpl"]
                    }
                },
            "settings": {
                "Enabled": "bool true",
            },
            "eventhooks": [],
            "chronos": {
                }
            }
    return manifest

async def example(message, bot):
    """The actual command that runs upon invoke"""
    if await bot.db.plugin_serverconfig_get(message.guild.id, "example", "Enabled"):
        return "This is an example for a command plugin. It doesn't really do anything (duh)." #Returns a response message that displays for 5 seconds or None

async def on_message(message, bot):
    """Example of the on_message event hook"""
    bot.log("New message: {0.author}: {0.content}".format(message))
    return None

async def example_man():
    """This defines the commands manual page."""
    embed = discord.Embed(title="Example manual - !example", description="How to use this command:", color=0xff886b)
    embed.add_field(name="!example", value="Explain what the command does here.", inline=False)
    embed.add_field(name="!example [attribute_1] [attribute_2]", value="List commands that use attributes like this.", inline=False)
    embed.set_footer(text="brought to you by Peach") #leave the footer like this
    return embed
