FROM grafana/grafana:master
MAINTAINER FlexConstructor <flexconstructor@gmail.com>
ENV DASHBOARDS_JSON_ENABLED true
RUN mkdir /var/lib/grafana/dashboards \
    && chmod 775 /var/lib/grafana/dashboards
COPY dashboards/docker-dashboard_rev5.json /var/lib/grafana/dashboards/docker-dashboard_rev5.json