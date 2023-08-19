FROM henrywhitaker3/srep-base:latest

RUN apt-get update && \
    apt-get install -y python3 \
                        lsof

COPY files/ /opt
