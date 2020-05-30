package main

import (
	"net/http"
	"time"

	"github.com/TarsCloud/TarsGo/tars"

)

func httpRootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write([]byte(time.Now().Local().Format("2006-01-02 15:04:05")))
	return
}

func main() {
	mux := &tars.TarsHttpMux{}
	mux.HandleFunc("/", httpRootHandler)
	cfg := tars.GetServerConfig()
	tars.AddHttpServant(mux, cfg.App+"."+cfg.Server+".SayHelloObj") //Register http server
	tars.Run()
}
