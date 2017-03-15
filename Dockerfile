FROM amarburg/golang-ffmpeg:wheezy-1.8

## Ideally this is cache busting...
#COPY ./.gitversion $GOPATH

RUN go get github.com/amarburg/go-lazycache-app

## Copy main so we use the version from the "pushed" copy, not pulling from Github
COPY main.go $GOPATH/src/github.com/amarburg/go-lazycache-app

RUN go install github.com/amarburg/go-lazycache-app

VOLUME ["/srv/image_store"]

ENV LAZYCACHE_IMAGESTORE=local  \
    LAZYCACHE_IMAGESTORE_LOCALROOT=/srv/image_store \
    LAZYCACHE_IMAGESTORE_URLROOT=file:///tmp/image_store

CMD ["/go/bin/go-lazycache-app", "--port", "8080"]
EXPOSE 8080
