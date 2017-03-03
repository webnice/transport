package transport

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/webdeskltd/transport/charmap"
	"github.com/webdeskltd/transport/methods"

	"golang.org/x/text/encoding"
)

// Transport is an interface
type Transport interface {
	Method() methods.Interface
	NewRequest(method methods.Value) Request
}

// Request is an request interface
type Request interface {
	Accept(string) Request
	AcceptEncoding(string) Request
	AcceptLanguage(string) Request
	Auth(string, string) Request
	ContentType(string) Request
	Cookies([]*http.Cookie) Request
	DataString(string) Request
	DataBytes([]byte) Request
	Data(*bytes.Reader) Request
	Method(methods.Value) Request
	ProxyURL(string) Request
	Referer(string) Request
	TimeOut(time.Duration) Request
	UserAgent(string) Request
	URL(string) Request
	Response(io.Writer) Request
	TLSVerifyOn() Request
	TLSVerifyOff() Request

	ClientSource() (*http.Client, error)
	RequestSource() (*http.Request, error)
	Do() (Response, error)
	Error() (Request, error)
	Header() HeaderInterface
}

// Response is an response result interface
type Response interface {
	Content() ContentInterface
	ContentLength() int64
	Cookies() []*http.Cookie
	Error() error
	Header() HeaderInterface
	Latency() time.Duration
	LatencyData() time.Duration
	Response() *http.Response
	StatusCode() int
	Status() string
	Charmap() charmap.Charmap
}

// HeaderInterface is an requestHeadersImplementation interface
type HeaderInterface interface {
	Add(string, string)
	Del(string)
	Get(string) string
	Set(string, string)
	Names() []string
}

// ContentInterface is an contentImplementation interface
type ContentInterface interface {
	Transcode(encoding.Encoding) ContentInterface
	Transform(TransformFunc) ContentInterface
	String() (string, error)
	Bytes() ([]byte, error)
	ReaderCloser() (io.ReadCloser, error)
	Write(io.Writer) error
	ContentUnmarshalJSON(interface{}) error
	ContentUnmarshalXML(interface{}) error
	Untar() ContentInterface
	Unzip() ContentInterface
	UnGzip() ContentInterface
}

// TransformFunc is an func for streaming content conversion
type TransformFunc func(io.Reader) (io.Reader, error)

// implementation is an implementation
type implementation struct {
	CollectionOfTemporaryFiles []string
}

// requestImplementation is an Request implementation
type requestImplementation struct {
	RequestMethod                methods.Value           // Метод запроса данных
	RequestURL                   string                  // Запрашиваемый URL без данных
	RequestReferer               string                  // Referer
	RequestUserAgent             string                  // User-Agent
	RequestContentType           string                  // Content-Type
	RequestAccept                string                  // Accept
	RequestAcceptEncoding        string                  // Accept-Encoding
	RequestAcceptLanguage        string                  // Accept-Language
	RequestProxyURL              *url.URL                // URL прокси сервера
	RequestTimeOut               time.Duration           // Таймаут получения данных. Если =0 - выключен. Если >0 - полное время на всю операцию, от подключения до полечения данных
	RequestCookies               []*http.Cookie          // Куки запроса
	AuthLogin                    string                  // Логин для авторизации, если не пустой то производится авторизация
	AuthPassword                 string                  // Пароль для авторизации
	RequestDataString            string                  // Данные для запроса в формате строки
	RequestDataBytes             []byte                  // Данные для запроса в формате среза байт
	RequestData                  *bytes.Reader           // Отправляемые данные
	RequestHeaders               *headerImplementation   // Заголовки запроса
	RequestError                 error                   // Ошибка
	RequestTLSSkipVerify         bool                    // Проверка подписи сертификата SSL/TLS. =true - проверка отключена, =false - проверка включена (по умолчанию)
	ResponseImplementation       *responseImplementation // Объект результата запроса
	HTTPRequest                  *http.Request           // Объект net/http запроса
	collectionOfTemporaryFilesFn func(string)            // Функция коллекционирует временные файлы, которые необходимо удалить при уничтожении объекта
	ResponseData                 io.WriteCloser          // Исходящий поток для загружаемых данных, если указан, то временный файл не создаётся а данные пишутся во writer
}

// headerImplementation is an RequestHeaders and ResponseHeaders implementation
type headerImplementation struct {
	Header http.Header
}

// responseImplementation is an Response implementation
type responseImplementation struct {
	HTTPResponse          *http.Response // Объект htp/http ответа на запрос
	ResponseBeginRequest  time.Time      // Дата и время начала запроса
	ResponseLatency       time.Duration  // Скорость ответа сервера (без передачи ответа на запрос)
	ResponseLatencyData   time.Duration  // Скорость ответа с передачей данных
	ResponseError         error          // Ошибка
	ResponseContentLength int64          // Размер загруженного контента
	ResponseCode          int            // HTTP код ответа
	ResponseStatus        string         // HTTP код ответа в строковом виде
	ResponseFHEnable      bool           // =true - открыт временный файл, =false - данные пишутся в исходящий поток, временного файла нет
	ResponseFHName        string         // Имя временного файла
	ResponseFH            *os.File       // Временный файл для получаемых данных
}

// contentImplementation is an Content implementation
type contentImplementation struct {
	ResponseFHName string            // Имя временного файла
	ResponseFH     *os.File          // Временный файл для получаемых данных
	transcode      encoding.Encoding // Если не равно nil, то контент перекодируется на лету из указанной кодировки
	transform      TransformFunc     // Функция потокового преобразования контента
	unzip          bool              // =true - контент разархивируется методом ZIP, возвращается первый файл в архиве
	untar          bool              // =true - контент разархивируется методом TAR, возвращается первый файл в архиве
	ungzip         bool              // =true - контент разархивируется методом GZIP, возвращается первый файл в архиве
}
