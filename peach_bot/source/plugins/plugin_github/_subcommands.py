import requests, json, discord

async def add(discordid, username, db):
    #add entry to database
    await db.plugin_userglobalconfig_update(discordid, "github", "username", {'data': username})
    return "Your GitHub account was successfully linked :white_check_mark:"

async def pull(userid, db, discordusername):
    #load github username from database
    dbdata = await db.plugin_userglobalconfig_get(userid, "github", "username")
    #pull information from github

    #create embed
    githubdata = json.loads(requests.get('https://api.github.com/users/'+dbdata["data"], headers={'Accept': 'application/vnd.github.v3+json'}).text)
    starred = len(json.loads(requests.get("https://api.github.com/users/"+dbdata["data"]+"/starred", headers={'Accept': 'application/vnd.github.v3+json'}).text))
    response=discord.Embed(title="GitHub - "+dbdata["data"], url="https://www.github.com/"+dbdata["data"], description=discordusername+"'s GitHub account:", color=0xff886b)
    response.set_thumbnail(url=githubdata["avatar_url"])
    response.add_field(name="Repositories:", value=githubdata["public_repos"], inline=True)
    response.add_field(name="Followers:", value=githubdata["followers"], inline=True)
    response.add_field(name="Following:", value=githubdata["following"], inline=True)
    response.add_field(name="Starred Repositories:", value=starred, inline=True)
    #send embed
    return response
