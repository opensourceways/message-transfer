FROM golang:latest as BUILDER
LABEL maintainer="shishupei"

ARG USER
ARG PASS
RUN echo "machine github.com login $USER password $PASS" >/root/.netrc
# build binary
RUN mkdir -p /go/src/github.com/opensourceways/message-transfer
COPY . /go/src/github.com/opensourceways/message-transfer
RUN cd /go/src/github.com/opensourceways/message-transfer && CGO_ENABLED=1 go build -v -o ./message-transfer main.go

# copy binary config and utils
FROM openeuler/openeuler:22.03

COPY --from=BUILDER /go/src/github.com/opensourceways/message-transfer /opt/app
WORKDIR /opt/app/
ENTRYPOINT ["/opt/app/message-transfer"]