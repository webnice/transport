package transport

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/webnice/transport/v4/request"
)

// Interface is an interface of object
type Interface interface {
	// RequestPoolSize Назначение количества процессов в бассейне выполнения запросов.
	RequestPoolSize(n uint16) Interface

	// ProxyFunc Назначение функции настроек прокси для выполнения запросов.
	ProxyFunc(f ProxyFunc) Interface

	// ProxyConnectHeader Назначение заголовков запроса установки соединения с прокси сервером.
	ProxyConnectHeader(v http.Header) Interface

	// DialContextTimeout Назначение максимального вреени ожидания на загрузку контента. 0-не ограничено.
	DialContextTimeout(t time.Duration) Interface

	// DialContextKeepAlive Назначение времени поддержания не активного соединения до его разрыва. 0-без ограничений.
	DialContextKeepAlive(t time.Duration) Interface

	// MaximumIdleConnections Назначение максимального количества открытых не активных соединений. 0-без ограничений.
	MaximumIdleConnections(v uint) Interface

	// MaximumIdleConnectionsPerHost Назначение максимального количества открытых не активных соединений для
	// каждого хоста. 0-без ограничений.
	MaximumIdleConnectionsPerHost(v uint) Interface

	// IdleConnectionTimeout Назначение максимального количества открытых не активных соединений для всех
	// хостов. 0-без ограничений.
	IdleConnectionTimeout(t time.Duration) Interface

	// TLSHandshakeTimeout Назначение максимального времени ожидания обмена рукопожатиями по протоколу
	// TLS. 0-без ограничений.
	TLSHandshakeTimeout(t time.Duration) Interface

	// TLSSkipVerify Установка режима отключения проверки TLS сертификатов.
	TLSSkipVerify(v bool) Interface

	// TLSClientConfig Настройки клиента TLS соединения.
	// Если =nil-используются настройки по умолчанию из стандартной библиотеки.
	TLSClientConfig(v *tls.Config) Interface

	// DialTLS Назначение функции установки TLS соединения для запросов к HTTPS хостам.
	DialTLS(fn DialTLSFunc) Interface

	// DialContextCustomFunc Назначение функции установки не шифрованного соединения с хостами.
	DialContextCustomFunc(fn DialContextFunc) Interface

	// DualStack Включение или отключения функции "Happy Eyeballs" RFC 6555.
	DualStack(v bool) Interface

	// TotalTimeout Установка ограничения времени на выполнение запроса и загрузку всех данных
	// ответа. 0-без ограничений.
	TotalTimeout(t time.Duration) Interface

	// Transport Объект http транспорта.
	Transport(tr *http.Transport) Interface

	// CookieJar Интерфейс объекта печенек.
	CookieJar(v http.CookieJar) Interface

	// RequestGet Получение из бассейна объекта request.Interface.
	// Полученный объект обязательно необходимо вернуть в бассейн методом RequestPut для избежания утечки памяти.
	RequestGet() request.Interface

	// RequestPut Возвращение в бассейн объекта request.Interface.
	RequestPut(req request.Interface)

	// Client Получение клиента http.Client.
	// В пределах одного экземпляра transport.impl, http.Client создаётся только один раз
	// при первом вызове данной функции. Эта функция так же вызывается при первом вызове функции Do().
	Client() *http.Client

	// Do Выполнение запроса в асинхронном режиме.
	Do(req request.Interface) Interface

	// Done Остановка процессов работников, завершение соединений.
	Done()

	// DebugFunc Включение или отключение режима отладки.
	// Если передана функция отладки не равная nil, режим отладки включается.
	// Передача функции отладки равной nil отключает режим отладки.
	DebugFunc(fn DebugFunc) Interface

	// ОШИБКИ

	// Error Последняя ошибка.
	Error() error

	// ErrorFunc Регистрация функции получения ошибок транспорта.
	ErrorFunc(fn ErrorFunc) Interface
}
