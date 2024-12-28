package request

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/webnice/dic"
	transportHeader "github.com/webnice/transport/v4/header"
	"github.com/webnice/transport/v4/response"
)

// Cancel Прерывание запроса.
func (r *Request) Cancel() Interface { r.contextCancelFunc(); return r }

// Done Ожидание завершения выполнения запроса.
func (r *Request) Done() Interface { <-r.context.Done(); return r }

// DoneWithContext Ожидание завершения выполнения запроса с передачей контекста с возможностью
// дополнительного контроля и прерывания.
func (r *Request) DoneWithContext(ctx context.Context) Interface {
	select {
	case <-r.context.Done():
	case <-ctx.Done():
		r.Cancel()
	}

	return r
}

// Error Последняя ошибка.
func (r *Request) Error() error { return r.err }

// DebugFunc Включение или отключение режима отладки.
// Если передана функция отладки не равная nil, режим отладки включается.
// Передача функции отладки равной nil отключает режим отладки.
func (r *Request) DebugFunc(fn DebugFunc) Interface {
	r.debugFunc = fn
	r.response.DebugFunc(response.DebugFunc(fn))
	return r
}

// Method Назначение метода выполнения запроса.
func (r *Request) Method(m dic.IMethod) Interface {
	const errMethod = "метод запроса не установлен или передано nil значение"

	if m == nil {
		r.err = errors.New(errMethod)
		return r
	}
	r.method = m

	return r
}

// URL Назначение URI адреса для выполнения запроса.
// Deprecated: Используйте метод Uri.
func (r *Request) URL(uri string) Interface { return r.Uri(uri) }

// Uri Назначение URI адреса для выполнения запроса.
func (r *Request) Uri(uri string) Interface {
	r.uri.Reset()
	r.uri.WriteString(uri)
	return r
}

// Referer Назначение заголовка Referer.
func (r *Request) Referer(referer string) Interface {
	r.header.Add(dic.Header().Referer.String(), referer)
	return r
}

// UserAgent Назначение заголовка UserAgent.
func (r *Request) UserAgent(userAgent string) Interface {
	r.header.Add(dic.Header().UserAgent.String(), userAgent)
	return r
}

// ContentType Назначение заголовка Content-Type.
func (r *Request) ContentType(contentType string) Interface {
	r.header.Add(dic.Header().ContentType.String(), contentType)
	return r
}

// Accept Назначение заголовка Accept.
func (r *Request) Accept(accept string) Interface {
	r.header.Add(dic.Header().Accept.String(), accept)
	return r
}

// AcceptEncoding Назначение заголовка Accept-Encoding.
func (r *Request) AcceptEncoding(acceptEncoding string) Interface {
	r.header.Add(dic.Header().AcceptEncoding.String(), acceptEncoding)
	return r
}

// AcceptLanguage Назначение заголовка Accept-Language.
func (r *Request) AcceptLanguage(acceptLanguage string) Interface {
	r.header.Add(dic.Header().AcceptLanguage.String(), acceptLanguage)
	return r
}

// AcceptCharset Назначение заголовка Accept-Charset.
func (r *Request) AcceptCharset(acceptCharset string) Interface {
	r.header.Add(dic.Header().AcceptCharset.String(), acceptCharset)
	return r
}

// CustomHeader Назначение заголовка с произвольным названием и значением.
func (r *Request) CustomHeader(name string, value string) Interface {
	r.header.Add(name, value)
	return r
}

// BasicAuth Назначение пользователя и пароля простой web авторизации.
func (r *Request) BasicAuth(username string, password string) Interface {
	r.username, r.password = username, password
	return r
}

// Cookies Добавление печенек в запрос.
func (r *Request) Cookies(cookies []*http.Cookie) Interface {
	r.cookie = append(r.cookie, cookies...)
	return r
}

// Header Интерфейс заголовка запроса.
func (r *Request) Header() transportHeader.Interface { return r.header }

// Latency Задержка ответа сервера на запрос.
// Значение получено без учёта времени на чтение заголовков и тела ответа.
func (r *Request) Latency() time.Duration { return r.timeLatency }

// Response Интерфейс ответа на запрос.
func (r *Request) Response() response.Interface { return r.response }

// DataStream Потоковые данные для тела запроса.
func (r *Request) DataStream(data io.Reader) Interface {
	r.requestData, r.requestDataInterface = &bytes.Reader{}, data
	return r
}

