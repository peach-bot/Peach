"""Assigns the command name to the right function call."""
import pluginimporter

class CommandHandler:
    def __init__(self):
        # load all modules in plugins folder
        plugins = pluginimporter.load_plugins()

        # create command dict
        commandlink = {}
        for plugin in plugins:
            pluginname = plugin.__name__[8:]
            if pluginname.startswith("cmd_"):
                commandlink[getattr(plugin, "chatinvoke")()] = plugin

if __name__ == "__main__":
    # debug stuff
    commandhandler = CommandHandler()