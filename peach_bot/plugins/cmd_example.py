"""This file serves as a template for building chat command plugins."""

def define():
    """This function defines the commands configuration"""
    plugindef = {
        "chatinvoke": "example", #string that invokes the command (prefix is defined in bot.py, leave out)
        "permreq": [], #list of required permissions
        "rolereq": 2, #int of required role level (0: Owner, 1: Admin, 2: Mod, 3: Trusted, 4: Verified, 5: Member, 6: Jail)
    }
    return plugindef

async def run():
    """The actual command that runs upon invoke"""
    pass