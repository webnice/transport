package header

import "net/http"

// New Конструктор объекта сущности пакета, возвращается интерфейс пакета.
func New(item ...http.Header) Interface {
	var (
		i   int
		hdr = new(impl)
	)

	for i = range item {
		hdr.Header = item[i]
		return hdr
	}
	hdr.Header = make(http.Header)

	return hdr
}

// Add Добавление заголовка с именем ключа и значением.
func (hdr *impl) Add(key string, value string) {
	hdr.Header.Add(key, value)
}

// Del Удаление заголовка с именем ключа.
func (hdr *impl) Del(key string) {
	hdr.Header.Del(key)
}

// Get Получение первого значения заголовка с указанным ключём.
func (hdr *impl) Get(key string) string {
	return hdr.Header.Get(key)
}

// IsSet Проверка существует ли заголовок с указанным ключём.
func (hdr *impl) IsSet(key string) (ok bool) {
	_, ok = hdr.Header[key]
	return
}

// Set Установка значения заголовка с указанным ключём.
func (hdr *impl) Set(key string, value string) {
	hdr.Header.Set(key, value)
}

// Names Получение списка всех ключей заголовка.
func (hdr *impl) Names() (ret []string) {
	var key string

	for key = range hdr.Header {
		ret = append(ret, key)
	}

	return
}

// Len Получение количества ключей заголовка.
func (hdr *impl) Len() int {
	return len(hdr.Header)
}

// Reset Очистка заголовка от всех ключей и их значений.
func (hdr *impl) Reset() {
	for kn := range hdr.Header {
		delete(hdr.Header, kn)
	}
}
