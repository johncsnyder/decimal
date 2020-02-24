# -*- mode: dockerfile -*-
FROM golang:1.12.2-stretch

RUN apt-get update && apt-get install -y unzip

# See https://gist.github.com/sofyanhadia/37787e5ed098c97919b8c593f0ec44d8
# The version in apt-get is for proto2 only.
WORKDIR /tmp
ARG PROTOC_VERSION=3.6.1
RUN curl -OL https://github.com/google/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip && \
    unzip protoc-${PROTOC_VERSION}-linux-x86_64.zip -d protoc3 && \
    mv protoc3/bin/* /usr/local/bin/ && \
    mv protoc3/include/* /usr/local/include/ && \
    rm -r protoc3 && rm protoc-${PROTOC_VERSION}-linux-x86_64.zip

RUN go get -u github.com/golang/protobuf/proto github.com/golang/protobuf/protoc-gen-go github.com/gogo/protobuf/jsonpb github.com/gogo/protobuf/protoc-gen-gogo github.com/gogo/protobuf/protoc-gen-gofast github.com/gogo/protobuf/protoc-gen-gogofast github.com/gogo/protobuf/protoc-gen-gogofaster github.com/gogo/protobuf/protoc-gen-gogoslick github.com/gogo/protobuf/gogoproto

ENTRYPOINT ["protoc"]
