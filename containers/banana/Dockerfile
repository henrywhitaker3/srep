FROM ghcr.io/henrywhitaker3/srep/base:latest

RUN add-apt-repository ppa:redislabs/redis

RUN apt-get update && \
    apt-get install -y redis

COPY root/ /
