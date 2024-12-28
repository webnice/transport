package header

import "net/http"

// Interface Интерфейс пакета.
type Interface interface {
	// Add Добавление заголовка с именем ключа и значением.
	Add(key string, value string)

	// Del Удаление заголовка с именем ключа.
	Del(key string)

	// Get Получение первого значения заголовка с указанным ключём.
	Get(key string) string

	// IsSet Проверка существует ли заголовок с указанным ключём.
	IsSet(key string) bool

	// Set Установка значения заголовка с указанным ключём.
	Set(key string, value string)

	// Names Получение списка всех ключей заголовка.
	Names() (ret []string)

	// Len Получение количества ключей заголовка.
	Len() int

	// Reset Очистка заголовка от всех ключей и их значений.
	Reset()
}

// Объект сущности пакета.
type impl struct {
	Header http.Header
}
