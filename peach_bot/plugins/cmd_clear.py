def define():
    return "clear"

async def run(message):
    await message.channel.purge(1)