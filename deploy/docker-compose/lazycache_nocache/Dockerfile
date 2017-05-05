FROM amarburg/golang-ffmpeg:wheezy-1.8

RUN go get -v github.com/amarburg/go-lazycache-app

#
## Hot-patch main.go from this directory into go-lazycache-app
#
ADD main.go $GOPATH/src/github.com/amarburg/go-lazycache-app/

RUN go build -v github.com/amarburg/go-lazycache-app
RUN go install -v github.com/amarburg/go-lazycache-app

VOLUME ["/srv/image_store"]
VOLUME ["/data"]

ENV LAZYCACHE_PORT=8080

## Strangely, it's installing the binary to $GOPATH
CMD $GOPATH/go-lazycache-app
EXPOSE 8080
