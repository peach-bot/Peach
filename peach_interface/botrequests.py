import socket

class requests:
    def __init__(self):
        self.PORT = 42069
        self.HOST = '127.0.0.1'

    def stop_bot(self):
        with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
            s.connect((self.HOST, self.PORT))
            s.sendall(b'stop the bot')
            while True:
                data = s.recv(1024)
                print(data.decode("utf-8"))
