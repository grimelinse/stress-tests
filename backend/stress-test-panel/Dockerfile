FROM andreptb/grunt
MAINTAINER FlexConstructor <flexconstructor@gmail.com>
RUN mkdir -p /var/lib/grafana/plugins/stress-test-panel
WORKDIR /var/lib/grafana/plugins/stress-test-panel
COPY . /var/lib/grafana/plugins/stress-test-panel
#
RUN npm install \
    && grunt    \
    && mkdir -p /usr/share/grafana/public/redis
CMD echo ${REDIS_CLIENT_URL} > /usr/share/grafana/public/redis/config.txt
VOLUME /var/lib/grafana/plugins
VOLUME /usr/share/grafana/public/redis