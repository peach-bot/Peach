from nameko.rpc import rpc

class BotService():
    name = "bot_service"

    @rpc
    def getservers(self):
        return bot.guilds
