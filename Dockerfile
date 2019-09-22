FROM alpine:latest

RUN apk add --no-cache python3-dev=3.7.3-r0 \
    && pip3 install --upgrade pip

WORKDIR /app

COPY . /app

RUN pip3 --no-cache-dir install -r requirements.txt --trusted-host pypi.python.org                                                     

EXPOSE 5000

CMD start.sh
