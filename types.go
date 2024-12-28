package transport

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"sync"
	"sync/atomic"
	"time"

	"github.com/webnice/transport/v4/request"
)

const (
	defaultRequestPoolSize               = uint16(1)        // Количество процессов в бассейне выполнения запросов.
	defaultDialContextTimeout            = time.Duration(0) // Время ожидания ответа на запрос. По умолчанию - бесконечно.
	defaultDialContextKeepAlive          = 30 * time.Second // Время поддержания не активного соединения до его разрыва.
	defaultMaximumIdleConnections        = 100              // Максимальное количество открытых не активных соединений. 0-без ограничений.
	defaultMaximumIdleConnectionsPerHost = 10               // Максимальное количество открытых не активных соединений для каждого хоста. 0-без ограничений.
	defaultIdleConnectionTimeout         = 90 * time.Second // Максимальное количество открытых не активных соединений для всех хостов. 0-без ограничений.
	defaultTLSHandshakeTimeout           = 10 * time.Second // Максимальное время ожидания обменом рукопожатиями по протоколу TLS. 0-без ограничений.
	defaultTLSInsecureSkipVerify         = false            // Отключение проверки TLS сертификатов.
	defaultDialContextDualStack          = true             // Включение "Happy Eyeballs" RFC 6555.
	requestChanBuffer                    = int(1000)        // Размер буферизированного канала обмена данными.
)

// ProxyFunc Функция настроек прокси для выполнения запроса.
type ProxyFunc func(*http.Request) (*url.URL, error)

// ErrorFunc Функция получения ошибок запроса.
type ErrorFunc func(err error)

// DebugFunc Функция отладки и мониторинга запросов.
type DebugFunc func(data []byte)

// DialTLSFunc Функция выполнения соединения TLS для запросов к HTTPS хостам.
type DialTLSFunc func(network, addr string) (net.Conn, error)

// DialContextFunc Функция выполнения подключения к не шифрованным TCP/IP хостам.
type DialContextFunc func(ctx context.Context, network, addr string) (net.Conn, error)

// Объект сущности пакета.
type impl struct {
	client                        *http.Client           // Объект http клиента.
	transport                     *http.Transport        // Объект http транспорта.
	cookieJar                     http.CookieJar         // Интерфейс CookieJar.
	requestChan                   chan request.Interface // Канал задач запросов.
	requestPoolLock               *sync.Mutex            // Блокировка от двойного запуска.
	requestPoolCancelFunc         []context.CancelFunc   // Массив функций остановки бассейна работников.
	requestPoolStarted            *atomic.Value          // =true Бассейн работников запущен.
	requestPoolDone               *sync.WaitGroup        // WaitGroup для полной корректной остановки пула.
	err                           error                  // Последняя ошибка.
	errFunc                       ErrorFunc              // Функция получения ошибок на стороне http клиента.
	debugFunc                     DebugFunc              // Функция отладки и мониторинга запросов.
	requestPoolInterface          request.Pool           // Интерфейс бассейна объектов запросов.
	requestPoolSize               uint16                 // Количество процессов в бассейне выполнения запросов.
	proxy                         ProxyFunc              // Функция настроек прокси для выполнения запроса.
	proxyConnectHeader            http.Header            // Не обязательные заголовки для установки соединения с прокси сервером.
	dialContextTimeout            time.Duration          // Максимальное время ожидания на загрузку контента. 0-не ограничено.
	dialContextKeepAlive          time.Duration          // Время поддержания не активного соединения до его разрыва.
	maximumIdleConnections        uint                   // Максимальное количество открытых не активных соединений. 0-без ограничений.
	maximumIdleConnectionsPerHost uint                   // Максимальное количество открытых не активных соединений для каждого хоста. 0-без ограничений.
	idleConnectionTimeout         time.Duration          // Максимальное количество открытых не активных соединений для всех хостов. 0-без ограничений.
	tlsHandshakeTimeout           time.Duration          // Максимальное время ожидания обменом рукопожатиями по протоколу TLS. 0-без ограничений.
	tlsInsecureSkipVerify         bool                   // Отключение проверки TLS сертификатов.
	tlsClientConfig               *tls.Config            // Настройки клиента TLS соединения. Если =nil-используются настройки по умолчанию из стандартной библиотеки.
	tlsDialFunc                   DialTLSFunc            // Функция выполнения соединения TLS для запросов к HTTPS хостам.
	dialContextCustomFunc         DialContextFunc        // Функция выполнения подключения к не шифрованным TCP/IP хостам.
	dialContextDualStack          bool                   // Включение "Happy Eyeballs" RFC 6555.
	totalTimeout                  time.Duration          // Ограничение времени на выполнение запроса и загрузку всех данных ответа. 0-без ограничений.
}
