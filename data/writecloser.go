package data

import "io"

// NewWriteCloser Создание нового объекта на основе io.Writer
func NewWriteCloser(w io.Writer) WriteCloser {
	return &writeCloserImplementation{essence: w}
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
