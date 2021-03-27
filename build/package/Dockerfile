FROM golang:1.12-alpine AS go-builder

LABEL stage=build

ENV GO111MODULE on

#ARG GOLANG_PROXY_IP
#ENV GOPROXY=http://${GOLANG_PROXY_IP}:3000

WORKDIR /go/src/kubeshell

COPY ./kubeshell /go/src/kubeshell
RUN apk add git && \
    cd /go/src/kubeshell && \
    CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o kubeshell cmd/kubeshell/main.go

FROM alpine AS base
WORKDIR /kubeshell
COPY --from=go-builder /go/src/kubeshell/kubeshell /kubeshell
RUN apk add --no-cache util-linux
RUN wget -O /kubectl https://storage.googleapis.com/kubernetes-release/release/v1.20.0/bin/linux/amd64/kubectl && \
    chmod +x /kubectl && \
    cp /kubectl /usr/local/bin/

ENV port "5666"

ENTRYPOINT [ "/kubeshell/kubeshell" ]
