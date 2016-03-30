package transport // import "github.com/webdeskltd/transport"

//import "github.com/webdeskltd/debug"
import "github.com/webdeskltd/log"
import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/webdeskltd/transport/methods"
)

// Do Выполнение запроса, ожидание и получение результата
func (r *requestImplementation) Do() (ret Response, err error) {
	var client *http.Client
	var i int

	r.ResponseImplementation = new(responseImplementation)
	ret = r.ResponseImplementation

	// Создание запроса
	if err = r.MakeRequest(); err != nil {
		r.RequestError = err
		return
	}

	// Авторизация
	if r.AuthLogin != "" {
		r.HTTPRequest.SetBasicAuth(r.AuthLogin, r.AuthPassword)
	}

	r.Header().Set(`Referer`, r.RequestReferer)
	r.Header().Set(`User-Agent`, r.RequestUserAgent)
	r.Header().Set(`Content-Type`, r.RequestContentType)

	// Установка заголовков запросу
	r.MakeHeaders()

	// Создание клиента запроса
	client = r.MakeHTTPClient()

	// Кукисы
	for i = range r.RequestCookies {
		r.HTTPRequest.AddCookie(r.RequestCookies[i])
	}

	// Засекаем время запроса
	r.ResponseImplementation.ResponseBeginRequest = time.Now().In(time.Local)

	// Запрос
	if r.ResponseImplementation.HTTPResponse, err = client.Do(r.HTTPRequest); err != nil {
		r.RequestError = err
		r.ResponseImplementation.ResponseError = err
	}

	// Подсчитываем время ушедшее на запрос
	r.ResponseImplementation.ResponseLatency = time.Since(r.ResponseImplementation.ResponseBeginRequest)

	// Если была ошибка, то данных нет, выход
	if err != nil {
		return
	}

	// Закрыти входящего потока данных по завершении функции
	defer func() {
		if err = r.ResponseImplementation.HTTPResponse.Body.Close(); err != nil {
			log.Warning("Error closing the incoming data stream: %s", err.Error())
		}
	}()

	// Загрузка данных
	err = r.LoadData()

	return
}

// LoadData Load all data from reader and copy to writer
func (r *requestImplementation) LoadData() (err error) {
	// Получение writer для записи ответа
	if err = r.MakeOutgoingStream(); err != nil {
		err = fmt.Errorf("Error open outgoing data temporary file or stream: %s", err.Error())
		r.RequestError = err
		r.ResponseImplementation.ResponseError = err
		return
	}

	// Закрытие исходящего потока по завершении чтения данных
	defer func() {
		if r.ResponseData != nil {
			if err = r.ResponseData.Close(); err != nil {
				log.Warning("Error closing the outgoing data stream: %s", err.Error())
			}
		}
	}()

	// Загрузка данных
	r.ResponseImplementation.ResponseContentLength, err = io.Copy(r.ResponseData, r.ResponseImplementation.HTTPResponse.Body)
	if err != nil {
		err = fmt.Errorf("Error reading content: %v", err)
		r.RequestError = err
		r.ResponseImplementation.ResponseError = err
		return
	}

	// Подсчитываем время ушедшее на загрузку данных
	r.ResponseImplementation.ResponseLatencyData = time.Since(r.ResponseImplementation.ResponseBeginRequest)

	// Проверка заявленного размера данных и загруженного
	if r.ResponseImplementation.HTTPResponse.ContentLength != -1 &&
		r.ResponseImplementation.ResponseContentLength != r.ResponseImplementation.HTTPResponse.ContentLength {
		log.Warning("Content-length wrong or incomplite!")
	}

	// Результирующие HTTP коды ответа
	r.ResponseImplementation.ResponseStatus = r.ResponseImplementation.HTTPResponse.Status
	r.ResponseImplementation.ResponseCode = r.ResponseImplementation.HTTPResponse.StatusCode

	return
}

// MakeRequest Создание запроса на основе метода запроса
func (r *requestImplementation) MakeRequest() (err error) {
	var mtd = methods.New()
	var url *bytes.Buffer

	// Если можно исправить ошибку самостоятельно то надо это делать
	if r.RequestData == nil {
		r.RequestData = bytes.NewReader([]byte{})
	}

	// GET
	if r.RequestMethod.Type() == mtd.Get().Type() {
		url = bytes.NewBufferString(r.RequestURL)
		if r.RequestData.Len() > 0 {
			_, _ = url.WriteString(`?`)
			_, _ = r.RequestData.WriteTo(url)
		}
		r.HTTPRequest, err = http.NewRequest(r.RequestMethod.String(), url.String(), nil)
		return
	}

	// Все остальные
	r.HTTPRequest, err = http.NewRequest(r.RequestMethod.String(), r.RequestURL, r.RequestData)

	return
}

