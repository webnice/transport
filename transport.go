package transport

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	runtimeDebug "runtime/debug"
	"sync"
	"sync/atomic"

	"github.com/webnice/transport/v4/request"
)

// New Конструктор объекта сущности пакета, возвращается интерфейс пакета.
func New() Interface {
	var trt = new(impl)
	trt.requestPoolInterface = request.New()
	trt.requestChan = make(chan request.Interface, requestChanBuffer)
	trt.requestPoolLock = new(sync.Mutex)
	trt.requestPoolStarted = new(atomic.Value)
	trt.requestPoolStarted.Store(false)
	trt.requestPoolDone = new(sync.WaitGroup)
	setDefaults(trt)
	return trt
}

// Error Последняя ошибка.
func (trt *impl) Error() error { return trt.err }

// ErrorFunc Регистрация функции получения ошибок транспорта.
func (trt *impl) ErrorFunc(fn ErrorFunc) Interface { trt.errFunc = fn; return trt }

// DebugFunc Включение или отключение режима отладки.
// Если передана функция отладки не равная nil, режим отладки включается.
// Передача функции отладки равной nil отключает режим отладки.
func (trt *impl) DebugFunc(fn DebugFunc) Interface { trt.debugFunc = fn; return trt }

// RequestGet Получение из бассейна объекта request.Interface.
// Полученный объект обязательно необходимо вернуть в бассейн методом RequestPut для избежания утечки памяти.
func (trt *impl) RequestGet() request.Interface {
	return trt.requestPoolInterface.RequestGet().DebugFunc(request.DebugFunc(trt.debugFunc))
}

// RequestPut Возвращение в бассейн объекта request.Interface.
func (trt *impl) RequestPut(req request.Interface) { trt.requestPoolInterface.RequestPut(req) }

// Client Получение клиента http.Client.
// В пределах одного экземпляра transport.impl, http.Client создаётся только один раз
// при первом вызове данной функции. Эта функция так же вызывается при первом вызове функции Do().
func (trt *impl) Client() (ret *http.Client) {
	if trt.client != nil {
		return trt.client
	}
	if trt.tlsClientConfig == nil && trt.tlsInsecureSkipVerify {
		trt.tlsClientConfig = &tls.Config{InsecureSkipVerify: trt.tlsInsecureSkipVerify}
	} else if trt.tlsClientConfig != nil && trt.tlsInsecureSkipVerify {
		trt.tlsClientConfig.InsecureSkipVerify = trt.tlsInsecureSkipVerify
	}
	// Создание объекта транспорта.
	if trt.transport == nil {
		trt.transport = &http.Transport{
			Proxy:               trt.proxy,
			ProxyConnectHeader:  trt.proxyConnectHeader,
			MaxIdleConns:        int(trt.maximumIdleConnections),
			MaxIdleConnsPerHost: int(trt.maximumIdleConnectionsPerHost),
			IdleConnTimeout:     trt.idleConnectionTimeout,
			TLSHandshakeTimeout: trt.tlsHandshakeTimeout,
			TLSClientConfig:     trt.tlsClientConfig,
			DialTLS:             trt.tlsDialFunc,
			DialContext:         trt.dialContextCustomFunc,
		}
		if trt.dialContextCustomFunc == nil {
			trt.transport.DialContext = (&net.Dialer{
				Timeout:   trt.dialContextTimeout,
				KeepAlive: trt.dialContextKeepAlive,
				DualStack: trt.dialContextDualStack,
			}).DialContext
		}
	}
	// Создание клиента http.
	trt.client = &http.Client{
		Transport: trt.transport,
		Timeout:   trt.totalTimeout,
		Jar:       trt.cookieJar,
	}

	return trt.client
}

// Do Выполнение запроса в асинхронном режиме.
func (trt *impl) Do(req request.Interface) Interface {
	// Создание транспорта, клиента http.
	if trt.client == nil {
		_ = trt.Client()
	}
	// Создание и запуск пула воркеров для обслуживания запросов.
	if !trt.requestPoolStarted.Load().(bool) {
		trt.makePool()
	}
	// Добавление запроса в пул задач.
	trt.requestChan <- req

	return trt
}

// Done Остановка процессов работников, завершение соединений.
func (trt *impl) Done() {
	defer func() {
		if e := recover(); e != nil {
			trt.err = fmt.Errorf("Catch panic: %s\nGoroutine stack is:\n%s", e.(error), string(runtimeDebug.Stack()))
			if trt.errFunc != nil {
				trt.errFunc(trt.err)
			}
			return
		}
	}()
	trt.requestPoolLock.Lock()
	defer trt.requestPoolLock.Unlock()
	// Выход, если пул воркеров остановлен
	if !trt.requestPoolStarted.Load().(bool) {
		return
	}
	// Завершаем все воркеры пула
	for i := range trt.requestPoolCancelFunc {
		if trt.requestPoolCancelFunc[i] != nil {
			trt.requestPoolCancelFunc[i]()
		}
	}
	//  Ожидание завершения воркеров
	trt.requestPoolDone.Wait()
	trt.requestPoolStarted.Store(false)
	close(trt.requestChan)
}
