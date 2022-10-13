FROM alpine:3.14
MAINTAINER devendra.turkar@gmail.com

WORKDIR /sample
COPY main main

ENTRYPOINT ["./main"]