// MakeHeaders Установка всех заголовков запросу, пропуская пустые заголовки
func (r *requestImplementation) MakeHeaders() {
	if r.RequestHeaders.Header != nil {
		for kn := range r.RequestHeaders.Header {
			var kv = r.RequestHeaders.Header.Get(kn)
			if kv == "" {
				continue
			}
			r.HTTPRequest.Header.Set(kn, kv)
		}
	}
}

// MakeHTTPClient Создание HTTP Client
func (r *requestImplementation) MakeHTTPClient() (client *http.Client) {
	var bit int
	var tlsConfig *tls.Config

	// TLS configuration
	tlsConfig = &tls.Config{
		// PreferServerCipherSuites controls whether the server selects the
		// client's most preferred ciphersuite, or the server's most preferred
		// ciphersuite. If true then the server's preference, as expressed in
		// the order of elements in CipherSuites, is used.
		PreferServerCipherSuites: true,

		// InsecureSkipVerify controls whether a client verifies the
		// server's certificate chain and host name.
		// If InsecureSkipVerify is true, TLS accepts any certificate
		// presented by the server and any host name in that certificate.
		// In this mode, TLS is susceptible to man-in-the-middle attacks.
		// This should be used only for testing.
		InsecureSkipVerify: r.RequestTLSSkipVerify,
	}

	if r.RequestTimeOut > 0 {
		bit += 1 << 1
	}
	if r.RequestProxyURL != nil {
		bit++
	}

	switch bit {
	case 0:
		// Таймаут = нет, Прокси = нет
		client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig:   tlsConfig,
				DisableKeepAlives: true,
			},
		}
	case 1:
		// Таймаут = нет, Прокси = есть
		client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig:   tlsConfig,
				DisableKeepAlives: true,
				Proxy:             http.ProxyURL(r.RequestProxyURL),
			},
		}
	case 2:
		// Таймаут = есть, Прокси = нет
		client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig:   tlsConfig,
				DisableKeepAlives: true,
				Dial:              r.MakeFnTimeoutDialler(r.RequestTimeOut),
			},
		}
	case 3:
		// Таймаут = есть, Прокси = есть
		client = &http.Client{
			Transport: &http.Transport{
				// TLSClientConfig specifies the TLS configuration to use with
				// tls.Client. If nil, the default configuration is used.
				TLSClientConfig: tlsConfig,

				// DisableKeepAlives, if true, prevents re-use of TCP connections
				// between different HTTP requests.
				DisableKeepAlives: true,

				// Dial specifies the dial function for creating unencrypted
				// TCP connections.
				// If Dial is nil, net.Dial is used.
				Dial: r.MakeFnTimeoutDialler(r.RequestTimeOut),

				// Proxy specifies a function to return a proxy for a given
				// Request. If the function returns a non-nil error, the
				// request is aborted with the provided error.
				// If Proxy is nil or returns a nil *URL, no proxy is used.
				Proxy: http.ProxyURL(r.RequestProxyURL),
			},
		}
	}

	return
}

// MakeFnTimeoutDialler Создание функции для контроля таймаута
func (r *requestImplementation) MakeFnTimeoutDialler(timeout time.Duration) func(net, addr string) (client net.Conn, err error) {
	return func(netw, addr string) (client net.Conn, err error) {
		client, err = net.DialTimeout(netw, addr, time.Duration(timeout))
		if err != nil {
			log.Warning("Request timeout exceeded, drop connection: %v", err)
			return
		}
		err = client.SetDeadline(time.Now().Add(timeout))
		return
	}
}

// CreateFileName Создание пути к месту хранения файла и полного имени файла
func (r *requestImplementation) MakeTemporaryFileName() (ret string) {
	var tm = time.Now().In(time.Local)
	ret = path.Join(
		os.TempDir(),
		fmt.Sprintf("%020d.tmp", tm.UnixNano()),
	)
	return
}

// MakeOutgoingStream Создание временного файла или выбор потока для загрузки данных
func (r *requestImplementation) MakeOutgoingStream() (err error) {
	// Если был указан WriteClosser, то ничего не меняем
	if r.ResponseData != nil {
		return
	}

	// Создаём временный файл для записи данных

	// Создание имени файла
	r.ResponseImplementation.ResponseFHName = r.MakeTemporaryFileName()

	// Открытие файла на запись
	r.ResponseImplementation.ResponseFH, err = os.OpenFile(r.ResponseImplementation.ResponseFHName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return
	}

	// Добавление в коллекцию автоудаления фалов
	r.ResponseImplementation.ResponseFHEnable = true
	r.collectionOfTemporaryFilesFn(r.ResponseImplementation.ResponseFHName)

	// io.WriteCloser
	r.ResponseData = r.ResponseImplementation.ResponseFH

	return
}
