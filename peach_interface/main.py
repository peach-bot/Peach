import logging
import random

import flask

from source import databasehandler

app = flask.Flask(__name__)
app.debug = False

@app.route("/")
def index():
    return flask.render_template("dashboard.html", botrunning=botrunning)

@app.route("/servers/")
def servers():
    return flask.render_template("servers.html")

@app.route("/stats/")
def stats():
    return flask.render_template("stats.html")

@app.route("/settings/")
def settings():
    return flask.render_template("settings.html")

@app.route("/logs/")
def logs():
    return flask.render_template("logs.html")

@app.route("/functions/dashboard/stop/")
def stop():
    global botrunning
    botrunning = False
    log.info("Attempting to stop bot")
    request.stop_bot()
    return flask.redirect(flask.url_for("index"), code=302)

@app.route("/functions/dashboard/start/")
def start():  
    global botrunning
    botrunning = True
    log.info("Attempting to start bot")
    request.start_bot()
    return flask.redirect(flask.url_for("index"), code=302)

if __name__ == "__main__":
    logging.basicConfig(format='%(name)s @ %(asctime)s - %(levelname)s: %(message)s', level=logging.INFO, datefmt='%Y-%m-%d %H:%M:%S')
    log = logging.getLogger('peach/interface')
    allowedloggers = ['peach/bot', 'peach/interface']
    for loggers in logging.Logger.manager.loggerDict:
        if loggers not in allowedloggers:
            logging.getLogger(loggers).disabled = True
        else:
            pass

    db = databasehandler.DatabaseHandler(log)
    
    botrunning = random.choice([True, False])
    log.info('Starting flask')
    app.run(host="0.0.0.0")
