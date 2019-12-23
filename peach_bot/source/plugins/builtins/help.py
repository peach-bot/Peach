def manifest():
    """Receive the plugins manifest."""
    manifest = {
            "commands": {
                "help": {
                    "function": "help",
                    "permreqs": [],
                    "deleteinvoke": True,
                    "invoke": "help",
                    "aliases": ["h"]
                    }
                },
            "settings": {
                "Enabled": "bool true",
            },
            "eventhooks": [],
            "chronos": {
                }
            }
    return manifest

async def help(message, bot):
    await message.channel.send("Hello, I am stupid")

async def help_man():
    embed = discord.Embed(title="Manual - !help", description="How to use this command:", color=0xff886b)
    embed.add_field(name="!help", value="Sends a message explaining how to use the bot.", inline=False)
    embed.set_footer(text="powered by Peach")
    return embed