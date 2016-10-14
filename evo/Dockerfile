FROM ubuntu:14.04
MAINTAINER FlexConstructor <flexconstructor@gmail.com>

ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update && apt-get install -y curl

ENV EVO_STREAM_BUILD 4491
ENV EVO_STREAM_VERSION 1.7.1
ENV DISTRIBUTION_VERSION 14.04
RUN mkdir -p /src && \
    curl -sSL http://tarballs.evostream.com/release/$EVO_STREAM_BUILD/evostreamms-$EVO_STREAM_VERSION.$EVO_STREAM_BUILD-x86_64-Ubuntu_$DISTRIBUTION_VERSION.tar.gz | tar -xzC /src && \
    ln -s /src/evostreamms-$EVO_STREAM_VERSION.$EVO_STREAM_BUILD-x86_64-Ubuntu_$DISTRIBUTION_VERSION /src/evostreamms
COPY ./License.lic /src/evostreamms/config/License.lic

WORKDIR /src/evostreamms/bin
VOLUME /src/evostreamms/config
#EXPOSE 1935 5544 6666 8080
CMD ["./run_console_ems.sh"]