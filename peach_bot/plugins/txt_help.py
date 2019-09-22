def define():
    """This function defines the commands configuration"""
    plugindef = {
        "chatinvoke": "help",
        "permreq": [],
    }
    return plugindef

async def run(message, bot):
    await message.channel.send("Hello, I am stupid")
