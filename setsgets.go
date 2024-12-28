package transport

import (
	"crypto/tls"
	"net/http"
	"time"
)

func setDefaults(trt *impl) {
	trt.requestPoolSize, trt.dialContextTimeout, trt.dialContextKeepAlive =
		defaultRequestPoolSize, defaultDialContextTimeout, defaultDialContextKeepAlive
	trt.maximumIdleConnections, trt.maximumIdleConnectionsPerHost, trt.idleConnectionTimeout, trt.tlsHandshakeTimeout, trt.tlsInsecureSkipVerify, trt.dialContextDualStack =
		defaultMaximumIdleConnections, defaultMaximumIdleConnectionsPerHost, defaultIdleConnectionTimeout, defaultTLSHandshakeTimeout, defaultTLSInsecureSkipVerify, defaultDialContextDualStack
	trt.ProxyFunc(http.ProxyFromEnvironment)
}

// RequestPoolSize Назначение количества процессов в бассейне выполнения запросов.
func (trt *impl) RequestPoolSize(v uint16) Interface {
	if v == 0 {
		return trt
	}
	trt.requestPoolSize = v
	return trt
}

// ProxyFunc Назначение функции настроек прокси для выполнения запросов.
func (trt *impl) ProxyFunc(f ProxyFunc) Interface {
	if f == nil {
		return trt
	}
	trt.proxy = f
	return trt
}

// ProxyConnectHeader Назначение заголовков запроса установки соединения с прокси сервером.
func (trt *impl) ProxyConnectHeader(v http.Header) Interface {
	trt.proxyConnectHeader = v
	return trt
}

// DialContextTimeout Назначение максимального времени ожидания на загрузку контента. 0-не ограничено.
func (trt *impl) DialContextTimeout(t time.Duration) Interface {
	trt.dialContextTimeout = t
	return trt
}

// DialContextKeepAlive Назначение времени поддержания не активного соединения до его разрыва. 0-без ограничений.
func (trt *impl) DialContextKeepAlive(t time.Duration) Interface {
	trt.dialContextKeepAlive = t
	return trt
}

// MaximumIdleConnections Назначение максимального количества открытых не активных соединений. 0-без ограничений.
func (trt *impl) MaximumIdleConnections(v uint) Interface {
	trt.maximumIdleConnections = v
	return trt
}

// MaximumIdleConnectionsPerHost Назначение максимального количества открытых не активных соединений для
// каждого хоста. 0-без ограничений.
func (trt *impl) MaximumIdleConnectionsPerHost(v uint) Interface {
	trt.maximumIdleConnectionsPerHost = v
	return trt
}

// IdleConnectionTimeout Назначение максимального количества открытых не активных соединений для всех
// хостов. 0-без ограничений.
func (trt *impl) IdleConnectionTimeout(t time.Duration) Interface {
	trt.idleConnectionTimeout = t
	return trt
}

// TLSHandshakeTimeout Назначение максимального времени ожидания обмена рукопожатиями по протоколу
// TLS. 0-без ограничений.
func (trt *impl) TLSHandshakeTimeout(t time.Duration) Interface {
	trt.tlsHandshakeTimeout = t
	return trt
}

// TLSSkipVerify Установка режима отключения проверки TLS сертификатов.
func (trt *impl) TLSSkipVerify(v bool) Interface {
	trt.tlsInsecureSkipVerify = v
	return trt
}

// TLSClientConfig Настройки клиента TLS соединения.
// Если =nil-используются настройки по умолчанию из стандартной библиотеки.
func (trt *impl) TLSClientConfig(v *tls.Config) Interface {
	trt.tlsClientConfig = v
	return trt
}

// DialTLS Назначение функции установки TLS соединения для запросов к HTTPS хостам.
func (trt *impl) DialTLS(fn DialTLSFunc) Interface {
	trt.tlsDialFunc = fn
	return trt
}

// DialContextCustomFunc Назначение функции установки не шифрованного соединения с хостами.
func (trt *impl) DialContextCustomFunc(fn DialContextFunc) Interface {
	trt.dialContextCustomFunc = fn
	return trt
}

// DualStack Включение или отключения функции "Happy Eyeballs" RFC 6555.
func (trt *impl) DualStack(v bool) Interface {
	trt.dialContextDualStack = v
	return trt
}

// TotalTimeout Установка ограничения времени на выполнение запроса и загрузку всех данных
// ответа. 0-без ограничений.
func (trt *impl) TotalTimeout(t time.Duration) Interface {
	trt.totalTimeout = t
	return trt
}

// Transport Объект http транспорта.
func (trt *impl) Transport(tr *http.Transport) Interface {
	trt.transport = tr
	return trt
}

// CookieJar Интерфейс объекта печенек.
func (trt *impl) CookieJar(v http.CookieJar) Interface {
	trt.cookieJar = v
	return trt
}
