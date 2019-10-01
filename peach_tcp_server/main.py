import socket
import logging
import interfacehandler
import _thread as thread

HOST = "127.0.0.1"
PORT = 42069


if __name__ == "__main__":
    logging.basicConfig(format='%(name)s @ %(asctime)s - %(levelname)s: %(message)s', level=logging.INFO, datefmt='%Y-%m-%d %H:%M:%S')
    log = logging.getLogger('peach/tcpserver')
    allowedloggers = ['peach/bot', 'peach/tcpserver', 'peach/interface']
    for loggers in logging.Logger.manager.loggerDict:
        if loggers not in allowedloggers:
            logging.getLogger(loggers).disabled = True
        else:
            pass
        
    ifhandler = interfacehandler.InterfaceHandler(log)

    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        s.bind((HOST, PORT))
        botconn = None
        botaddr = None
        while True:
            log.info("Listening")
            s.listen()
            conn, addr = s.accept()
            data = conn.recv(4096).decode("utf-8")
            if data.startswith("-auth"):
                splitdatastr = data.split(" ")
                if splitdatastr[1] == "interface":
                    log.info("New interface connection")
                    interfaceconn = conn
                    interfaceaddr = addr
                    thread.start_new_thread(ifhandler.loop, (conn, s, botconn, botaddr))

                if splitdatastr[1] == "bot":
                    log.info("New bot connection")
                    botconn = conn
                    ifhandler.botconn = botconn
                    botaddr = addr
                    ifhandler.botaddr = botaddr

