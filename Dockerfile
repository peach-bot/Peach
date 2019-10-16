FROM python:3.7

ENV BOTTOKEN=<insert_discord_token_here>

ENV DBHOST=<insert_db_hostname>

ENV DBUSER=<insert_db_user>

ENV DBPASSWORD=<insert_db_password>

EXPOSE 5000

RUN apt update

RUN apt install nodejs npm -y

RUN npm i -g yarn

WORKDIR /app

COPY . /app

RUN pip3 install -r requirements.txt

RUN yarn

CMD [ "yarn", "dev" ]