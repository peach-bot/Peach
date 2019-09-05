import socket

class requests:
    def __init__(self):
        self.PORT = 42069
        self.HOST = '127.0.0.1'
        self.s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.s.connect((self.HOST, self.PORT))
        self.s.sendall(b'-auth interface')

    def stop_bot(self):
        self.s.sendall(b'-relay stop bot')
   
    def start_bot(self):
        self.s.sendall(b'-relay start bot')
