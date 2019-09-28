FROM python:latest

WORKDIR /usr/src/app

COPY . .

RUN pip3 --no-cache-dir install -r requirements.txt --trusted-host pypi.python.org                                                     

EXPOSE 5000

CMD ./start.sh