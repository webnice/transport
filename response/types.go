package response

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/webnice/transport/v4/charmap"
	"github.com/webnice/transport/v4/content"
	"github.com/webnice/transport/v4/data"
	"github.com/webnice/transport/v4/header"
)

const (
	// Максимальный размер данных загружаемый в память 250Mb.
	maxDataSizeLoadedInMemory = uint64(250 * 1024 * 1024)
)

// Pool Бассейн переиспользования объектов.
type Pool interface {
	// ResponseGet Извлечение из бассейна нового элемента Response.
	ResponseGet() Interface

	// ResponsePut Возврат в бассейн использованного элемента Response.
	ResponsePut(req Interface)
}

// Interface Интерфейс пакета.
type Interface interface {
	// DebugFunc Включение или отключение режима отладки.
	// Если передана функция отладки не равная nil, режим отладки включается.
	// Передача функции отладки равной nil отключает режим отладки.
	DebugFunc(fn DebugFunc) Interface

	// Do Выполнение запроса и получение Response.
	Do(client *http.Client, request *http.Request) error

	// Load Загрузка данных ответа на запрос.
	Load() error

	// Error Последняя ошибка.
	Error() error

	// Response Возвращает http.Response как есть.
	Response() *http.Response

	// ContentLength Длинна контента ответа на запрос.
	ContentLength() int64

	// Cookies Разбор заголовка с печеньками и возврат списка переданных печенек.
	Cookies() []*http.Cookie

	// Latency Задержка ответа сервера на запрос.
	// Значение получено без учёта времени на чтение заголовков и тела ответа.
	Latency() time.Duration

	// StatusCode Код http ответа.
	StatusCode() int

	// Status Строковый статус http ответа.
	Status() string

	// Header Интерфейс работы с заголовками, представлен в виде карты ключ=значение.
	Header() header.Interface

	// Charmap Интерфейс работы с кодировкой.
	Charmap() charmap.Charmap

	// RetryAfter Значение заголовка RetryAfter.
	// Если возвращается значение 0 - заголовок отсутствовал.
	RetryAfter() (ret time.Duration)

	// Content Интерфейс работы с контентом.
	Content() content.Interface
}

// Объект сущности пакета.
type impl struct {
	responsePool *sync.Pool // Бассейн объектов Response.
}

// DebugFunc Описание функции отладки запросов.
type DebugFunc func(data []byte)

// Response Объект Response.
type Response struct {
	err                   error                          // Последняя ошибка.
	response              *http.Response                 // Объект http.Response.
	debugFunc             DebugFunc                      // Функция отладки и мониторинга.
	timeBegin             time.Time                      // Дата и время начала загрузки результата запроса.
	timeLatency           time.Duration                  // Время ушедшее за выполнение загрузки результата запроса.
	contentInMemory       bool                           // =true - Результат в памяти, =false - результат во временном файле.
	contentData           *bytes.Buffer                  // Результат запроса в памяти.
	contentFilename       string                         // Название временного файла результата запроса.
	contentFh             *os.File                       // Интерфейс файлового дескриптора временного файла.
	contentTemporaryFiles []string                       // Названия временных файлов.
	contentWriteCloser    io.WriteCloser                 // Интерфейс io.WriteCloser для доступа к результату запроса в памяти.
	contentLength         int64                          // Размер загруженных данных.
	contentReader         data.ReadAtSeekerWriteToCloser // Интерфейс к данным результата запроса.
	charmap               charmap.Charmap                // Интерфейс работы с кодировкой.
	// Переменные.
	tmpOk     bool      // Общая переменная.
	tmpTm     time.Time // Общая переменная.
	tmpString string    // Общая переменная.
	tmpI      int       // Общая переменная.
}
