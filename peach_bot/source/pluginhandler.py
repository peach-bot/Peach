"""Assigns the command name to the right function call."""
import asyncio

from source import pluginimporter


class PluginHandler:
    """Handles plugin importing and execution."""
    def __init__(self, bot, log):
        # bind bot and logger
        self.bot = bot
        self.log = log

    async def mapplugins(self):
        # load all modules in plugins folder
        self.log.info("Loading plugins...")
        plugins = pluginimporter.load_plugins(self.log)
        self.log.info("Loading plugins complete")

        #load plugins from Database
        data = await self.bot.db.query_return("SELECT * FROM plugins")
        dbids = []
        for dib in data:
            dbids.append(dib[0])

        #map plugin info
        self.log.info("Mapping plugins...")
        self.commandmap = {}
        self.eventmap = {}
        for plugin in plugins:
            if plugin.__name__.split(".")[3] not in dbids:
                await self.bot.db.query("INSERT INTO plugins VALUES ('{0}')".format(plugin.__name__.split(".")[3]))
            pluginmanifest = getattr(plugin, "manifest")()
            
            #update plugin settings
            position = 0
            for setting, setting_type in pluginmanifest["settings"].items():
                await self.bot.db.plugin_defaults_update("'{}'".format(plugin.__name__.split(".")[3]), "'{}'".format(setting), {"type": setting_type.split()[0], "value": setting_type.split()[1]}, position)
                position += 1

            #map plugin event hooks
            for event in pluginmanifest["eventhooks"]:
                try:
                    #add plugin to event's list
                    self.eventmap[event].append(plugin)
                except KeyError:
                    #if list doesn't exist add to new list
                    self.eventmap[event] = [plugin]
            
            #map commands
            for commandname, command in pluginmanifest["commands"].items():

                #map main invoke
                if command["invoke"] in self.commandmap: #check if invoke already in map
                    #invoke shouldn't be in map, bad moderation smh
                    self.log.error("Commands invoke is already in commandmap")
                else:
                    #if everything's fine add map
                    self.commandmap[command["invoke"]] = [0, command, plugin]
                
                #map aliases
                for index, alias in enumerate(command["aliases"]):
                    #index 0 is reserved for main invoke
                    index += 1
                    if alias in self.commandmap:
                        if index < self.commandmap[alias][0]:
                            #only overwrite alias if index is lower
                            self.commandmap[alias] = [index, command, plugin]
                    else:
                        #if alias doesn't exist add it
                        self.commandmap[alias] = [index, command, plugin]
        self.log.info("Mapping plugins complete")

    async def runevent(self, event: str, *args):
        """Run an event trigger in all plugins.
        
        event -- event name

        args  -- event related discord.py objects"""
        try:
            for plugin in self.eventmap[event]:
                getattr(plugin, event)(self.bot, *args)
        except KeyError:
            #no plugins for on_ready
            pass

    async def runcommand(self, message):
        """Run a plugin command. Takes in raw discord message and processes it."""
        #filter the command they invoked
        command = message.content.split()[0][1:]

        #look if command exists
        if command in self.commandmap:
            #get authors permissions
            permissions = message.author.permissions_in(message.channel)
            #grab module from mapdict
            prioritylevel, command, plugin = self.commandmap[command]
            #check if author's permissions are sufficient
            has_perms = True
            for required in command["permreqs"]:
                if not getattr(permissions, required):
                    has_perms = False

            #run command
            if has_perms:
                #delete invoke from channel
                if command["deleteinvoke"]:
                    #try to delete invoke message, if not it isn't important
                    try:
                        await message.delete()
                    except Exception:
                        pass
                response = await getattr(plugin, command["function"])(message, self.bot)
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
        await message.channel.send(embed = await getattr(self.commandmap[command][2], command+"_man")())
