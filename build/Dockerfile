FROM golang:1.9.4-alpine3.7 AS build-goman
WORKDIR /go/src/github.com/zaaksam/goman/go
COPY go .
RUN cd main/web && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o goman

FROM scratch
LABEL maintainer="zaaksam@qq.com"
COPY --from=build-goman /go/src/github.com/zaaksam/goman/go/main/web/goman /app/
EXPOSE 8080
ENTRYPOINT [ "/app/goman", "-ip", "0.0.0.0", "-port", "8080" ]