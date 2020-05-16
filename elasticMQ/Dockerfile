FROM java:8

WORKDIR /var/elasticMQ

EXPOSE 9324

RUN wget -q https://s3-eu-west-1.amazonaws.com/softwaremill-public/elasticmq-server-0.13.9.jar

COPY custom.conf /var/elasticMQ

ENTRYPOINT ["/usr/bin/java", "-Dconfig.file=custom.conf", "-jar", "elasticmq-server-0.13.9.jar"]