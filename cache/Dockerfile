FROM            scratch
MAINTAINER      MAINTAINER Oleg Dolya "oleg.dolya@gmail.com"

RUN mkdir -p /app/bin
RUN mkdir -p /data
RUN mkdir -p /run/redis

EXPOSE 6379
WORKDIR /app

VOLUME ["/data"]

ADD rootfs.tar /

ADD bin/test /app/bin/test
ADD bin/boot /app/bin/boot

ADD redis-cli /app/bin/redis-cli
ADD redis-server /app/bin/redis-server
#ADD redis-check-aof /app/bin/redis-check-aof
#ADD redis-check-dump /app/bin/redis-check-dump

# set timezone /usr/share/zoneinfo/Europe/Kiev
#ADD localtime  /etc/localtime
#ADD timezone /etc/timezone

# install etcdctl
RUN curl -sSL -o /usr/local/bin/etcdctl https://s3-us-west-2.amazonaws.com/opdemand/etcdctl-v0.4.6 \
    && chmod +x /usr/local/bin/etcdctl

CMD ["/app/bin/boot"]
ADD redis.conf /etc/redis/redis.conf
