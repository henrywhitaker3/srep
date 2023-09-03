FROM ubuntu:22.04

RUN apt-get update && \
    apt-get install -y software-properties-common \
                        net-tools \
                        vim \
                        jq

COPY metadata.json /opt
COPY root/ /

RUN echo "/opt/message.sh" >> /root/.bashrc

ENTRYPOINT [ "/entrypoint.sh" ]
