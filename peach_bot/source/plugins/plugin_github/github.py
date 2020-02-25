import discord
import source.plugins.plugin_github._subcommands as subcommands


def manifest():
    """Receive the plugins manifest."""
    manifest = {
            "commands": {
                "github": {
                    "function": "github",
                    "permreqs": [],
                    "deleteinvoke": True,
                    "invoke": "github",
                    "aliases": ["gh"]
                    }
                },
            "settings": {
                "Enabled": "bool true",
                "Account Adding": "bool true",
                "Account Lookup": "bool true",
            },
            "eventhooks": [],
            "chronos": {
                }
            }
    return manifest

async def github(message, bot):
    if await bot.db.plugin_serverconfig_get(message.guild.id, "github", "Enabled"):
        #if !github add then add entry
        messagelen = len(message.content.split(" "))
        if messagelen > 1:
            if await bot.db.plugin_serverconfig_get(message.guild.id, "github", "Account Adding"):
                if message.content.split(" ")[1] == "add":
                    await message.channel.send(await subcommands.add(message.author.id, message.content.split()[2], bot.db))
            if await bot.db.plugin_serverconfig_get(message.guild.id, "github", "Account Lookup"):
                if message.content.split(" ")[1].startswith("<@!"):
                    userid = message.content.split(" ")[1][3:21]
                    await message.channel.send(embed=await subcommands.pull(userid, bot.db, message.author.name+"#"+message.author.discriminator))
                elif message.content.split(" ")[1].startswith("<@"):
                    userid = message.content.split(" ")[1][2:20]
                    await message.channel.send(embed=await subcommands.pull(userid, bot.db, message.author.name+"#"+message.author.discriminator))

        else:
            if await bot.db.plugin_serverconfig_get(message.guild.id, "github", "Account Lookup"):
                userid = message.author.id
                await message.channel.send(embed=await subcommands.pull(userid, bot.db, message.author.name+"#"+message.author.discriminator))

        return None

async def github_man():
    embed = discord.Embed(title="Manual - !github", description="How to use this command:", color=0xff886b)
    embed.add_field(name="!github add [github username]", value="Links given GitHub account to your Discord ID", inline=False)
    embed.add_field(name="!github @user", value="Displays user's github profile", inline=False)
    embed.add_field(name="!github ", value="Displays your github profile", inline=False)
    embed.set_footer(text="brought to you by Peach")
    return embed