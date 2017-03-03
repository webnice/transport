package transport

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"net/http"
	"time"

	"gopkg.in/webnice/transport.v1/charmap"
)

// Response Return net/http.Response object as is
func (rsp *responseImplementation) Response() *http.Response {
	return rsp.HTTPResponse
}

// Latency Скорость ответа сервера (без передачи ответа на запрос)
func (rsp *responseImplementation) Latency() time.Duration {
	return rsp.ResponseLatency
}

// LatencyData Скорость ответа с передачей данных
func (rsp *responseImplementation) LatencyData() time.Duration {
	return rsp.ResponseLatencyData
}

// Error Ошибка
func (rsp *responseImplementation) Error() error {
	return rsp.ResponseError
}

// ContentLength Размер в байтах загруженного ответа
func (rsp *responseImplementation) ContentLength() int64 {
	return rsp.ResponseContentLength
}

// StatusCode HTTP Код ответа на запрос
func (rsp *responseImplementation) StatusCode() int {
	return rsp.ResponseCode
}

// Status HTTP Код ответа на запрос в виде строки
func (rsp *responseImplementation) Status() string {
	return rsp.ResponseStatus
}

// Header Интерфейс к получению заголовков ответа
func (rsp *responseImplementation) Header() HeaderInterface {
	return &headerImplementation{rsp.HTTPResponse.Header}
}

// Cookies Получение всех переданных с сервера кукисов
func (rsp *responseImplementation) Cookies() (ret []*http.Cookie) {
	return rsp.HTTPResponse.Cookies()
}

// Content Интерфейс работы с ответом сервера
// Интерфейс работает только если не был передан io.Writer для ответа сервера
// Если io.Writer был передан, то все данные пишутся в него и интерфейсу не счем раотать
func (rsp *responseImplementation) Content() ContentInterface {
	return &contentImplementation{
		ResponseFHName: rsp.ResponseFHName,
		ResponseFH:     rsp.ResponseFH,
	}
}

// Charmap Charmap interface
func (rsp *responseImplementation) Charmap() charmap.Charmap {
	return charmap.NewCharmap()
}
