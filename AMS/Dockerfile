FROM flexconstructor/docker-adobe-media-server
MAINTAINER FlexConstructor <flexconstructor@gmail.com>
#-------- Copy FMS Serial key --------------

COPY serial_key /opt/adobe/ams/conf/serial_key
RUN echo SERVER.LICENSEINFO = $(cat /opt/adobe/ams/conf/serial_key) >> /opt/adobe/ams/conf/ams.ini
