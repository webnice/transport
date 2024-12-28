package response

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"path"
	runtimeDebug "runtime/debug"
	"strconv"
	"time"

	"github.com/webnice/dic"
	"github.com/webnice/transport/v4/charmap"
	"github.com/webnice/transport/v4/content"
	"github.com/webnice/transport/v4/data"
	"github.com/webnice/transport/v4/header"
)

// DebugFunc Включение или отключение режима отладки.
// Если передана функция отладки не равная nil, режим отладки включается.
// Передача функции отладки равной nil отключает режим отладки.
func (r *Response) DebugFunc(fn DebugFunc) Interface { r.debugFunc = fn; return r }

// Do Выполнение запроса и получение Response.
func (r *Response) Do(client *http.Client, request *http.Request) (err error) {
	r.response, err = client.Do(request)

	return
}

// Создание пути к месту хранения файла и полного имени временного файла.
func (r *Response) makeTemporaryFileName() {
	r.tmpTm = time.Now().In(time.Local)
	r.contentFilename = path.Join(
		os.TempDir(),
		fmt.Sprintf("%020d.tmp", r.tmpTm.UnixNano()),
	)

	return
}

// Создание WriteCloser для загрузки результата запроса.
func (r *Response) makeContentContainer(size int64) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("Паника: %s\nСтек выполнения:\n%s", e.(error), string(runtimeDebug.Stack()))
			return
		}
	}()
	switch size < 0 || size > int64(maxDataSizeLoadedInMemory) {
	case true:
		// Создание временного файла для данных.
		r.makeTemporaryFileName()
		r.contentFh, err = os.OpenFile(r.contentFilename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.FileMode(0600))
		r.contentTemporaryFiles = append(r.contentTemporaryFiles, r.contentFilename)
		r.contentWriteCloser = r.contentFh
	default:
		// Чтение в память.
		r.contentData.Grow(int(r.response.ContentLength)) // Grow может паниковать (дебильный код).
		r.contentInMemory = true
		r.contentWriteCloser = data.NewWriteCloser(r.contentData)
	}

	return
}

// Создание единого интерфейса чтения загруженного контента.
func (r *Response) makeContentReader() {
	const defaultMode = os.FileMode(0600)

	switch r.contentInMemory {
	case true:
		r.contentReader = data.NewReadAtSeekerWriteToCloser(r.contentData)
	default:
		if r.contentFh, r.err = os.OpenFile(r.contentFilename, os.O_RDONLY, defaultMode); r.err != nil {
			return
		}
		r.contentReader = data.NewReadAtSeekerWriteToCloser(r.contentFh)
	}
}

// Load Загрузка данных ответа на запрос.
func (r *Response) Load() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("Паника: %s\nСтек выполнения:\n%s", e.(error), string(runtimeDebug.Stack()))
			return
		}
	}()
	// Создание интерфейса io.WriteCloser к результату запроса в памяти.
	if err = r.makeContentContainer(r.response.ContentLength); err != nil {
		return
	}
	// Засекаем время выполнения загрузки.
	r.timeBegin = time.Now().In(time.Local)
	// Загрузка данных.
	r.contentLength, err = io.Copy(r.contentWriteCloser, r.response.Body)
	// Подсчитываем время ушедшее на загрузку.
	r.timeLatency = time.Since(r.timeBegin)
	_ = r.contentWriteCloser.Close()
	_ = r.response.Body.Close()
	if err != nil {
		return
	}
	r.makeContentReader()
	// Если включён дебаг.
	if r.debugFunc != nil {
		r.response.Body = r.contentReader
		if buf, err := httputil.DumpResponse(r.response, true); err == nil {
			r.debugResponse(buf)
		}
		_, _ = r.contentReader.Seek(0, io.SeekStart)
	}

	return
}

// Error Последняя ошибка.
func (r *Response) Error() error { return r.err }

// Response Возвращает http.Response как есть.
func (r *Response) Response() (ret *http.Response) {
	_, _ = r.contentReader.Seek(0, io.SeekStart)
	r.response.Body, ret = r.contentReader, r.response

	return
}

// ContentLength Длинна контента ответа на запрос.
func (r *Response) ContentLength() int64 { return r.response.ContentLength }

// Cookies Разбор заголовка с печеньками и возврат списка переданных печенек.
func (r *Response) Cookies() []*http.Cookie { return r.response.Cookies() }

// Latency Задержка ответа сервера на запрос.
// Значение получено без учёта времени на чтение заголовков и тела ответа.
func (r *Response) Latency() time.Duration { return r.timeLatency }

// StatusCode Код http ответа.
func (r *Response) StatusCode() int { return r.response.StatusCode }

// Status Строковый статус http ответа.
func (r *Response) Status() string { return r.response.Status }

// Header Интерфейс работы с заголовками, представлен в виде карты ключ=значение.
func (r *Response) Header() header.Interface { return header.New(r.response.Header) }

// Charmap Интерфейс работы с кодировкой.
func (r *Response) Charmap() charmap.Charmap { return r.charmap }

// RetryAfter Значение заголовка RetryAfter.
// Если возвращается значение 0 - заголовок отсутствовал или не соответствовал стандартам.
func (r *Response) RetryAfter() (ret time.Duration) {
	const errConv = "конвертация значения заголовка RetryAfter из %q в число прервана ошибкой: %s"
	var (
		err error
		raw string
		num int64
	)

	if !r.Header().IsSet(dic.Header().RetryAfter.String()) {
		return
	}
	if raw = r.Header().Get(dic.Header().RetryAfter.String()); raw == "" {
		return
	}
	if num, err = strconv.ParseInt(raw, 10, 64); err != nil {
		log.Printf(errConv, raw, err)
		return
	}
	ret = time.Second * time.Duration(num)

	return
}

// Content Интерфейс работы с контентом.
func (r *Response) Content() content.Interface { return content.New(r.contentReader) }
