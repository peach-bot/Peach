"""Assigns the command name to the right function call."""
import asyncio

import pluginimporter


class PluginHandler:
    """"""
    def __init__(self, bot, log):
        # bind bot and logger
        self.bot = bot
        self.log = log

        # load all modules in plugins folder
        self.log.info("Loading plugins...")
        plugins = pluginimporter.load_plugins()
        self.log.info("Loading plugins complete")

        self.log.info("Linking plugins...")
        # create command dict
        self.commandlink = {}
        for plugin in plugins:
            pluginname = plugin.__name__[8:]
            plugindef = getattr(plugin, "define")()
            self.commandlink[plugindef["chatinvoke"]] = (plugin, plugindef)
        self.log.info("Linking plugins complete")

    async def runcommand(self, message):
        #delete invoke from channel
        await message.delete()
        #filter the command they invoked
        command = message.content.split()[0][1:]

        #look if command exists
        if  command in self.commandlink:
            #get authors permissions
            permissions = message.author.permissions_in(message.channel)
            #grab module from linkdict
            plugin, plugindef = self.commandlink[command]
            #check if author's permissions are sufficient
            has_perms = True
            for required in plugindef["permreq"]:
                if not getattr(permissions, required):
                    has_perms = False
            
            #run command
            if has_perms:
                response = await plugin.run(message, self.bot)
                if response != None:
                    responsemessage = await message.channel.send(response)
                    await asyncio.sleep(5)
                    await responsemessage.delete()
            
            # if user doesn't have sufficient permissions
            else:
                refusalmessage = await message.channel.send("I'm sorry. You do not have sufficient permissions to use this command. :pensive:")
                await asyncio.sleep(5)
                await refusalmessage.delete()

    async def man(self, message):
        #delete invoke from channel
        await message.delete()
        #filter the command they invoked
        command = message.content.split()[1]
        await message.channel.send(embed = await self.commandlink[command][0].man())

    async def updatestate(self, message):

        #When invoked from chat
        if message != None:
            pass

        #invoked by interface
        else:
            pass
