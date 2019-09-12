"""This file serves as a template for building chat command plugins."""
import discord

def define():
    """This function defines the commands configuration"""
    plugindef = {
        "chatinvoke": "example", #string that invokes the command (prefix is defined in bot.py, leave out)
        "permreq": [], #list of required permissions
        }
    return plugindef

async def run(message, bot):
    """The actual command that runs upon invoke"""
    return "This is an example for a command plugin. It doesn't really do anything (duh)." #Returns a response message that displays for 5 seconds or None

async def man():
    """This defines the commands manual page."""
    embed = discord.Embed(title="Example manual - !example", description="How to use this command:", color=0x9eff82)
    embed.add_field(name="!example", value="Explain what the command does here.", inline=False)
    embed.add_field(name="!example [attribute_1] [attribute_2]", value="List commands that use attributes like this.", inline=False)
    embed.set_footer(text="brought to you by Peach") #leave the footer like this
    return embed