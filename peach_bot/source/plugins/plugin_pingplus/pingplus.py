def manifest():
    """Receive the plugins manifest."""
    manifest = {
            "commands": {
                "pingplus": {
                    "function": "ping",
                    "permreqs": [],
                    "deleteinvoke": True,
                    "invoke": "ping",
                    "aliases": []
                    }
                },
            "settings": {
                "Enabled": "bool true",
            },
            "eventhooks": [],
            "chronos": {
                }
            }
    return manifest

async def ping(message, bot):
    if await bot.db.plugin_serverconfig_get(message.guild.id, "pingplus", "Enable"):
        return "Ping lol"