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
RUN dnf -y update && \
    dnf in -y shadow && \
    groupadd -g 1000 message && \
    useradd -u 1000 -g defect -s /bin/bash -m message

USER message

COPY  --chown=defect --from=BUILDER /go/src/github.com/opensourceways/message-transfer /opt/app
WORKDIR /opt/app/
ENTRYPOINT ["/opt/app/message-transfer"]