package transport // import "github.com/webdeskltd/transport"

//import "github.com/webdeskltd/debug"
import "github.com/webdeskltd/log"
import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/webdeskltd/transport/charmap"
	"github.com/webdeskltd/transport/data"

	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

// Write Запись контента в io.Writer
func (cnt *contentImplementation) Write(wr io.Writer) (err error) {
	var rdc io.ReadCloser

	if rdc, err = cnt.ReaderCloser(); err != nil {
		return
	}
	defer func() {
		if err = rdc.Close(); err != nil {
			log.Warning("Erro close temporary file '%s': %s", cnt.ResponseFHName, err)
		}
	}()
	_, err = io.Copy(wr, rdc)
	return
}

// ReaderCloser Получение io.ReadCloser для контента
func (cnt *contentImplementation) ReaderCloser() (rdr io.ReadCloser, err error) {
	var fh *os.File
	fh, err = os.Open(cnt.ResponseFHName)
	if err != nil {
		err = fmt.Errorf("Error to open temporary file '%s': %s", cnt.ResponseFHName, err)
		return
	}

	// Перекодирование контента если установлена кодировка from
	if cnt.transcode != nil {
		// Создание ReadCloser из Reader + func Close
		rdr = data.NewReadCloser(
			transform.NewReader(fh, cnt.transcode.NewDecoder()),
			fh.Close,
		)
	} else {
		rdr = fh
	}

	return
}

// String Получение контента в виде строки
func (cnt *contentImplementation) String() (ret string, err error) {
	var tmp *bytes.Buffer
	tmp = bytes.NewBuffer(nil)
	err = cnt.Write(tmp)
	ret = tmp.String()
	return
}

// Bytes Получение контента в виде среза байт
func (cnt *contentImplementation) Bytes() (ret []byte, err error) {
	var tmp *bytes.Buffer
	tmp = bytes.NewBuffer(nil)
	err = cnt.Write(tmp)
	ret = tmp.Bytes()
	return
}

// Transcode Перекодирование контента из указанной кодировки в UTF-8
func (cnt *contentImplementation) Transcode(from encoding.Encoding) ContentInterface {
	return &contentImplementation{
		ResponseFHName: cnt.ResponseFHName,
		ResponseFH:     cnt.ResponseFH,
		transcode:      from,
	}
	return cnt
}

// UnmarshalJSON Декодирование контента в структуру, предполагается что контент является json
func (cnt *contentImplementation) UnmarshalJSON(i interface{}) (err error) {
	var rdc io.ReadCloser
	var decoder *json.Decoder

	if rdc, err = cnt.ReaderCloser(); err != nil {
		return
	}
	defer func() {
		if err = rdc.Close(); err != nil {
			log.Warning("Erro close temporary file '%s': %s", cnt.ResponseFHName, err)
		}
	}()

	decoder = json.NewDecoder(rdc)
	err = decoder.Decode(i)

	return
}

// UnmarshalXML Декодирование контента в структуру, предполагается что контент является xml
func (cnt *contentImplementation) UnmarshalXML(i interface{}) (err error) {
	var rdc io.ReadCloser
	var decoder *xml.Decoder

	if rdc, err = cnt.ReaderCloser(); err != nil {
		return
	}
	defer func() {
		if err = rdc.Close(); err != nil {
			log.Warning("Erro close temporary file '%s': %s", cnt.ResponseFHName, err)
		}
	}()

	decoder = xml.NewDecoder(rdc)
	decoder.CharsetReader = cnt.MakeCharsetReader()
	err = decoder.Decode(i)

	return
}

// MakeCharsetReader Создание функции для потокового чтения данных с перекодированием
func (cnt *contentImplementation) MakeCharsetReader() func(string, io.Reader) (io.Reader, error) {
	return func(cs string, input io.Reader) (rd io.Reader, err error) {
		// Перекодирование контента на уровне вышестоящего ридера
		if cnt.transcode != nil {
			rd = input
			return
		}

		// Поиск кодовой страницы
		var enc = charmap.NewCharmap().FindByName(cs)
		if enc == nil {
			err = fmt.Errorf("Could not find the code page '%s'", cs)
			return
		}

		// Новый ридер с перекодированием
		rd = data.NewReadCloser(
			transform.NewReader(input, enc.NewDecoder()),
			nil, // Поток будет закрыт в родительской функции, Closer не требуется
		)
		return
	}
}
