FROM ubuntu:22.04

RUN apt-get update && \
    apt-get install -y software-properties-common \
                        net-tools \
                        vim

COPY root/ /

ENTRYPOINT [ "/entrypoint.sh" ]
