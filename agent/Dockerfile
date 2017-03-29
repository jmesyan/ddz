FROM golang:1.8
MAINTAINER mg <masterg@yeah.net>
WORKDIR /go/src/
COPY . /go/src/github.com/master-g/golandlord/agent
RUN go install -v github.com/master-g/golandlord/agent
ENTRYPOINT ["/go/bin/agent"]
EXPOSE 40001