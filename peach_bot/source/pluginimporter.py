import importlib
import os
import re
import pathlib
import sys


def load_plugins(logger):
    """loads all modules in plugins folder"""

    plugin_path = pathlib.Path(os.path.dirname(__file__), 'plugins')
    pluginfiles = [i for i in plugin_path.glob("**/*.py")]
    source_prnt = pathlib.Path(os.path.dirname(__file__), "plugins")
    plugins = [path.relative_to(source_prnt) for path in pluginfiles]
    plugins = [path for path in plugins if not path.name.startswith("_")]
    plugins = ["." + str(i.with_suffix("")).replace(os.path.sep, ".") for i in plugins] # get a neat list like [".txt_clear", ".sub_txt_github.subcommands"]

    importlib.import_module('source.plugins')
    modules = []
    for plugin in plugins:
        print("Considering", plugin)
        if not plugin.startswith('__'): # This is crappy
            modules.append(importlib.import_module(plugin, package="source.plugins"))

    logger.info("Found {0} plugins".format(len(modules)))

    return modules
