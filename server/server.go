package server

import (
	"github.com/gorilla/mux"
	"github.com/yuwe1/delaynet/ping"
	"net/http"
	"strconv"
)

type Server struct {
	Server        *http.Server
}

func (s *Server)GetServerHttp(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["delaytime"]
	var delayTime int64
	var err error
	if delayTime,err = strconv.ParseInt(id[:len(id)-2],10,64);err!=nil{
		w.Write([]byte(err.Error()))
		return
	}

	if ping.P.DelayTime != 0{
		ping.P.Rescyle <-struct{}{}
		<-ping.P.CtlTime
	}
	ping.P.DelayTime = delayTime
	ping.P.CtlTime  <- struct{}{}
	w.Write([]byte(id))
}