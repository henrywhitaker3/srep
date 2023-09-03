FROM ghcr.io/henrywhitaker3/srep/base:1.0.1

RUN apt-get update && \
    apt-get install -y python3 \
                        lsof

COPY metadata.json /opt
COPY root/ /
