FROM ubuntu:22.04

RUN apt-get update && \
    apt-get install -y software-properties-common \
                        net-tools \
                        vim \
                        jq \
                        curl \
                        wget

COPY metadata.json /opt
COPY root/ /

WORKDIR /root

ENTRYPOINT [ "/entrypoint.sh" ]
