FROM debian:wheezy
MAINTAINER Sabbagh Fares <sabbagh.fares@gmail.com>

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update &&                           \
    apt-get -yq install                         \
            php5-curl                           \
            php5-cli                            \
            wget                                \
            gcc                                 \
            make                                \
            libevent-dev                        \
            libssl-dev                          \
            netcat                              \
            curl


RUN mkdir -p /apps/vendor
WORKDIR /apps

RUN sed -i "s/variables_order.*/variables_order = \"EGPCS\"/g" /etc/php5/cli/php.ini
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer && \
    composer create-project --no-dev predis/predis /apps/vendor/predis

RUN wget --no-check-certificate https://www.torproject.org/dist/tor-0.2.5.10.tar.gz && \
    tar xzf tor-0.2.5.10.tar.gz && \
    cd tor-0.2.5.10 && ./configure && make install && \
    cd /apps/ && rm -rf tor-0.2.5.10*

ADD http://sourceforge.net/projects/simplehtmldom/files/simple_html_dom.php/download /apps/vendor/simple_html_dom.php

COPY src/ startup.sh /apps/
COPY conf/torrc /usr/local/etc/tor/torrc

RUN chmod 755 /apps/*sh && /apps/startup.sh

CMD ["/bin/bash"]
