FROM golang:1.8-onbuild
MAINTAINER mg <masterg@yeah.net>
COPY . /go/src/github.com/master-g/golandlord/agent
RUN go install github.com/master-g/golandlord/agent
ENTRYPOINT /go/bin/agent
EXPOSE 8888 8888/udp
