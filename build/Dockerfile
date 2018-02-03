FROM golang:1.9.3-alpine3.7 AS build-goman
WORKDIR /go/src/github.com/zaaksam/goman/go
COPY go .
RUN cd main/web && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o goman

FROM scratch
LABEL maintainer="zaaksam@qq.com"
WORKDIR /app
COPY --from=build-goman /go/src/github.com/zaaksam/goman/go/main/web/goman /app/
EXPOSE 8080
ENTRYPOINT [ "./goman", "-ip", "0.0.0.0", "-port", "8080" ]

# cd .. && docker build -t goman:latest -t goman:0.2.2 -f ./build/Dockerfile .