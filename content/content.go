package content

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/flate"
	"compress/gzip"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"

	"github.com/webnice/transport/v4/charmap"
	"github.com/webnice/transport/v4/data"

	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

// New Конструктор объекта сущности пакета, возвращается интерфейс пакета.
func New(d data.ReadAtSeekerWriteToCloser) Interface { return &impl{essence: d} }

// WriteTo Реализация  интерфейса io.WriterTo.
func (cnt *impl) WriteTo(w io.Writer) (n int64, err error) {
	if err = cnt.ReaderCloser(); err != nil || cnt.rdc == nil {
		return
	}
	n, err = io.Copy(w, cnt.rdc)

	return
}

// ReaderCloser Получение io.ReadCloser для контента.
func (cnt *impl) ReaderCloser() (err error) {
	// Разархивация ZIP.
	if cnt.unZip {
		if cnt.rdc, err = cnt.UncompressZip(cnt.essence); err != nil {
			return
		}
	}
	// Разархивация TAR.
	if cnt.unTar {
		if cnt.rdc, err = cnt.UncompressTar(cnt.essence); err != nil {
			return
		}
	}
	// Разархивация GZIP.
	if cnt.unGzip {
		if cnt.rdc, err = cnt.UncompressGzip(cnt.essence); err != nil {
			return
		}
	}
	// Разархивация FLATE.
	if cnt.unFlate {
		if cnt.rdc, err = cnt.UncompressFlate(cnt.essence); err != nil {
			return
		}
	}
	// Перекодирование контента если установлен транскодер.
	if cnt.transcode != nil && cnt.essence != nil {
		// Создание ReadCloser из Reader + func Close.
		cnt.rdc = data.NewReadCloser(
			transform.NewReader(cnt.essence, cnt.transcode.NewDecoder()),
			cnt.essence.Close,
		)
	} else if cnt.rdc == nil && cnt.essence != nil {
		cnt.rdc = cnt.essence
	}
	// Преобразование контента если установлен трансформер.
	if cnt.transform != nil && cnt.rdc != nil {
		var newReader io.Reader
		newReader, err = cnt.transform(cnt.rdc)
		cnt.rdc = data.NewReadCloser(newReader, cnt.rdc.Close)
	}

	return
}

// UncompressZip Контент представлен ZIP архивом.
func (cnt *impl) UncompressZip(r data.ReadAtSeekerWriteToCloser) (rdr io.ReadCloser, err error) {
	const (
		errZip  = "чтение zip архива прервано ошибкой: %s"
		errNop  = "в zip архиве не найдено файлов"
		errFile = "открытие файла %q в zip архиве прервано ошибкой: %s"
	)
	var zipReader *zip.Reader

	if zipReader, err = zip.NewReader(r, r.Size()); err != nil {
		err = fmt.Errorf(errZip, err.Error())
		return
	}
	if len(zipReader.File) <= 0 {
		err = errors.New(errNop)
		return
	}
	if rdr, err = zipReader.File[0].Open(); err != nil {
		err = fmt.Errorf(errFile, zipReader.File[0].Name, err.Error())
		return
	}

	return
}

// UncompressTar Контент представлен TAR архивом.
func (cnt *impl) UncompressTar(r data.ReadAtSeekerWriteToCloser) (rdr io.ReadCloser, err error) {
	const (
		errNop = "в tar архиве не найдено файлов"
		errTar = "чтение tar архива прервано ошибкой: %s"
	)
	var tarReader *tar.Reader

	tarReader = tar.NewReader(r)
	_, err = tarReader.Next()
	if err == io.EOF {
		err = errors.New(errNop)
		return
	}
	if err != nil {
		err = fmt.Errorf(errTar, err.Error())
		return
	}
	rdr = data.NewReadCloser(tarReader, r.Close)

	return
}

// UncompressGzip Контент представлен GZIP архивом.
func (cnt *impl) UncompressGzip(r data.ReadAtSeekerWriteToCloser) (rdr io.ReadCloser, err error) {
	const errGzip = "чтение gzip архива прервано ошибкой: %s"
	var gzipReader *gzip.Reader

	switch gzipReader, err = gzip.NewReader(r); {
	case err != nil && err != io.EOF:
		err = fmt.Errorf(errGzip, err.Error())
		return
	case err == io.EOF:
		rdr, err = r, nil
		return
	}
	rdr = data.NewReadCloser(gzipReader, func() error { _ = gzipReader.Close(); return r.Close() })

	return
}

// UncompressFlate Контент представлен FLATE архивом.
func (cnt *impl) UncompressFlate(r data.ReadAtSeekerWriteToCloser) (rdr io.ReadCloser, err error) {
	const errFlate = "чтение flate архива прервано ошибкой: %s"
	var flateReader io.ReadCloser

	switch flateReader = flate.NewReader(r); {
	case flateReader == nil && err != io.EOF:
		err = errors.New(errFlate)
		return
	case err == io.EOF:
		rdr, err = r, nil
		return
	}
	rdr = data.NewReadCloser(flateReader, func() error { _ = flateReader.Close(); return r.Close() })

	return
}

