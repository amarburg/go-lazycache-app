
FROM amarburg/golang-ffmpeg:wheezy-1.8

## Ideally this is cache busting...

COPY ./.gitversion $GOPATH

## Copy main so we use the version from the "pushed" copy, not pulling from Github
COPY main.go $GOPATH/src/github/amarburg/go-lazycache-app/

RUN go get github.com/amarburg/go-lazycache-app
RUN go install github.com/amarburg/go-lazycache-app

CMD ["/go/bin/go-lazycache-app", "--port", "80"]
EXPOSE 80
