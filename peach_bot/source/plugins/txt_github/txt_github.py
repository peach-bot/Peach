import discord
import source.plugins.txt_github._subcommands as subcommands


def define():
    """This function defines the commands configuration"""
    plugindef = {
        "type": "txt",
        "chatinvoke": "github",
        "aliases": "gh",
        "eventhooks": [],
        "deleteinvoke": False,
        "permreq": [],
        "interval": 60,
        }
    return plugindef

async def run(message, bot):
    #if !github add then add entry
    if message.content.split(" ")[1] == "add":
        response = await subcommands.add(message.author.id, message.content.split()[2], bot.db)

    else:
        messagelen = len(message.content.split(" "))    
        #if no other user is mentioned
        if messagelen == 2:
            #userid = author's id
            userid = message.author.id
            response = subcommands.pull(message.author, userid)
        else:
            #get mentioned user's id
            userid = message.content.split(" ")[1]
            response = subcommands.pull(message.author, userid)

    return response
    

async def man():
    embed = discord.Embed(title="Manual - !github", description="How to use this command:", color=0xff886b)
    embed.add_field(name="!github add [github username]", value="Links given GitHub account to your Discord ID", inline=False)
    embed.add_field(name="!github @user", value="Displays user's github profile", inline=False)
    embed.add_field(name="!github ", value="Displays your github profile", inline=False)
    embed.set_footer(text="brought to you by Peach")
    return embed