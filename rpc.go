package goyar

import (
	_ "errors"
	"log"
	"net/http"
	"net/rpc"
)

type YarRpcServer struct {
	*rpc.Server
}

func NewYarRpcServer() *YarRpcServer {
	return &YarRpcServer{rpc.NewServer()}
}

func (s *YarRpcServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Println("got a request")
	conn, _, err := w.(http.Hijacker).Hijack()
	if err != nil {
		log.Print("rpc hijacking ", req.RemoteAddr, ": ", err.Error())
		return
	}

	codec := NewServerCodec(conn, w, req)
	log.Println("ServeCodec")

	s.Server.ServeCodec(codec)
	log.Println("finished serving request")
}
