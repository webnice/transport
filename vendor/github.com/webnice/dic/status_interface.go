package dic

// Объект-одиночка справочника статусов HTTP ответов.
var singletonStatus Statuses

// IStatus Интерфейс статусов HTTP ответов.
type IStatus interface {
	// IsEqual Истина, если статусы HTTP ответов эквивалентны.
	IsEqual(m IStatus) bool

	// IsEqualCode Истина, если коды статусо HTTP ответов эквивалентны.
	IsEqualCode(c int) bool

	// Code Числовой код статуса HTTP ответа.
	Code() int

	// String Полностью весть MIME тип с подтипом и опциями.
	String() string

	// Bytes Статус HTTP ответа в виде среза байт.
	Bytes() []byte
}

// Структура объекта статуса HTTP ответа.
type tStatus struct {
	status string // Эльфийское название статуса.
	code   int    // Числовой код статуса.
}
