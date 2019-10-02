import socket


class Request:

    def __init__(self, log):
        self.PORT = 42069
        self.HOST = '127.0.0.1'
        self.log = log
        self.s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.s.connect((self.HOST, self.PORT))
        self.s.sendall(b'-auth interface')
        self.log.info("Connected to tcp server")

    def send(self, message):
        try:
            self.s.sendall(bytes(message, "utf-8"))
        except ConnectionResetError:       
            self.log.info("Lost tcp connection, reconnecting...")
            self.__init__(self.log)
            self.s.sendall(bytes(message, "utf-8"))

    def stop_bot(self):
        self.send('-relay stop bot')
   
    def start_bot(self):
        self.send('-relay start bot')