// DataString Строковые данные тела запроса.
func (r *Request) DataString(data string) Interface {
	r.requestDataInterface, r.requestData = nil, bytes.NewReader([]byte(data))
	return r
}

// DataBytes Данные тела запроса, представленные в качестве среза байт.
func (r *Request) DataBytes(data []byte) Interface {
	r.requestDataInterface, r.requestData = nil, bytes.NewReader(data)
	return r
}

// DataJSON Данные тела запроса представленные в качестве объекта.
// Объект перед передачей сериализуется в JSON.
func (r *Request) DataJSON(data any) Interface {
	if r.tmpBytes, r.err = json.Marshal(data); r.err != nil {
		r.contextCancelFunc()
	}
	r.DataBytes(r.tmpBytes)

	return r
}

// DataXML Данные тела запроса представленные в качестве объекта.
// Объект перед передачей сериализуется в XML.
func (r *Request) DataXML(data any) Interface {
	if r.tmpBytes, r.err = xml.Marshal(data); r.err != nil {
		r.contextCancelFunc()
	}
	r.DataBytes(r.tmpBytes)

	return r
}

// MakeRequest Создание запроса.
func (r *Request) MakeRequest() (err error) {
	// Данные передаём через интерфейс:
	// - если интерфейс =nil;
	// - если есть данные;
	if r.requestDataInterface == nil && r.requestData.Len() > 0 {
		r.requestDataInterface = r.requestData
	}
	// Для метода GET, перенос данных в параметры URN.
	// Если в URL нет "?" И есть данные.
	if r.method.IsEqual(dic.Method().Get) && bytes.Index(r.uri.Bytes(), []byte(`?`)) < 0 && r.requestData.Len() > 0 {
		if _, err = r.uri.WriteString(`?`); err != nil {
			return
		}
		if _, err = r.requestData.WriteTo(r.uri); err != nil {
			return
		}
		r.requestDataInterface = nil
	}
	r.request, err = http.NewRequestWithContext(r.context, r.method.String(), r.uri.String(), r.requestDataInterface)

	return
}

// Request Возвращается http.Request подготовленный к выполнению запроса.
func (r *Request) Request() (ret *http.Request, err error) {
	err = r.MakeRequest()
	ret = r.request
	return
}

// Do Создание и выполнение запроса инициализация и возврат интерфейса работы с ответом.
func (r *Request) Do(client *http.Client) error {
	const errRequest = "запрос завершился вернув nil в качестве ответа"

	defer r.contextCancelFunc()
	// Создание запроса.
	if r.err = r.MakeRequest(); r.err != nil {
		return r.err
	}
	// Заголовки простой авторизации.
	if r.username != "" {
		r.request.SetBasicAuth(r.username, r.password)
	}
	// Печеньки запроса.
	if len(r.cookie) > 0 {
		for r.tmpCounter = range r.cookie {
			r.request.AddCookie(r.cookie[r.tmpCounter])
		}
	}
	// Заголовки.
	if r.header.Len() > 0 {
		r.tmpArr = r.header.Names()
		for r.tmpCounter = range r.tmpArr {
			if _, r.tmpOk = r.request.Header[r.tmpArr[r.tmpCounter]]; r.tmpOk {
				r.request.Header.Set(r.tmpArr[r.tmpCounter], r.header.Get(r.tmpArr[r.tmpCounter]))
			} else {
				r.request.Header.Add(r.tmpArr[r.tmpCounter], r.header.Get(r.tmpArr[r.tmpCounter]))
			}
		}
	}
	// Засекаем время запроса.
	r.timeBegin = time.Now().In(time.Local)
	// Выполнение запроса.
	r.err = r.response.Do(client, r.request)
	// Подсчитываем время ушедшее на запрос.
	r.timeLatency = time.Since(r.timeBegin)
	if r.err != nil {
		return r.err
	}
	// Отладка и мониторинг запроса.
	if r.debugFunc != nil {
		if r.requestData.Size() > 0 {
			_, _ = r.requestData.Seek(0, io.SeekStart)
		}
		if buf, err := httputil.DumpRequestOut(r.request, true); err == nil {
			buf = bytes.Join([][]byte{[]byte("URI: " + r.uri.String() + "\r\n"), buf}, []byte(``))
			r.debugRequest(buf)
		}
	}
	if r.response == nil {
		return errors.New(errRequest)
	}
	// Загрузка всех входящих данных.
	if r.err = r.response.Load(); r.err != nil {
		return r.err
	}

	return r.err
}
