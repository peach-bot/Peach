import importlib
import os
import re
import sys


def load_plugins():
    """loads all modules in plugins folder"""
    # casually taken from https://copyninja.info/blog/dynamic-module-loading.html
    pysearchre = re.compile('.py$', re.IGNORECASE)
    pluginfiles = filter(pysearchre.search, os.listdir(os.path.join(os.path.dirname(__file__), 'plugins')))
    form_module = lambda fp: '.' + os.path.splitext(fp)[0]
    plugins = map(form_module, pluginfiles)
    # import parent module / namespace
    importlib.import_module('plugins')
    modules = []
    for plugin in plugins:
             if not plugin.startswith('__'):
                 modules.append(importlib.import_module(plugin, package="plugins"))

    return modules
