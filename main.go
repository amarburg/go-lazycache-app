package main

import (
	"flag"
	kitlog "github.com/go-kit/kit/log"
	"os"
	"github.com/amarburg/go-lazycache"
)

var OOIRawDataRootURL = "https://rawdata.oceanobservatories.org/files/"

func main() {

	var (
		port          = flag.Int("port", 5000, "Network port to listen on (default: 5000)")
		bind          = flag.String("bind", "0.0.0.0", "Network interface to bind to (defaults to 0.0.0.0)")
		image_store   = flag.String("image-store", "", "Type of image store (none, google)")
		google_bucket = flag.String("image-store-bucket", "", "Bucket used for Google image store")
	)
	flag.Parse()

	defaultLogger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))

	//fmt.Println(config)
	lazycache.ConfigureImageStore(*image_store, *google_bucket, defaultLogger)

	server := lazycache.StartLazycacheServer(*bind, *port)
	defer server.Stop()

	lazycache.AddMirror(OOIRawDataRootURL)

	server.Wait()
}
