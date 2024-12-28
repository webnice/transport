package request

import (
	"bytes"
	"context"
	"sync"

	"github.com/webnice/dic"
	"github.com/webnice/transport/v4/header"
	"github.com/webnice/transport/v4/response"
)

// New Конструктор объекта сущности пакета, возвращается интерфейс пакета.
func New() Pool {
	var rqt = new(impl)
	rqt.responsePool = response.New()
	rqt.requestPool = new(sync.Pool)
	rqt.requestPool.New = rqt.NewRequestItem

	return rqt
}

// NewRequestItem Конструктор объектов бассейна для Request.
func (rqt *impl) NewRequestItem() interface{} {
	var req = &Request{
		method:      dic.Method().Get,
		header:      header.New(),
		uri:         &bytes.Buffer{},
		requestData: &bytes.Reader{},
		response:    rqt.responsePool.ResponseGet(),
	}
	req.context, req.contextCancelFunc = context.WithCancel(context.Background())

	return req
}

// RequestGet Извлечение из бассейна нового объекта Request.
func (rqt *impl) RequestGet() Interface {
	var req = rqt.requestPool.Get().(*Request)
	req.response = rqt.responsePool.ResponseGet().DebugFunc(response.DebugFunc(req.debugFunc))
	return req
}

// RequestPut Возврат в бассейн элемента Request.
func (rqt *impl) RequestPut(req Interface) {
	rqt.requestClean(req.(*Request))
	rqt.requestPool.Put(req)
}

// Очистка данных объекта Request, подготовка к возврату в бассейн для повторного использования.
func (rqt *impl) requestClean(req *Request) {
	req.context, req.contextCancelFunc = context.WithCancel(context.Background())
	rqt.responsePool.ResponsePut(req.response)
	req.response = nil
	req.method = dic.Method().Get
	req.header.Reset()
	req.err = nil
	req.debugFunc = nil
	req.uri.Reset()
	req.request = nil
	req.requestData = &bytes.Reader{}
	req.requestDataInterface = nil
	req.username = req.username[:0]
	req.password = req.password[:0]
	req.cookie = req.cookie[:0]
	// Переменные для внутренних целей.
	req.tmpArr, req.tmpOk, req.tmpCounter, req.tmpBytes = req.tmpArr[:0], false, 0, req.tmpBytes[:0]
}
