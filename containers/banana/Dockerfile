FROM ghcr.io/henrywhitaker3/srep/base:1.0.2

RUN add-apt-repository ppa:redislabs/redis

RUN apt-get update && \
    apt-get install -y redis

COPY metadata.json /opt
COPY root/ /
