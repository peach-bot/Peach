"""Uptime plugin"""

import datetime, discord

def manifest():
    """Receive the plugins manifest."""
    manifest = {
            "commands": {
                "uptime": {
                    "function": "uptime",
                    "permreqs": [],
                    "deleteinvoke": True,
                    "invoke": "uptime",
                    "aliases": ["up"]
                    }
                },
            "settings": {},
            "eventhooks": [],
            "chronos": {}
            }
    return manifest

async def uptime(message, bot):
    
    updatestring = "on {} at {}".format(bot.startuptime.replace(microsecond=0).strftime("%d.%m.%Y"), bot.startuptime.replace(microsecond=0).time())
    uptimestring = ""
    y = datetime.datetime.now()-bot.startuptime.replace(microsecond=0)
    hours = int(y.seconds/60/60)
    minutes = int(y.seconds/60-hours*60)
    if y.days > 0:
        uptimestring = uptimestring+"{} days".format(y.days)

        if hours > 0 and minutes > 0:
            uptimestring = uptimestring+", {} hours and {} minutes".format(hours, minutes)
        elif minutes > 0:
            uptimestring = uptimestring+"and {} minutes".format(minutes)
        elif hours > 0:
            uptimestring = uptimestring+"and {} hours".format(hours)

    else:
        if hours > 0 and minutes > 0:
            uptimestring = uptimestring+"{} hours and {} minutes".format(hours, minutes)
        elif minutes > 0:
            uptimestring = "{} minutes".format(minutes)
        elif hours > 0:
            uptimestring = "{} hours".format(hours)

    if y.days == 0 and hours == 0 and minutes == 0:
        uptimestring = "0 minutes"


    embed=discord.Embed(title="UPTIME", color=0xff886b)
    embed.add_field(name="The bot was started", value=str(updatestring), inline=False)
    embed.add_field(name="and has been running for", value=uptimestring, inline=False)
    embed.set_footer(text="powered by Peach")
    await message.channel.send(embed=embed)

async def uptime_man():
    embed = discord.Embed(title="Manual - !uptime", description="How to use this command:", color=0xff886b)
    embed.add_field(name="!uptime", value="Tells you how long the bot has been operating for since the last restart.", inline=False)
    embed.set_footer(text="powered by Peach")
    return embed