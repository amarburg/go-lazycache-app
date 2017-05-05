package main

import (
	"github.com/spf13/viper"
	"github.com/amarburg/go-lazycache"
	"net/http"
	"github.com/amarburg/go-stoppable-http-server"
	"fmt"
)



var ooiRawDataRootURL = "https://rawdata.oceanobservatories.org/files/"


// StartLazycacheServer starts a HTTP server at http://bind:port/ and registers default Lazycache handlers
func StartLazycacheServer(bind string, port int) *stoppable_http_server.SLServer {
	http.DefaultServeMux = http.NewServeMux()

	msg := fmt.Sprintf("Listening on http://%s:%d/",bind, port)
	lazycache.DefaultLogger.Log("msg", msg)

	server := stoppable_http_server.StartServer(func(config *stoppable_http_server.HttpConfig) {
		config.Host = bind
		config.Port = port
	} )

	lazycache.RegisterDefaultHandlers()

	return server
}

// RunOOIServer start an Lazycache server and registers the standard Rutgers CI destination
func RunOOIServer( bind string, port int ) (*stoppable_http_server.SLServer) {
	server := StartLazycacheServer( bind, port )

	lazycache.AddMirror(ooiRawDataRootURL)

	// lazycache.AddMirror( "http://nginx_data/" )
	//
	// // Berna-specific configuration (hardcoded for now)
	// fs,err := lazycache.OpenLocalFS("/data")
	//
	// if err != nil {
	// 	panic(fmt.Sprintf("Error opening Local File Source: %s", err.Error()))
	// }
	//
	// lazycache.MakeRootNode(fs, fmt.Sprintf("/v1/berna%s/",fs.Path) )

	return server
}


func main() {
	lazycache.ConfigureFromViper()

	server := RunOOIServer( viper.GetString("bind"), viper.GetInt("port") )
	defer server.Stop()

	// Handle Google App Engine health requests
	http.HandleFunc("/_ah/health", healthCheckHandler)

	server.Wait()
}
