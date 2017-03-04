FROM golang:1.8
MAINTAINER mg <masterg@yeah.net>
WORKDIR /go/src/
COPY . /go/src/github.com/master-g/golandlord/snowflake
RUN go install -v github.com/master-g/golandlord/snowflake
ENTRYPOINT ["/go/bin/snowflake"]
EXPOSE 40001
