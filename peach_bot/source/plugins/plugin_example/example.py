"""This file serves as a template for building chat command plugins."""
import discord

def define():
    """This function defines the commands configuration"""
    plugindef = {
        "type": "mlt", #mark as a plugin invoked by multiple triggers
        "chatinvoke": "example", #string that invokes the command (prefix is defined in bot.py, leave out)
        "aliases": ["xmpl"], #other strings that invoke the command (for further explenation look in the docs under invoke hierarchy)
        "eventhooks": [], #list of events that invoke the plugin (e.g.: ["on_message", "on_member_join"])
        "deleteinvoke": True, #delete the invoking message
        "permreq": [], #list of required permissions
        "interval": 60, #chrono interval in minutes
        }
    return plugindef

async def run_txt(message, bot):
    """The actual command that runs upon invoke"""
    return "This is an example for a command plugin. It doesn't really do anything (duh)." #Returns a response message that displays for 5 seconds or None

async def on_message(message, bot):
    """Example of the on_message event hook"""
    bot.log("New message: {0.author}: {0.content}".format(message))

async def man():
    """This defines the commands manual page."""
    embed = discord.Embed(title="Example manual - !example", description="How to use this command:", color=0xff886b)
    embed.add_field(name="!example", value="Explain what the command does here.", inline=False)
    embed.add_field(name="!example [attribute_1] [attribute_2]", value="List commands that use attributes like this.", inline=False)
    embed.set_footer(text="brought to you by Peach") #leave the footer like this
    return embed
