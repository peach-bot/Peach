def define():
    """This function defines the commands configuration"""
    plugindef = {
        "chatinvoke": "clear",
        "permreq": ["manage_messages"],
    }
    return plugindef

async def run(message, bot):
    try:
        amount = int(message.content.split()[1])
        if 0 < amount <= 50:
            cleared = await message.channel.purge(limit = amount)
            return "Cleared **{0}** messages for you. :slight_smile:".format(len(cleared))
        else:
            return "I was unable to clear that amount of messages for you. A maximum of 50 messages is allowed."
    except IndexError:
        return "Please give an amount of messages to clear (min: 1, max: 50)"