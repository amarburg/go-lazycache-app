package main

import (
	"github.com/spf13/viper"
	"github.com/amarburg/go-lazycache"
	"net/http"
	"github.com/amarburg/go-stoppable-http-server"
	"fmt"
)


var OOIRawDataRootURL = "https://rawdata.oceanobservatories.org/files/"


func StartLazycacheServer(bind string, port int) *stoppable_http_server.SLServer {
	http.DefaultServeMux = http.NewServeMux()

	lazycache.RegisterDefaultHandlers()

	lazycache.DefaultLogger.Log("msg",fmt.Sprintf("Listening on http://%s:%d/",bind, port))

	server := stoppable_http_server.StartServer(func(config *stoppable_http_server.HttpConfig) {
		config.Host = bind
		config.Port = port
	} )

	return server
}

func main() {
	lazycache.ViperConfiguration()

	lazycache.ConfigureImageStoreFromViper( lazycache.DefaultLogger )
	lazycache.AddMirror(OOIRawDataRootURL)

	server := StartLazycacheServer(viper.GetString("bind"), viper.GetInt("port") )
	defer server.Stop()

	server.Wait()
}
