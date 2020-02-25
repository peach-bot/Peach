import discord


def manifest():
    """Receive the plugins manifest."""
    manifest = {
            "commands": {
                "clear": {
                    "function": "clear",
                    "permreqs": ["manage_messages"],
                    "deleteinvoke": True,
                    "invoke": "clear",
                    "aliases": ["c"]
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

async def clear(message, bot):
    if await bot.db.plugin_serverconfig_get(message.guild.id, "clear", "Enabled"):
        try:
            amount = int(message.content.split()[1])
            if 0 < amount <= 50:
                cleared = await message.channel.purge(limit = amount)
                return "Cleared **{0}** messages for you. :slight_smile:".format(len(cleared))
                bot.log.info("Plugin txt_clear: cleared {0} messages".format(len(cleared)))
            else:
                return "I was unable to clear that amount of messages for you. A maximum of 50 messages is allowed."
        except IndexError:
            return "Please give an amount of messages to clear (min: 1, max: 50)"

async def clear_man():
    embed = discord.Embed(title="Manual - !clear", description="How to use this command:", color=0xff886b)
    embed.add_field(name="!clear [amount]", value="Clear a set amount of messages from the invoking channel.\nAmount can range from 1 to 50 messages.", inline=False)
    embed.set_footer(text="powered by Peach")
    return embed