// String Получение контента в виде строки.
func (cnt *impl) String() (ret string, err error) {
	var tmp = &bytes.Buffer{}

	if _, err = cnt.WriteTo(tmp); err != nil {
		return
	}
	ret = tmp.String()

	return
}

// Bytes Получение контента в виде среза байт.
func (cnt *impl) Bytes() (ret []byte, err error) {
	var tmp = &bytes.Buffer{}

	if _, err = cnt.WriteTo(tmp); err != nil {
		return
	}
	ret = tmp.Bytes()

	return
}

// Transcode Перекодирование контента из указанной кодировки в UTF-8.
func (cnt *impl) Transcode(from encoding.Encoding) Interface {
	return &impl{
		essence:   cnt.essence,
		transcode: from,
		transform: cnt.transform,
		unTar:     cnt.unTar,
		unZip:     cnt.unZip,
		unGzip:    cnt.unGzip,
		unFlate:   cnt.unFlate,
	}
}

// UnmarshalJson Декодирование контента в структуру, предполагается что контент является json.
func (cnt *impl) UnmarshalJson(data any) (err error) {
	var decoder *json.Decoder

	if err = cnt.ReaderCloser(); err == io.EOF {
		return nil
	} else if err != nil {
		return
	}
	defer func() { _ = cnt.rdc.Close() }()
	decoder = json.NewDecoder(cnt.rdc)
	err = decoder.Decode(data)

	return
}

// UnmarshalXml Декодирование контента в структуру, предполагается что контент является xml.
func (cnt *impl) UnmarshalXml(data any) (err error) {
	var decoder *xml.Decoder

	switch err = cnt.ReaderCloser(); {
	case err == io.EOF:
		return nil
	case err != nil:
		return
	}
	defer func() { _ = cnt.rdc.Close() }()
	decoder = xml.NewDecoder(cnt.rdc)
	decoder.CharsetReader = cnt.MakeCharsetReader()
	err = decoder.Decode(data)

	return
}

// MakeCharsetReader Создание функции потоковой конвертации данных.
func (cnt *impl) MakeCharsetReader() func(string, io.Reader) (io.Reader, error) {
	const errCode = "не найдена кодировка %q"

	return func(cs string, input io.Reader) (rd io.Reader, err error) {
		// Перекодирование контента на уровне вышестоящего ридера.
		if cnt.transcode != nil {
			rd = input
			return
		}
		// Поиск кодовой страницы.
		var enc = charmap.NewCharmap().FindByName(cs)
		if enc == nil {
			err = fmt.Errorf(errCode, cs)
			return
		}
		// Новый ридер с перекодированием.
		rd = data.NewReadCloser(
			transform.NewReader(input, enc.NewDecoder()),
			nil, // Поток будет закрыт в родительской функции, Closer не требуется.
		)
		return
	}
}

// Transform Трансформирование исходного контента путём пропуска контента через переданный в функции ридер.
func (cnt *impl) Transform(fn TransformFunc) Interface {
	return &impl{
		essence:   cnt.essence,
		transcode: cnt.transcode,
		transform: fn,
		unTar:     cnt.unTar,
		unZip:     cnt.unZip,
		unGzip:    cnt.unGzip,
		unFlate:   cnt.unFlate,
	}
}

// UnTar Разархивация контента методом TAR.
func (cnt *impl) UnTar() Interface {
	return &impl{
		essence:   cnt.essence,
		transcode: cnt.transcode,
		transform: cnt.transform,
		unTar:     true,
		unZip:     cnt.unZip,
		unGzip:    cnt.unGzip,
		unFlate:   cnt.unFlate,
	}
}

// UnZip Разархивация контента методом ZIP (извлекается только первый файл).
func (cnt *impl) UnZip() Interface {
	return &impl{
		essence:   cnt.essence,
		transcode: cnt.transcode,
		transform: cnt.transform,
		unTar:     cnt.unTar,
		unZip:     true,
		unGzip:    cnt.unGzip,
		unFlate:   cnt.unFlate,
	}
}

// UnGzip Разархивация контента методом GZIP.
func (cnt *impl) UnGzip() Interface {
	return &impl{
		essence:   cnt.essence,
		transcode: cnt.transcode,
		transform: cnt.transform,
		unTar:     cnt.unTar,
		unZip:     cnt.unZip,
		unGzip:    true,
		unFlate:   cnt.unFlate,
	}
}

// UnFlate Разархивация контента методом FLATE.
func (cnt *impl) UnFlate() Interface {
	return &impl{
		essence:   cnt.essence,
		transcode: cnt.transcode,
		transform: cnt.transform,
		unTar:     cnt.unTar,
		unZip:     cnt.unZip,
		unGzip:    cnt.unGzip,
		unFlate:   true,
	}
}

// BackToBegin Перемещение точки чтения контента в начало контента.
func (cnt *impl) BackToBegin() (err error) {
	if cnt.essence == nil {
		err = fmt.Errorf("request failed, response object is nil")
		return
	}
	_, err = cnt.essence.Seek(0, io.SeekStart)

	return
}
