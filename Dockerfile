
FROM amarburg/golang-ffmpeg:wheezy-1.8

## Ideally this is cache busting...
ENV dummy=$IMG_NAME

RUN go get github.com/amarburg/go-lazycache-app
RUN go install github.com/amarburg/go-lazycache-app

CMD ["/go/bin/go-lazycache-app", "--port", "80"]
EXPOSE 80
