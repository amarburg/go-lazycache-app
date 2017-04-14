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

	msg := fmt.Sprintf("Listening on http://%s:%d/",bind, port)
	lazycache.DefaultLogger.Log("msg", msg)

	server := stoppable_http_server.StartServer(func(config *stoppable_http_server.HttpConfig) {
		config.Host = bind
		config.Port = port
	} )

	return server
}

func RunOOIServer( bind string, port int ) (*stoppable_http_server.SLServer) {
	server := StartLazycacheServer( bind, port )

	lazycache.RegisterDefaultHandlers()
	lazycache.AddMirror(OOIRawDataRootURL)

		lazycache.AddMirror( "http://localhost:7081/" )

	return server
}


func main() {
	lazycache.ConfigureFromViper()

	server := RunOOIServer( viper.GetString("bind"), viper.GetInt("port") )

	http.HandleFunc("/_ah/health", healthCheckHandler)

	defer server.Stop()
	server.Wait()
}
