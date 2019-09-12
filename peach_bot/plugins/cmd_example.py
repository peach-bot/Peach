"""This file serves as a template for building chat command plugins."""

def define():
    """This function defines the commands configuration"""
    plugindef = {
        "chatinvoke": "example", #string that invokes the command (prefix is defined in bot.py, leave out)
        "permreq": [], #list of required permissions
        }
    return plugindef

async def run(message, bot):
    """The actual command that runs upon invoke"""
    return "This is an example for a command plugin. It doesn't really do anything (duh)."

async def man():
    """This defines the commands manual page."""
    return "WIP"