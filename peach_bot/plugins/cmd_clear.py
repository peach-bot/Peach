async def run(message):
    await message.channel.purge(1)

def chatinvoke():
    return "clear"