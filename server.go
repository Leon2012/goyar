package goyar

import (
	_ "errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/rpc"
	"sync"
)

type yarCodec struct {
	rwc     io.ReadWriteCloser
	rw      http.ResponseWriter
	request *http.Request
	mutex   sync.Mutex
	seq     uint64
	yar     *Yar
}

func NewServerCodec(conn io.ReadWriteCloser, w http.ResponseWriter, req *http.Request) rpc.ServerCodec {
	return &yarCodec{
		rwc:     conn,
		rw:      w,
		request: req,
	}
}

func (c *yarCodec) ReadRequestHeader(r *rpc.Request) error {
	data, err := ioutil.ReadAll(c.request.Body)
	fmt.Println("data:", data)
	if err != nil {
		return err
	}

	yar, err := Unpack(data)
	fmt.Println(yar)
	if err != nil {
		return err
	}

	r.ServiceMethod = yar.Request.Method
	r.Seq = yar.Request.Id
	c.yar = yar

	return nil
}

func (c *yarCodec) ReadRequestBody(body interface{}) error {
	body = c.yar.Request.Params[0]
	return nil
}

type flusher interface {
	Flush() error
}

func (c *yarCodec) WriteResponse(r *rpc.Response, x interface{}) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return nil
}

func (c *yarCodec) Close() error {
	return c.rwc.Close()
}

// func ServeConn(conn io.ReadWriteCloser) {
// 	rpc.ServeCodec(NewServerCodec(conn))
// }
