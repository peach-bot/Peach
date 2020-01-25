import psycopg2
import os
import json
import time
import datetime

class DatabaseHandler:

    def __init__(self, log):
        self.log = log
        self.dbconn = psycopg2.connect("host={0} port={1} user={2} password={3} dbname={4} application_name='peach - interface'".format(os.environ['DBHOST'], os.environ['DBPORT'], os.environ['DBUSER'], os.environ['DBPASSWORD'], os.environ['DBNAME'],))

        self.dbcur = self.dbconn.cursor()

    def fetch_settings(self, serverid):
        """Returns an epic dictionary containing a server's settings."""
        rawsettings = self.plugin_serverconfig_get(serverid)
        settings = {
            "plugins": {}
        }
        for setting in rawsettings:
            if setting[0] not in settings["plugins"]:
                settings["plugins"][setting[0]] = {}
            settings["plugins"][setting[0]][setting[1]] = {
                "type": setting[4]["type"],
                "value": setting[4]["value"],
            }
        return settings

    def plugin_serverconfig_get(self, serverid):
        """Grabs a server's settings from the database."""
        self.dbcur.execute("SELECT defaults.pluginid, defaults.cfgkey, serverconfig.serverid, defaults.position, CASE WHEN serverconfig.cfgvalue IS NULL THEN defaults.cfgvalue ELSE serverconfig.cfgvalue END AS result FROM defaults LEFT JOIN (SELECT * FROM serverconfig WHERE serverid = {}) serverconfig ON defaults.pluginid = serverconfig.pluginid AND defaults.cfgkey = serverconfig.cfgkey ORDER BY pluginid, position".format(serverid))
        data = self.dbcur.fetchall()
        return data

    def updatesettings(self, form, serverid):
        for setting in form:
            if setting.type != "CSRFTokenField":
                if setting.name != "submit":
                    plugin = setting.name.split("_", 1)[0]
                    try:
                        dbsetting = setting.name.split("_", 1)[1].replace("_", " ", -1).title()
                    except IndexError:
                        dbsetting = setting.name.split("_", 1)[1]
                    if setting.type == "BooleanField":
                        settingtype = "bool"
                    self.dbcur.execute("INSERT INTO serverconfig VALUES ({0}, '{1}', '{2}', '{3}') ON CONFLICT (serverid, pluginid, cfgkey) DO UPDATE SET cfgvalue = '{3}' WHERE serverconfig.serverid = {0} AND serverconfig.pluginid = '{1}' AND serverconfig.cfgkey = '{2}'".format(serverid, plugin, dbsetting, json.dumps({"type": settingtype, "value": str(setting.data).lower()})))
                    self.dbconn.commit()

    def getactivitydatamonth(self, serverid):
        today = int(datetime.datetime.fromtimestamp(time.time()).replace(hour=0, minute=0, second=0, microsecond=0).timestamp())

        #get the timestamp for 30 days ago
        timestamp = today - 86400*30

        #create empty lists n stuff
        activitybuffer = []
        activitydata = []

        #fetch data from database
        self.dbcur.execute("SELECT unixtimestamp, SUM(messages) FROM channelstats WHERE serverid = {} GROUP BY unixtimestamp".format(serverid))
        dbdata = self.dbcur.fetchall()

        #calulate all the points for the buffer
        while True:
            for activity in dbdata:
                if timestamp == activity[0]:
                    activitybuffer.append({"x": timestamp, "y": activity[1]})
                    break
            else:
                activitybuffer.append({"x": timestamp, "y": 0})
            timestamp = timestamp+1800
            if timestamp > today:
                break

        #revert timestamp back to 30 days ago
        timestamp = today - 86400*30

        #merge timeframes together
        day = 0
        nextday = timestamp + 86400
        for activity in activitybuffer:
            day += activity["y"]
            timestamp = timestamp+1800
            if timestamp == nextday:
                activitydata.append({"x": nextday-86400, "y": day})
                nextday = timestamp + 86400
                day = 0

        return activitydata

    def getactivitydatayear(self, serverid):
        today = int(datetime.datetime.fromtimestamp(time.time()).replace(hour=0, minute=0, second=0, microsecond=0).timestamp())

        #get the timestamp for a year ago
        timestamp = today - 86400*365

        #create empty lists n stuff
        activitybuffer = []
        activitydata = []

        #fetch data from database
        self.dbcur.execute("SELECT unixtimestamp, SUM(messages) FROM channelstats WHERE serverid = {} GROUP BY unixtimestamp".format(serverid))
        dbdata = self.dbcur.fetchall()

        #calulate all the points for the buffer
        while True:
            for activity in dbdata:
                if timestamp == activity[0]:
                    activitybuffer.append({"x": timestamp, "y": activity[1]})
                    break
            else:
                activitybuffer.append({"x": timestamp, "y": 0})
            timestamp = timestamp+1800
            if timestamp == today:
                break

        #revert timestamp back to a year ago
        timestamp = today - 86400*365

        #merge timeframes together
        day = 0
        nextday = timestamp + 86400
        for activity in activitybuffer:
            day += activity["y"]
            timestamp = timestamp+1800
            if timestamp == nextday:
                activitydata.append({"x": nextday-86400, "y": day})
                nextday = timestamp + 86400
                day = 0

        #revert timestamp back to a year ago
        timestamp = today - 86400*365
        activitybuffer = activitydata
        activitydata = []
        month = 0
        first = activitybuffer[0]["x"]
        #merge days into months
        for activity in activitybuffer:
            if datetime.datetime.fromtimestamp(activity["x"]).day == 1:
                activitydata.append({"x": first, "y": month})
                month = 0
                first = activity["x"]
            month += activity["y"]
        activitydata.append({"x": first, "y": month})

        print(activitydata)
        return activitydata