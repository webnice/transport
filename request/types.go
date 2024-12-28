package request

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/webnice/dic"
	"github.com/webnice/transport/v4/header"
	"github.com/webnice/transport/v4/response"
)

// Pool Интерфейс бассейна повторно используемых объектов.
type Pool interface {
	// RequestGet Извлечение из бассейна нового объекта Request.
	RequestGet() Interface

	// RequestPut Возврат в бассейн элемента Request.
	RequestPut(req Interface)
}

// Interface Интерфейс пакета.
type Interface interface {
	// Cancel Прерывание запроса.
	Cancel() Interface

	// Done Ожидание завершения выполнения запроса.
	Done() Interface

	// DoneWithContext Ожидание завершения выполнения запроса с передачей контекста с возможностью
	// дополнительного контроля и прерывания.
	DoneWithContext(ctx context.Context) Interface

	// Error Последняя ошибка.
	Error() error

	// DebugFunc Включение или отключение режима отладки.
	// Если передана функция отладки не равная nil, режим отладки включается.
	// Передача функции отладки равной nil отключает режим отладки.
	DebugFunc(fn DebugFunc) Interface

	// ЗАПРОС

	// Request Возвращается http.Request подготовленный к выполнению запроса.
	Request() (*http.Request, error)

	// Method Назначение метода выполнения запроса.
	Method(dic.IMethod) Interface

	// URL Назначение URI адреса для выполнения запроса.
	// Deprecated: Используйте метод Uri.
	URL(url string) Interface

	// Uri Назначение URI адреса для выполнения запроса.
	Uri(uri string) Interface

	// Referer Назначение заголовка Referer.
	Referer(referer string) Interface

	// UserAgent Назначение заголовка UserAgent.
	UserAgent(userAgent string) Interface

	// ContentType Назначение заголовка Content-Type.
	ContentType(contentType string) Interface

	// Accept Назначение заголовка Accept.
	Accept(accept string) Interface

	// AcceptEncoding Назначение заголовка Accept-Encoding.
	AcceptEncoding(acceptEncoding string) Interface

	// AcceptLanguage Назначение заголовка Accept-Language.
	AcceptLanguage(acceptLanguage string) Interface

	// AcceptCharset Назначение заголовка Accept-Charset.
	AcceptCharset(acceptCharset string) Interface

	// CustomHeader Назначение заголовка с произвольным названием и значением.
	CustomHeader(name string, value string) Interface

	// BasicAuth Назначение пользователя и пароля простой web авторизации.
	BasicAuth(username string, password string) Interface

	// Cookies Добавление печенек в запрос.
	Cookies(cookies []*http.Cookie) Interface

	// Header Интерфейс заголовка запроса.
	Header() header.Interface

	// Latency Задержка ответа сервера на запрос.
	// Значение получено без учёта времени на чтение заголовков и тела ответа.
	Latency() time.Duration

	// Response Интерфейс ответа на запрос.
	Response() response.Interface

	// ДАННЫЕ

	// DataStream Потоковые данные для тела запроса.
	DataStream(data io.Reader) Interface

	// DataString Строковые данные тела запроса.
	DataString(data string) Interface

	// DataBytes Данные тела запроса, представленные в качестве среза байт.
	DataBytes(data []byte) Interface

	// DataJSON Данные тела запроса представленные в качестве объекта.
	// Объект перед передачей сериализуется в JSON.
	DataJSON(data any) Interface

	// DataXML Данные тела запроса представленные в качестве объекта.
	// Объект перед передачей сериализуется в XML.
	DataXML(data any) Interface

	// ВЫПОЛНЕНИЕ

	// Do Создание и выполнение запроса инициализация и возврат интерфейса работы с ответом.
	Do(client *http.Client) error
}

// Объект сущности пакета.
type impl struct {
	requestPool  *sync.Pool    // Бассейн объектов Request.
	responsePool response.Pool // Интерфейс бассейна объектов Response.
}

// DebugFunc Описание функции отладки запросов.
type DebugFunc func(data []byte)

// Request Объект Request.
type Request struct {
	context              context.Context    // Интерфейс контекста.
	contextCancelFunc    context.CancelFunc // Функция прерывания через контекст.
	method               dic.IMethod        // Метод запроса данных.
	header               header.Interface   // Заголовки запроса.
	err                  error              // Последняя ошибка.
	debugFunc            DebugFunc          // Функция отладки и мониторинга.
	uri                  *bytes.Buffer      // Запрашиваемый URI без данных.
	request              *http.Request      // Объект http.Request.
	requestData          *bytes.Reader      // Данные запроса.
	requestDataInterface io.Reader          // Интерфейс данных запроса.
	username             string             // Имя пользователя для авторизации, если указан, то передаются заголовки авторизации.
	password             string             // Пароль пользователя для авторизации.
	cookie               []*http.Cookie     // Печеньки запроса.
	timeBegin            time.Time          // Дата и время начала запроса.
	timeLatency          time.Duration      // Время ушедшее за выполнение запроса.
	response             response.Interface // Интерфейс результата запроса.
	// Переменные.
	tmpArr     []string // Общая переменная.
	tmpOk      bool     // Общая переменная.
	tmpCounter int      // Общая переменная.
	tmpBytes   []byte   // Общая переменная.
}
