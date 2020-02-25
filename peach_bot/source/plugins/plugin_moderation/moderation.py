"""Basic moderation functionality."""

def manifest():
    """Receive the plugins manifest."""
    manifest = {
            "commands": {
                "ban": {
                    "function": "ban",
                    "permreqs": ["ban_members"],
                    "deleteinvoke": False,
                    "invoke": "ban",
                    "aliases": []
                    },
                "kick": {
                    "function": "kick",
                    "permreqs": ["kick_members"],
                    "deleteinvoke": False,
                    "invoke": "kick",
                    "aliases": []
                    }
                },
            "settings": {
                "Enabled": "bool true",
                "Kick": "bool true",
                "Ban": "bool true"
            },
            "eventhooks": [],
            "chronos": {
                }
            }
    return manifest

async def ban(message, bot):
    """Bans a member from the guild."""
    if await bot.db.plugin_serverconfig_get(message.guild.id, "moderation", "Ban") and await bot.db.plugin_serverconfig_get(message.guild.id, "moderation", "Enabled"):
        return "Banned"

async def kick(message, bot):
    """Kicks a member from the guild."""
    if await bot.db.plugin_serverconfig_get(message.guild.id, "moderation", "Kick") and await bot.db.plugin_serverconfig_get(message.guild.id, "moderation", "Enabled"):
        return "Kicked"