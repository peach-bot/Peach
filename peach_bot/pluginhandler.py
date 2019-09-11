"""Assigns the command name to the right function call."""
import pluginimporter

class PluginHandler:
    """"""
    def __init__(self):
        # load all modules in plugins folder
        plugins = pluginimporter.load_plugins()

        # create command dict
        self.commandlink = {}
        for plugin in plugins:
            pluginname = plugin.__name__[8:]
            if pluginname.startswith("cmd_"):
                plugindef = getattr(plugin, "define")()
                self.commandlink[plugindef["chatinvoke"]] = (plugin, plugindef)

    async def runcommand(self, message):
        command = message.content.split()[0][1:]
        await self.commandlink[command][0](message)

if __name__ == "__main__":
    # debug stuff
    plguinhandler = PluginHandler()