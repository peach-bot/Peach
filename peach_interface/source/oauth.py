import os, requests

class Oauth(object):
    client_id = os.environ["CLIENT_ID"]
    client_secret = os.environ["CLIENT_SECRET"]
    scope = "identify%20guilds"
    redirect_uri = "http://127.0.0.1:5000/login"
    discord_login_url = "https://discordapp.com/api/oauth2/authorize?client_id={}&redirect_uri={}&response_type=code&scope={}".format(client_id, redirect_uri, scope)
    discord_token_url = "https://discordapp.com/api/oauth2/token"
    discord_api_url = "https://discordapp.com/api"

    @staticmethod
    def get_access_token(code):
        payload = {
            "client_id": Oauth.client_id,
            "client_secret": Oauth.client_secret,
            "grant_type": "authorization_code",
            "code": code,
            "redirect_uri": Oauth.redirect_uri,
            "scope": Oauth.scope
        }
        headers = {
            'Content-Type': 'application/x-www-form-urlencoded'
        }
        access_token = requests.post(url = Oauth.discord_token_url, data=payload, headers=headers)
        return access_token.json().get("access_token")

    @staticmethod
    def get_user_json(access_token):
        url = Oauth.discord_api_url+"/users/@me"

        headers = {
            "Authorization": "Bearer {}".format(access_token)
        }

        user = requests.get(url = url, headers = headers)
        return user.json()

