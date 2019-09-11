def chatinvoke():
    return "help"

async def run(message):
    await message.channel.send("Hello, I am stupid")
