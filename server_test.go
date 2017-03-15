package main

import "testing"

import "net/http"
import "encoding/json"


func JsonQuery( t *testing.T, url string, result interface{} ) {
  resp,err := http.Get(url)
  if err != nil {
    t.Errorf("Got error response from local server")
  }

  decoder := json.NewDecoder( resp.Body )

  err = decoder.Decode( result )

  if err != nil {
    t.Errorf("Got error decoding JSON RootMap: %s", err.Error() )
  }
}


func TestRoot(t *testing.T) {
  server := RunOOIServer( "127.0.0.1", 8080 )
  defer server.Stop()

  rootmap := make( map[string]interface{} )
  JsonQuery(t, "http://127.0.0.1:8080/", &rootmap )

  if len(rootmap) == 0 {
    t.Error("Zero-length RootMap")
  }
}


func TestOOIRoot(t *testing.T) {
  server := RunOOIServer( "127.0.0.1", 8080 )
  defer server.Stop()

  rootmap := make( map[string]interface{} )
  JsonQuery(t, "http://127.0.0.1:8080/v1/org/oceanobservatories/rawdata/files/", &rootmap )

  if rootmap["Files"] == nil {
    t.Errorf("Path=\"Files\" doesn't exist")
    }

  if rootmap["Path"] != "/" {
    t.Errorf("Expected Path=\"/\", got: %s", rootmap["Path"])
  }
}
