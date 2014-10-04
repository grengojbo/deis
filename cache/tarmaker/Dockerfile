FROM ubuntu:14.04

ENV DEBIAN_FRONTEND noninteractive

# install common packages
RUN apt-get update && apt-get install -y curl

# build Redis from source
RUN buildDeps='gcc libc6-dev make'; \
    set -x; \
    apt-get update && apt-get install -y $buildDeps net-tools --no-install-recommends \
    && mkdir -p /usr/src/redis \
    && mkdir -p /app \
    && mkdir -p /tmp/rootfs \
    && cd /usr/src \
    && curl -sSL http://download.redis.io/releases/redis-2.8.17.tar.gz -o redis.tar.gz \
    && tar -xzf redis.tar.gz -C /usr/src/redis --strip-components=1 \
    && cd /usr/src/redis \
    && make \
    && make PREFIX=/app install \
    && rm /app/bin/redis-benchmark \
    && cd /app/bin \
    && tar -czf /tmp/redis-dist.tar.gz . \
    && apt-get purge -y $buildDeps \
    && apt-get autoremove -y \
    && apt-get clean
