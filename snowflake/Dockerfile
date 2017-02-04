FROM golang:latest
MAINTAINER mg <masterg@yeah.net>
COPY . /go/src/snowflake
RUN go install snowflake
ENTRYPOINT /go/bin/snowflake
EXPOSE 40001
