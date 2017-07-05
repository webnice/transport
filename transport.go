package transport

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"log"
	"os"
	"runtime"

	"gopkg.in/webnice/transport.v1/methods"
)

// NewTransport Function create new transport implementation
func NewTransport() Transport {
	var obj = new(implementation)
	runtime.SetFinalizer(obj, destructor)
	return obj
}

// NewMethod Function create new methods implementation and returns interface
func NewMethod() methods.Interface { return methods.New() }

// destructor Вызывается при уничтожении объекта
func destructor(obj *implementation) {
	//log.Debug(" --- Запуск деструктора")
	for i := range obj.CollectionOfTemporaryFiles {
		//log.Debug(" ---- Удаление файла: %s", obj.CollectionOfTemporaryFiles[i])
		if obj.CollectionOfTemporaryFiles[i] != "" {
			if err := os.Remove(obj.CollectionOfTemporaryFiles[i]); err != nil {
				log.Printf("Warning, delete temporary file '%s' error: %s", obj.CollectionOfTemporaryFiles[i], err)
			}
		}
	}
}

// Method Создание нового объекта метода и возврат интерфейса к его реализации
func (t *implementation) Method() methods.Interface { return methods.New() }

// TemporaryFile Вызов коллекционирует временные файлы, которые необходимо удалить при уничтожении объекта
func (t *implementation) TemporaryFile(fileName string) {
	t.CollectionOfTemporaryFiles = append(t.CollectionOfTemporaryFiles, fileName)
	return
}

// NewRequest Создание нового запроса
func (t *implementation) NewRequest(m methods.Value) Request {
	var req = new(requestImplementation)
	if m != nil {
		req.RequestMethod = m
	} else {
		req.RequestMethod = methods.New().Get()
	}

	// Установка функции коллекционирования временных файлов
	req.collectionOfTemporaryFilesFn = t.TemporaryFile

	return req
}
