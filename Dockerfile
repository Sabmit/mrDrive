FROM debian:wheezy
MAINTAINER Sabbagh Fares <sabbagh.fares@gmail.com>

ENV DEBIAN_FRONTEND noninteractive

RUN \
    echo "deb http://ppa.launchpad.net/webupd8team/java/ubuntu precise main" > /etc/apt/sources.list.d/webupd8team-java.list  && \
    apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys EEA14886  && \
    apt-get update &&                           \
    echo oracle-java8-installer shared/accepted-oracle-license-v1-1 select true | /usr/bin/debconf-set-selections  && \
    apt-get -yq install                         \
            oracle-java8-installer              \
            php5-curl                           \
            php5-cli                            \
            wget                                \
            gcc                                 \
            make                                \
            libevent-dev                        \
            libssl-dev                          \
            netcat                              \
            curl                                \
	    bzr 				\
	    git					\
	    mercurial && \
    rm -rf /etc/apt/sources.list.d/webupd8team-java.list

# Python modules
RUN apt-get -yq install --no-install-recommends \
    python-scipy                                \
    python-pandas                               \
    python-dev                                  \
    g++

ENV GOLANG_VERSION 1.3.3

RUN curl -sSL https://golang.org/dl/go$GOLANG_VERSION.src.tar.gz \
		| tar -v -C /usr/src -xz
RUN cd /usr/src/go/src && ./make.bash --no-clean 2>&1
ENV PATH /usr/src/go/bin:$PATH
ENV GOPATH /usr/src/go/

RUN mkdir -p /apps/vendor
WORKDIR /apps

RUN sed -i "s/variables_order.*/variables_order = \"EGPCS\"/g" /etc/php5/cli/php.ini
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer && \
    composer create-project --no-dev predis/predis /apps/vendor/predis

RUN cd /tmp && \
    wget --no-check-certificate https://download.elasticsearch.org/elasticsearch/elasticsearch/elasticsearch-1.4.2.tar.gz && \
    tar xvzf elasticsearch-1.4.2.tar.gz && \
    rm -f elasticsearch-1.4.2.tar.gz && \
    mv /tmp/elasticsearch-1.4.2 /elasticsearch

RUN cd /tmp/ && \
    wget --no-check-certificate https://www.torproject.org/dist/tor-0.2.5.10.tar.gz && \
    tar xzf tor-0.2.5.10.tar.gz && \
    cd tor-0.2.5.10 && ./configure && make install && \
    cd /tmp/ && rm -rf tor-0.2.5.10*

RUN cd /tmp && \
    git clone https://github.com/scikit-learn/scikit-learn.git && \
    cd scikit-learn && \
    python setup.py install --user && \
    cd && rm -rf /tmp/scikit-learn

RUN cd /tmp && \
    git clone https://github.com/networkx/networkx.git && \
    cd networkx && \
    python setup.py install --user && \
    cd && rm -rf /tmp/networkx

ADD http://sourceforge.net/projects/simplehtmldom/files/simple_html_dom.php/download /apps/vendor/simple_html_dom.php

COPY src/ startup.sh /apps/
COPY conf/torrc /usr/local/etc/tor/torrc

RUN chmod 755 /apps/*sh && /apps/startup.sh

#ONLY FOR TESTING PURPOSE
EXPOSE 9200
EXPOSE 9300

CMD ["/bin/bash"]
