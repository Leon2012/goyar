package goyar

import (
	"log"
	"net/http"
	"testing"
)

type Args struct {
	A, B int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	log.Println(args)

	*reply = args.A * args.B
	return nil
}

func TestNewRpc(t *testing.T) {
	log.Println("start server...")
	server := NewYarRpcServer()

	arith := new(Arith)
	server.Register(arith)

	http.Handle("/api", server)
	http.ListenAndServe(":8000", nil)
}
