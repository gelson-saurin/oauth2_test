FROM golang:alpine AS build
RUN mkdir /go/src/oauth2-server
ADD ./ /go/src/oauth2-server
WORKDIR /go/src/oauth2-server
RUN apk --update upgrade && rm -rf /var/cache/apk/* && apk add build-base
RUN CGO_ENABLED=1 GOOS=linux go build -a -tags musl -installsuffix cgo -ldflags '-w -s -extldflags "-static"' -o /main cmd/main.go
EXPOSE 8080
CMD ["/main"]