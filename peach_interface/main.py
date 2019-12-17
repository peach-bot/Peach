import logging
import random
import flask

from source import databasehandler
from source.oauth import Oauth

app = flask.Flask(__name__)
app.debug = False

@app.route("/")
def index():
    return flask.render_template("login.html")

@app.route("/discord_auth/", methods = ["get"])
def discord_auth():
    return flask.redirect(Oauth.discord_login_url)

@app.route("/login/", methods = ["get"])
def login():
    global access_token
    code = flask.request.args.get("code")
    access_token = Oauth.get_access_token(code)
    return flask.redirect(flask.url_for("dashboard"), code=302)

@app.route("/dashboard/")
def dashboard():
    user_json = Oauth.get_user_json(access_token)
    print(user_json)
    avatar_url = "https://cdn.discordapp.com/avatars/{}/{}.png?size=32".format(user_json.get("id"), user_json.get("avatar"))
    return flask.render_template("dashboard.html", username=user_json.get("username"), avatar_url=avatar_url)

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
