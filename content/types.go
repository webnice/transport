package content

import (
	"io"

	"github.com/webnice/transport/v4/data"

	"golang.org/x/text/encoding"
)

// Interface Интерфейс пакета.
type Interface interface {
	io.WriterTo

	// Transcode Перекодирование контента из указанной кодировки в UTF-8.
	Transcode(e encoding.Encoding) Interface

	// Transform Трансформирование исходного контента путём пропуска контента через переданный в функции ридер.
	Transform(fn TransformFunc) Interface

	// String Получение контента в виде строки.
	String() (string, error)

	// Bytes Получение контента в виде среза байт.
	Bytes() ([]byte, error)

	// UnmarshalJson Декодирование контента в структуру, предполагается что контент является json.
	UnmarshalJson(o any) error

	// UnmarshalXml Декодирование контента в структуру, предполагается что контент является xml.
	UnmarshalXml(data any) error

	// UnTar Разархивация контента методом TAR.
	UnTar() Interface

	// UnZip Разархивация контента методом ZIP (извлекается только первый файл).
	UnZip() Interface

	// UnGzip Разархивация контента методом GZIP.
	UnGzip() Interface

	// UnFlate Разархивация контента методом FLATE.
	UnFlate() Interface

	// BackToBegin Перемещение точки чтения контента в начало контента.
	BackToBegin() error
}

// TransformFunc Описание функции конвертации и трансформации контента.
type TransformFunc func(r io.Reader) (io.Reader, error)

// Объект сущности пакета.
type impl struct {
	essence data.ReadAtSeekerWriteToCloser // Данные контента.
	rdc     io.ReadCloser                  // Интерфейс чтения контента.
	// Контент.
	transcode encoding.Encoding // Если не равно nil, то контент перекодируется на лету из указанной кодировки.
	transform TransformFunc     // Функция потокового преобразования контента.
	unZip     bool              // =true - контент сжатый алгоритмом сжатия ZIP, возвращается первый файл в архиве.
	unTar     bool              // =true - контент сжатый алгоритмом сжатия TAR.
	unGzip    bool              // =true - контент сжатый алгоритмом сжатия GZIP.
	unFlate   bool              // =true - контент сжатый алгоритмом сжатия FLATE.
}
