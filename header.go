package transport // import "github.com/webdeskltd/transport"

// Add adds the key, value pair to the header.
// It appends to any existing values associated with key.
func (h *headerImplementation) Add(key string, value string) {
	h.Header.Add(key, value)
}

// Del deletes the values associated with key.
func (h *headerImplementation) Del(key string) {
	h.Header.Del(key)
}

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns "".
// To access multiple values of a key, access the map directly
// with CanonicalHeaderKey.
func (h *headerImplementation) Get(key string) string {
	return h.Header.Get(key)
}

// Set sets the header entries associated with key to
// the single element value.  It replaces any existing
// values associated with key.
func (h *headerImplementation) Set(key string, value string) {
	h.Header.Set(key, value)
}

// Names Получение списка всех имён заголовков
func (h *headerImplementation) Names() (ret []string) {
	for kn := range h.Header {
		ret = append(ret, kn)
	}
	return
}
