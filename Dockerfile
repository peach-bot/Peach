FROM python:3.7

RUN apt update

RUN apt install nodejs -y

WORKDIR /app

COPY . /app

RUN pip3 --no-cache-dir install -r requirements.txt

RUN npm i -g nodemon

EXPOSE 5000

ENTRYPOINT [ "npm", "run", "dev" ]