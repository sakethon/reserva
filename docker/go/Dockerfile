FROM golang:1.15.0-alpine3.12
RUN apk add --update --no-cache ca-certificates git build-base gcc && \
    wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.29.0 && \
    mkdir /src
WORKDIR /src
