package request

import (
	"bytes"
	"context"
	"sync"

	"github.com/webnice/transport/v3/header"
	"github.com/webnice/transport/v3/methods"
	"github.com/webnice/transport/v3/response"
)

// New creates a new object and return interface
func New() Pool {
	var rqt = new(impl)
	rqt.methods = methods.New()
	rqt.responsePool = response.New()
	rqt.requestPool = new(sync.Pool)
	rqt.requestPool.New = rqt.NewRequestItem
	return rqt
}

// NewRequestItem Конструктор sync.Pool для Request
func (rqt *impl) NewRequestItem() interface{} {
	var req = &Request{
		method:      rqt.methods.Get(),
		header:      header.New(),
		url:         &bytes.Buffer{},
		requestData: &bytes.Reader{},
		response:    rqt.responsePool.ResponseGet(),
	}
	req.context, req.contextCancelFunc = context.WithCancel(context.Background())
	return req
}

// RequestGet Извлечение из pool нового элемента Request
func (rqt *impl) RequestGet() Interface {
	var req = rqt.requestPool.Get().(*Request)
	req.response = rqt.responsePool.ResponseGet().DebugFunc(response.DebugFunc(req.debugFunc))
	return req
}

// RequestPut Возврат в sync.Pool использованного элемента Request
func (rqt *impl) RequestPut(req Interface) {
	rqt.requestClean(req.(*Request))
	rqt.requestPool.Put(req)
}

// Очистка данных объекта Request, подготовка к переиспользованию
func (rqt *impl) requestClean(req *Request) {
	req.context, req.contextCancelFunc = context.WithCancel(context.Background())
	rqt.responsePool.ResponsePut(req.response)
	req.response = nil
	req.method = rqt.methods.Get()
	req.header.Reset()
	req.err = nil
	req.debugFunc = nil
	req.url.Reset()
	req.request = nil
	req.requestData = &bytes.Reader{}
	req.requestDataInterface = nil
	req.username = req.username[:0]
	req.password = req.password[:0]
	req.cookie = req.cookie[:0]
	// Переменные для внутренних целей
	req.tmpArr, req.tmpOk, req.tmpCounter, req.tmpBytes = req.tmpArr[:0], false, 0, req.tmpBytes[:0]
}
