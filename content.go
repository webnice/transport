package transport

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/webnice/transport.v1/charmap"
	"gopkg.in/webnice/transport.v1/data"

	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

// Write Запись контента в io.Writer
func (cnt *contentImplementation) Write(wr io.Writer) (err error) {
	var rdc io.ReadCloser

	if rdc, err = cnt.ReaderCloser(); err == io.EOF {
		err = nil
		return
	} else if err != nil {
		return
	}
	defer func() {
		if e := rdc.Close(); e != nil {
			log.Printf("Warning, close temporary file '%s' error: %s", cnt.ResponseFHName, e)
		}
	}()
	if _, err = io.Copy(wr, rdc); err == io.EOF {
		err = nil
	}

	return
}

// ReaderCloser Получение io.ReadCloser для контента
func (cnt *contentImplementation) ReaderCloser() (rdr io.ReadCloser, err error) {
	var fh *os.File
	var tmp io.ReadCloser

	if fh, err = os.Open(cnt.ResponseFHName); err != nil {
		err = fmt.Errorf("Error to open temporary file '%s': %s", cnt.ResponseFHName, err)
		return
	}
	tmp = fh

	// Разархивация ZIP
	if cnt.unzip {
		if tmp, err = cnt.UncompressZip(fh); err != nil {
			return
		}
	}

	// Разархивация TAR
	if cnt.untar {
		if tmp, err = cnt.UncompressTar(fh); err != nil {
			return
		}
	}

	// Разархивация GZIP
	if cnt.ungzip {
		if tmp, err = cnt.UncompressGzip(fh); err != nil {
			return
		}
	}

	// Перекодирование контента если установлен транскодер
	if cnt.transcode != nil {
		// Создание ReadCloser из Reader + func Close
		rdr = data.NewReadCloser(
			transform.NewReader(tmp, cnt.transcode.NewDecoder()),
			tmp.Close,
		)
	} else {
		rdr = tmp
	}

	// Преобразование контента если установлен трансформер
	if cnt.transform != nil {
		var newReader io.Reader
		newReader, err = cnt.transform(rdr)
		rdr = data.NewReadCloser(newReader, rdr.Close)
	}

	return
}

// UncompressZip Uncompress content as zip
func (cnt *contentImplementation) UncompressZip(fh *os.File) (rdr io.ReadCloser, err error) {
	var zipReader *zip.Reader
	var fi os.FileInfo

	if fi, err = fh.Stat(); err != nil {
		return
	}
	if zipReader, err = zip.NewReader(fh, fi.Size()); err != nil {
		err = fmt.Errorf("Zip archive error: %s", err.Error())
		return
	}
	if len(zipReader.File) <= 0 {
		err = fmt.Errorf("There are no files in the archive")
		return
	}
	rdr, err = zipReader.File[0].Open()
	if err != nil {
		err = fmt.Errorf("Zip archive error, can't open file '%s': %s", zipReader.File[0].Name, err.Error())
		return
	}

	return
}

// UncompressTar Uncompress content as tar
func (cnt *contentImplementation) UncompressTar(fh *os.File) (rdr io.ReadCloser, err error) {
	var tarReader *tar.Reader

	tarReader = tar.NewReader(fh)
	_, err = tarReader.Next()
	if err == io.EOF {
		err = fmt.Errorf("There are no files in the archive")
		return
	}
	if err != nil {
		err = fmt.Errorf("Tar archive error: %s", err.Error())
		return
	}
	rdr = data.NewReadCloser(
		tarReader,
		fh.Close,
	)

	return
}

// UncompressGzip Uncompress content as gzip
func (cnt *contentImplementation) UncompressGzip(fh io.ReadCloser) (rdr io.ReadCloser, err error) {
	var gzipReader *gzip.Reader

	gzipReader, err = gzip.NewReader(fh)
	if err != nil {
		err = fmt.Errorf("GZIP content error: %s", err.Error())
		return
	}
	rdr = data.NewReadCloser(
		gzipReader,
		func() error {
			_ = gzipReader.Close()
			return fh.Close()
		},
	)

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
		transform:      cnt.transform,
		untar:          cnt.untar,
		unzip:          cnt.unzip,
	}
}

// ContentUnmarshalJSON Декодирование контента в структуру, предполагается что контент является json
func (cnt *contentImplementation) ContentUnmarshalJSON(i interface{}) (err error) {
	var rdc io.ReadCloser
	var decoder *json.Decoder

	if rdc, err = cnt.ReaderCloser(); err == io.EOF {
		err = nil
		return
	} else if err != nil {
		return
	}
	defer func() {
		if e := rdc.Close(); e != nil {
			log.Printf("Warning, close temporary file '%s' error: %s", cnt.ResponseFHName, e)
		}
	}()

	decoder = json.NewDecoder(rdc)
	err = decoder.Decode(i)

	return
}

// ContentUnmarshalXML Декодирование контента в структуру, предполагается что контент является xml
func (cnt *contentImplementation) ContentUnmarshalXML(i interface{}) (err error) {
	var rdc io.ReadCloser
	var decoder *xml.Decoder

	if rdc, err = cnt.ReaderCloser(); err == io.EOF {
		err = nil
		return
	} else if err != nil {
		return
	}
	defer func() {
		if e := rdc.Close(); e != nil {
			log.Printf("Warning, close temporary file '%s' error: %s", cnt.ResponseFHName, e)
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

// Transform Трансформирование исходного контента путём пропуска контента через переданный в функции ридер
func (cnt *contentImplementation) Transform(fn TransformFunc) ContentInterface {
	return &contentImplementation{
		ResponseFHName: cnt.ResponseFHName,
		ResponseFH:     cnt.ResponseFH,
		transcode:      cnt.transcode,
		transform:      fn,
		untar:          cnt.untar,
		unzip:          cnt.unzip,
		ungzip:         cnt.ungzip,
	}
}

// Unzip Разархивация контента методом TAR
func (cnt *contentImplementation) Untar() ContentInterface {
	return &contentImplementation{
		ResponseFHName: cnt.ResponseFHName,
		ResponseFH:     cnt.ResponseFH,
		transcode:      cnt.transcode,
		transform:      cnt.transform,
		untar:          true,
		unzip:          cnt.unzip,
		ungzip:         cnt.ungzip,
	}
}

// Unzip Разархивация контента методом ZIP
func (cnt *contentImplementation) Unzip() ContentInterface {
	return &contentImplementation{
		ResponseFHName: cnt.ResponseFHName,
		ResponseFH:     cnt.ResponseFH,
		transcode:      cnt.transcode,
		transform:      cnt.transform,
		untar:          cnt.untar,
		unzip:          true,
		ungzip:         cnt.ungzip,
	}
}

// UnGzip Разархивация контента методом GZIP
func (cnt *contentImplementation) UnGzip() ContentInterface {
	return &contentImplementation{
		ResponseFHName: cnt.ResponseFHName,
		ResponseFH:     cnt.ResponseFH,
		transcode:      cnt.transcode,
		transform:      cnt.transform,
		untar:          cnt.untar,
		unzip:          cnt.unzip,
		ungzip:         true,
	}
}
