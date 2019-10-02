import time


class InterfaceHandler:
    """Handles the connection to the interface and deals with its requests."""

    def __init__(self, log):
        self.log = log
        self.botconn = None
        self.botaddr = None

    def loop(self, conn, s, botconn, botaddr):
        """Listener loop"""
        self.botconn = botconn
        self.botaddr = botaddr
        self.conn = conn
        s.listen()
        while True:    
            self.receive()

    def receive(self):
        """blocking listening function n stuff"""
        data = self.conn.recv(4096).decode("utf-8")
        if not data: pass
        self.log.info("Received from interface: {}".format(data))

        if data != "":
            if data.split(" ")[0] == "-relay":
                self.log.info("Relaying")
                if self.botconn != None:
                    self.botconn.sendto(bytes(" ".join(data.split(" ")[1:]), "utf-8"), self.botaddr)
                else:
                    self.log.info("No bot connection")
