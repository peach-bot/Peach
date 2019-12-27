import logging
import random
import flask
import os

from source import databasehandler
from source.oauth import Oauth

app = flask.Flask(__name__)
app.debug = False

@app.route("/")
def index():
    try:
        flask.session["access_token"]
        return flask.redirect(flask.url_for("dashboard"), code=302)
    except KeyError:
        return flask.render_template("login.html")

@app.route("/discord_auth/", methods = ["get"])
def discord_auth():
    return flask.redirect(Oauth.discord_login_url)

@app.route("/login/", methods = ["get"])
def login():
    flask.session["access_token"] = Oauth.get_access_token(flask.request.args.get("code"))
    user_json = Oauth.get_user_json(flask.session["access_token"])
    flask.session["user_guilds"] = Oauth.get_user_servers(flask.session["access_token"])
    flask.session["selected_server"] = flask.session["user_guilds"][0]
    flask.session["avatar_url"] = "https://cdn.discordapp.com/avatars/{}/{}.png?size=128".format(user_json.get("id"), user_json.get("avatar"))
    flask.session["username"] = user_json.get("username")
    return flask.redirect(flask.url_for("dashboard"), code=302)

@app.route("/select_server/")
def select_server():
    try:
        flask.session["access_token"]
        if flask.request.args.get("id") == None:
            return flask.render_template("select_server.html")
        for server in flask.session["user_guilds"]:
            if server[2] == flask.request.args.get("id"):
                flask.session["selected_server"] = server
        if flask.request.referrer != None:
            print(db.fetch_settings(flask.session["selected_server"][2]))
            return flask.redirect(flask.request.referrer)
        else:
            return flask.redirect(flask.url_for("index"), code=302)
    except KeyError:
        return flask.redirect(flask.url_for("index"), code=302)

@app.route("/dashboard/")
def dashboard():
    try:
        return flask.render_template("dashboard.html", username=flask.session["username"], avatar_url=flask.session["avatar_url"], servers=flask.session["user_guilds"], current_server=flask.session["selected_server"])
    except KeyError:
        return flask.redirect(flask.url_for("index"), code=302)

@app.route("/servers/")
def servers():
    try:
        return flask.render_template("servers.html", username=flask.session["username"], avatar_url=flask.session["avatar_url"], servers=flask.session["user_guilds"], current_server=flask.session["selected_server"])
    except KeyError:
        return flask.redirect(flask.url_for("index"), code=302)

@app.route("/logout/")
def logout():
    flask.session.clear()
    return flask.redirect(flask.url_for("index"), code=302)

@app.route("/stats/")
def stats():
    try:
        return flask.render_template("stats.html", username=flask.session["username"], avatar_url=flask.session["avatar_url"], servers=flask.session["user_guilds"], current_server=flask.session["selected_server"])
    except KeyError:
        return flask.redirect(flask.url_for("index"), code=302)

@app.route("/settings/")
def settings():
    try:
        return flask.render_template("settings.html", username=flask.session["username"], avatar_url=flask.session["avatar_url"], servers=flask.session["user_guilds"], current_server=flask.session["selected_server"])
    except KeyError:
        return flask.redirect(flask.url_for("index"), code=302)

@app.route("/logs/")
def logs():
    try:
        return flask.render_template("logs.html", username=flask.session["username"], avatar_url=flask.session["avatar_url"], servers=flask.session["user_guilds"], current_server=flask.session["selected_server"])
    except KeyError:
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
    log.info('Starting flask')
    app.secret_key = os.urandom(24)
    app.templates_auto_reload = True
    app.debug = True
    app.run(host="0.0.0.0")
