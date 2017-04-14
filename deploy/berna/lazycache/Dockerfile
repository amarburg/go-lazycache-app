FROM amarburg/golang-ffmpeg:wheezy-1.8

RUN go get github.com/amarburg/go-lazycache-app

#
## Hot-patch main.go from this directory into go-lazycache-app
#
ADD main.go $GOPATH/src/github.com/amarburg/go-lazycache-app/

RUN ls -al $GOPATH/src/github.com/amarburg/go-lazycache-app
RUN go build -x -v github.com/amarburg/go-lazycache-app
RUN go install github.com/amarburg/go-lazycache-app

VOLUME ["/srv/image_store"]
VOLUME ["/data"]

ENV LAZYCACHE_PORT=8080 \
    LAZYCACHE_IMAGESTORE=local  \
    LAZYCACHE_IMAGESTORE_ROOT=/srv/image_store \
    LAZYCACHE_IMAGESTORE_URL=http://localhost:7082/

## Strangely, it's installing the binary to $GOPATH
CMD $GOPATH/go-lazycache-app
EXPOSE 8080
