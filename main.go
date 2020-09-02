package main

import (
	"fmt"
	"github.com/yuwe1/delaynet/listen"
	"github.com/yuwe1/delaynet/ping"
	"github.com/yuwe1/delaynet/route"
	"github.com/yuwe1/delaynet/server"
	"net/http"

	_ "net/http/pprof"
)

func main(){
	l ,_:=listen.RetryListen("tcp", "0.0.0.0:2332")
	srv :=server.Server{
		Server :&http.Server{},
	}
	route :=route.NewRoute()
	route.HandleFunc("/latency/{delaytime}",srv.GetServerHttp).Methods("GET")
	srv.Server.Handler =route
	go ping.P.Listen()
	fmt.Println("start.....")
	srv.Server.Serve(l)
}