package transport

import (
	"context"
	"fmt"
	runtimeDebug "runtime/debug"

	"github.com/webnice/transport/v4/request"
)

// Создание и запуск бассейна работников для обслуживания запросов.
func (trt *impl) makePool() {
	var (
		i             uint16
		ctx           context.Context
		ctxCancelFunc context.CancelFunc
	)

	// Если процессы бассейна запущены, тогда выход.
	if trt.requestPoolStarted.Load().(bool) {
		return
	}
	// Защита от двойного запуска.
	trt.requestPoolLock.Lock()
	defer trt.requestPoolLock.Unlock()
	// Запуск процессов работников.
	trt.requestPoolCancelFunc = make([]context.CancelFunc, 0, trt.requestPoolSize)
	for i = 0; i < trt.requestPoolSize; i++ {
		ctx, ctxCancelFunc = context.WithCancel(context.Background())
		trt.requestPoolCancelFunc = append(trt.requestPoolCancelFunc, ctxCancelFunc)
		trt.requestPoolDone.Add(1)
		go trt.poolWorker(ctx)
	}
	trt.requestPoolStarted.Store(true)
}

// Процесс работника бассейна.
func (trt *impl) poolWorker(ctx context.Context) {
	var req request.Interface

	defer trt.requestPoolDone.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case req = <-trt.requestChan:
			if req == nil {
				continue
			}
		}
		trt.request(req)
	}
}

// Выполнение запроса.
func (trt *impl) request(req request.Interface) {
	var err error

	defer func() {
		if e := recover(); e != nil {
			trt.err = fmt.Errorf("Catch panic: %s\nGoroutine stack is:\n%s", e.(error), string(runtimeDebug.Stack()))
			if trt.errFunc != nil {
				trt.errFunc(trt.err)
			}
			return
		}
	}()
	if err = req.Do(trt.client); err == nil {
		return
	}
	if trt.err = err; trt.errFunc != nil {
		trt.errFunc(err)
	}
}
