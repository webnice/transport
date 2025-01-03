package response

import "bytes"

// Отображение ответа в режиме отладки.
func (r *Response) debugResponse(data []byte) {
	const prefixKey = "< "
	var (
		buf []byte
		tmp [][]byte
		i   int
	)

	defer func() { _ = recover() }()
	tmp, buf = bytes.Split(data, []byte("\n")), buf[:0]
	for i = range tmp {
		tmp[i] = bytes.TrimRight(tmp[i], "\r")
		buf = bytes.Join([][]byte{buf, []byte(prefixKey), tmp[i], []byte("\r\n")}, []byte(""))
	}
	r.debugFunc(bytes.TrimRight(buf, "\r\n"))
}
