FROM amarburg/golang-ffmpeg:wheezy-1.8

RUN go get github.com/amarburg/go-lazycache-app

## Copy main so we use the version from the "pushed" copy, not pulling from Github
ADD main.go $GOPATH/src/github.com/amarburg/go-lazycache-app/

RUN go build github.com/amarburg/go-lazycache-app
RUN go install github.com/amarburg/go-lazycache-app

VOLUME ["/srv/image_store"]
VOLUME ["/data"]

ENV LAZYCACHE_PORT=8080 \
    LAZYCACHE_IMAGESTORE=local  \
    LAZYCACHE_IMAGESTORE_LOCALROOT=/srv/image_store \
    LAZYCACHE_IMAGESTORE_URLROOT=file:///tmp/image_store

## Strangely, it's installing the binary to $GOPATH
CMD $GOPATH/go-lazycache-app
EXPOSE 8080
