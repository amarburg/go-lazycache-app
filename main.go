package main

import (
	kitlog "github.com/go-kit/kit/log"
	"github.com/spf13/viper"
	flag "github.com/spf13/pflag"
	"os"
	"github.com/amarburg/go-lazycache"
	"net/http"
	"github.com/amarburg/go-stoppable-http-server"
	"fmt"
	"strings"
)

var OOIRawDataRootURL = "https://rawdata.oceanobservatories.org/files/"


func StartLazycacheServer(bind string, port int) *stoppable_http_server.SLServer {
	http.DefaultServeMux = http.NewServeMux()
	http.HandleFunc("/", lazycache.IndexHandler)

	DefaultLogger.Log("msg",fmt.Sprintf("Listening on http://%s:%d/",bind, port))

	server := stoppable_http_server.StartServer(func(config *stoppable_http_server.HttpConfig) {
		config.Host = bind
		config.Port = port
	} )

	return server
}

var DefaultLogger kitlog.Logger

func main() {

	// Configuration
	viper.SetDefault("port", 8080 )
	viper.SetDefault("bind","0.0.0.0")
	viper.SetDefault("imagestore", "")
	viper.SetDefault("imagestore.bucket","camhd-image-cache")

	viper.SetConfigName("lazycache")
	viper.AddConfigPath("/etc/lazycache")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil  { // Handle errors reading the config file
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			// ignore
		default:
	    panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}

	viper.SetEnvPrefix("lazycache")
	viper.AutomaticEnv()
	// Convert '.' to '_' in configuration variable names
	viper.SetEnvKeyReplacer( strings.NewReplacer(".", "_") )

	// var (
	// 	bindFlag          = flag.String("bind", "0.0.0.0", "Network interface to bind to (defaults to 0.0.0.0)")
	// 	ImageStoreFlag   = flag.String("image-store", "", "Type of image store (none, google)")
	// 	ImageBucketFlag = flag.String("image-store-bucket", "", "Bucket used for Google image store")
	// )
	flag.Int("port", 80, "Network port to listen on (default: 8080)")
	flag.String("bind", "0.0.0.0", "Network interface to bind to (defaults to 0.0.0.0)")
	flag.String("image-store", "", "Type of image store (none, google)")
	flag.String("image-store-bucket", "camhd-image-cache", "Bucket used for Google image store")

	viper.BindPFlag("port", flag.Lookup("port"))
	viper.BindPFlag("bind", flag.Lookup("bind"))
	viper.BindPFlag("imagestore", flag.Lookup("image-store"))
	viper.BindPFlag("imagestore.bucket", flag.Lookup("image-store-bucket"))

	flag.Parse()

	DefaultLogger = kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))

	//fmt.Println(config)
	lazycache.ConfigureImageStore(viper.GetString("imagestore"), viper.GetString("imagestore.bucket"), DefaultLogger)

	server := StartLazycacheServer(viper.GetString("bind"), viper.GetInt("port") )
	defer server.Stop()

	lazycache.AddMirror(OOIRawDataRootURL)

	server.Wait()
}
