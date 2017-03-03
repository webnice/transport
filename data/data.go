package data

import (
	"io"
)

// WriteCloser is an interface
type WriteCloser interface {
	Write([]byte) (int, error)
	Close() error
}

// ReadCloser is an interface
type ReadCloser interface {
	Read([]byte) (int, error)
	Close() error
}

// writeCloserImplementation is an implementation
type writeCloserImplementation struct {
	essence io.Writer
	closer  func() error
}

// readCloserImplementation is an implementation
type readCloserImplementation struct {
	essence io.Reader
	closer  func() error
}

// NewWriteCloser Создание нового объекта на основе io.Writer
func NewWriteCloser(w io.Writer) WriteCloser {
	return &writeCloserImplementation{essence: w}
}

// NewReadCloser Создание нового объекта на основе io.Reader
func NewReadCloser(w io.Reader, fn func() error) ReadCloser {
	return &readCloserImplementation{essence: w, closer: fn}
}

// Write Реализация Writer
func (wr *writeCloserImplementation) Write(p []byte) (int, error) {
	return wr.essence.Write(p)
}

// Close Реализация Close
func (wr *writeCloserImplementation) Close() error {
	if wr.closer != nil {
		return wr.closer()
	}
	return nil
}

// Read Реализация Writer
func (rd *readCloserImplementation) Read(p []byte) (int, error) {
	return rd.essence.Read(p)
}

// Close Реализация Close
func (rd *readCloserImplementation) Close() error {
	if rd.closer != nil {
		return rd.closer()
	}
	return nil
}
