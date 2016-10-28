FROM ubuntu:14.04
MAINTAINER FlexConstructor <flexconstructor@gmail.com>
ENV TERM linux
RUN apt-get -y update                                                                            \
    && apt-get install -y wget python-pip                                                        \
    && wget -q -O - http://debian.erlyvideo.org/binary/gpg.key | apt-key add -                   \
    && echo "deb http://debian.erlyvideo.org binary/" > /etc/apt/sources.list.d/erlyvideo.list   \
    && apt-get -y update                                                                         \
    && apt-get -y install flussonic flussonic-ffmpeg flussonic-python flussonic-erlang           \
    && pip install supervisor                                                                    \
    && apt-get clean autoclean                                                                   \
    && apt-get autoremove -y                                                                     \
    && rm -rf /var/lib/{apt,dpkg,cache,log}
COPY license.txt  /etc/flussonic/license.txt
COPY flussonic.conf /etc/flussonic/flussonic.conf
ADD supervisord.conf /etc/supervisord.conf
VOLUME ["/var/log/flussonic"]

EXPOSE 80 8080 1935 554
CMD ["/usr/local/bin/supervisord"]