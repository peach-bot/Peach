import socket


class InterfaceHandler:
    """This class manages the responses and actions taken from commands sent by the interface."""

    tcpresponses = {
        
    }

    def __init__(self, log, bot, pluginhandler):
        self.log = log
        self.bot = bot
        self.pluginhandler = pluginhandler

    def tcploop(self):
            self.log.info("Creating tcp connection")
            self.PORT = 42069
            self.HOST = '127.0.0.1'
            with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
                s.connect((self.HOST, self.PORT))
                s.sendall(b'-auth bot')
                self.log.info("tcploop is listening")
                while True:
                    data = s.recv(1024)
                    self.log.info("Received from interface:", data.decode("utf-8"))
