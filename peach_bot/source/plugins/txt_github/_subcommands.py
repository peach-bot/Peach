async def add(discordid, username, db):
    #add entry to database
    data = {
        'username': username
    } 
    await db.plugin_updateuser(discordid, "github", data)
    return "Your GitHub account was successfully linked :white_check_mark:"

async def pull(author):
    #load github username from database

    #pull information from github

    #create embed

    #send embed
    return None
