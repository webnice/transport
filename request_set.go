package transport

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/webnice/transport.v1/data"
	"gopkg.in/webnice/transport.v1/methods"
)

// Method Set request method
func (r *requestImplementation) Method(m methods.Value) Request {
	if m == nil {
		log.Printf("Warning, request method is nil, method not set")
		return r
	}
	r.RequestMethod = m
	return r
}

// Error Ошибка возникшая в ходе работы
func (r *requestImplementation) Error() (Request, error) {
	return r, r.RequestError
}

// URL Set request URL
func (r *requestImplementation) URL(url string) Request {
	r.RequestURL = url
	return r
}

// Referer Установка Referer запросу
func (r *requestImplementation) Referer(referer string) Request {
	r.RequestReferer = referer
	return r
}

// UserAgent Установка UserAgent запросу
func (r *requestImplementation) UserAgent(value string) Request {
	r.RequestUserAgent = value
	return r
}

// ContentType Установка Content-Type запросу
func (r *requestImplementation) ContentType(contentType string) Request {
	r.RequestContentType = contentType
	return r
}

// ProxyURL Установка ProxyURL запросу
func (r *requestImplementation) ProxyURL(proxyURL string) Request {
	if proxyURL != "" {
		r.RequestProxyURL, r.RequestError = new(url.URL).Parse(proxyURL)
	} else {
		r.RequestProxyURL = nil
	}
	return r
}

// TimeOut Установка максимального времени которое будет выполняться запрос.
// Если время таймаута истекает, то соединение разрывается, несмотря на то что данные в это время могут еще поступать или передаваться
// Если =0 - таймаут выключен. Если >0 - полное время на всю операцию, от подключения до полечения данных
func (r *requestImplementation) TimeOut(t time.Duration) Request {
	r.RequestTimeOut = t
	return r
}

// TLSVerifyOn При запросе адреса работающего на SSL/TLS включает проверку подписи сертификата SSL
func (r *requestImplementation) TLSVerifyOn() Request {
	r.RequestTLSSkipVerify = false
	return r
}

// TLSVerifyOff При запросе адреса работающего на SSL/TLS отключает проверку подписи сертификата SSL
func (r *requestImplementation) TLSVerifyOff() Request {
	r.RequestTLSSkipVerify = true
	return r
}

// Header Интерфейс работы с заголовками запроса
func (r *requestImplementation) Header() HeaderInterface {
	if r.RequestHeaders == nil {
		var hdr = make(http.Header)
		r.RequestHeaders = &headerImplementation{hdr}
	}
	return r.RequestHeaders
}

// Cookies Установка запросу кукисов с заменой, проверка по имени кука
func (r *requestImplementation) Cookies(cookies []*http.Cookie) Request {
	var newCookies []*http.Cookie
	var ok bool
	var i, n int

	// Все старые кроме новых
	for n = range r.RequestCookies {
		ok = false
		for i = range cookies {
			if cookies[i] == nil {
				continue
			}
			if strings.EqualFold(r.RequestCookies[n].Name, cookies[i].Name) {
				ok = true
			}
		}
		if !ok {
			newCookies = append(newCookies, r.RequestCookies[n])
		}
	}

	// Все новые с защитой от ошибки
	for i = range cookies {
		if cookies[i] == nil {
			log.Printf("Warning, cookie is nil, cookie object number %d skipped", i)
			continue
		}
		newCookies = append(newCookies, cookies[i])
	}

	return r
}

// Accept Установка Accept запросу
func (r *requestImplementation) Accept(value string) Request {
	r.RequestAccept = value
	return r
}

// AcceptEncoding Установка Accept-Encoding запросу
func (r *requestImplementation) AcceptEncoding(value string) Request {
	r.RequestAcceptEncoding = value
	return r
}

// AcceptLanguage Установка Accept-Language запросу
func (r *requestImplementation) AcceptLanguage(value string) Request {
	r.RequestAcceptLanguage = value
	return r
}

// Auth Установка логина/пароля для авторизации
func (r *requestImplementation) Auth(login string, password string) Request {
	r.AuthLogin = login
	r.AuthPassword = password
	return r
}

// DataString Данные для запроса в формате строки
func (r *requestImplementation) DataString(data string) Request {
	r.RequestDataString = data
	r.RequestData = bytes.NewReader([]byte(r.RequestDataString))
	return r
}

// DataBytes Данные для запроса в формате среза байт
func (r *requestImplementation) DataBytes(data []byte) Request {
	r.RequestDataBytes = append(r.RequestDataBytes, data...)
	r.RequestData = bytes.NewReader(r.RequestDataBytes)
	return r
}

// Data Данные для запроса в формате ридера
func (r *requestImplementation) Data(data *bytes.Reader) Request {
	r.RequestData = data
	return r
}

// DataJson Сериализация данных из объекта в JSON
func (r *requestImplementation) DataJson(obj interface{}) Request {
	var buf []byte
	buf, r.RequestError = json.Marshal(obj)
	if r.RequestError != nil {
		return r
	}
	r.DataBytes(buf)
	r.Header().Set(`Content-Type`, `application/json`)
	return r
}

// DataXml Сериализация данных из объекта в XML
func (r *requestImplementation) DataXml(obj interface{}) Request {
	var buf []byte
	buf, r.RequestError = xml.Marshal(obj)
	if r.RequestError != nil {
		return r
	}
	r.DataBytes(buf)
	r.Header().Set(`Content-Type`, `application/xml`)
	return r
}

// Response Установка writer в которй будут выгружены данные запроса
// В случае если метод не вызывался или высывался с аргументом nil,
// то для данных будет создан временный файл, который сам удалится при уничтожении объекта transport
func (r *requestImplementation) Response(w io.Writer) Request {
	r.ResponseData = data.NewWriteCloser(w)
	return r
}
