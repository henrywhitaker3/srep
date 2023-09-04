FROM ghcr.io/henrywhitaker3/srep/base:1.0.2

COPY metadata.json /opt
COPY root/ /

RUN apt-get update && \
    apt-get install -y fuse-overlayfs \
                        busybox \
                        iptables

COPY --from=rancher/k3s:v1.27.4-k3s1 /bin/ /usr/bin/
