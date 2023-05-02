FROM ubuntu:18.04

RUN apt-get update && \
apt-get install -y redis-server && \
apt-get clean

EXPOSE 6379

ENTRYPOINT ["redis-server", "--protected-mode no"]
