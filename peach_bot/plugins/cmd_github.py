import discord


def define():
    """This function defines the commands configuration"""
    plugindef = {
        "chatinvoke": "github",
        "permreq": [],
    }
    return plugindef

async def run(message, bot):
    subcommand = message.content.split()[1]
    

async def man():
    embed = discord.Embed(title="Manual - !clear", description="How to use this command:", color=0xff886b)
    embed.add_field(name="!clear [amount]", value="Clear a set amount of messages from the invoking channel.\nAmount can range from 1 to 50 messages.", inline=False)
    embed.set_footer(text="brought to you by Peach")
    return embed
