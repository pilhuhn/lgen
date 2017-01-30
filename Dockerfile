FROM alpine
MAINTAINER <hwr@pilhuhn.de>

EXPOSE 8080
COPY lgen /lgen
ENTRYPOINT /lgen
