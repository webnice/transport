package methods

import (
	"strings"
)

// New Function create new implementation of interface
func New() Interface {
	return new(Implementation)
}

// Options Return HTTP method OPTIONS
func (m *Implementation) Options() Value {
	return &methodType{optionsMethod}
}

// Get Return HTTP method GET
func (m *Implementation) Get() Value {
	return &methodType{getMethod}
}

// Head Return HTTP method GET
func (m *Implementation) Head() Value {
	return &methodType{headMethod}
}

// Post Return HTTP method POST
func (m *Implementation) Post() Value {
	return &methodType{postMethod}
}

// Put Return HTTP method PUT
func (m *Implementation) Put() Value {
	return &methodType{putMethod}
}

// Delete Return HTTP method DELETE
func (m *Implementation) Delete() Value {
	return &methodType{deleteMethod}
}

// Trace Return HTTP method TRACE
func (m *Implementation) Trace() Value {
	return &methodType{traceMethod}
}

// Connect Return HTTP method CONNECT
func (m *Implementation) Connect() Value {
	return &methodType{connectMethod}
}

// Parse string and return interface
func (m *Implementation) Parse(inp string) Value {
	var tmp string
	var key Type
	var ret *methodType

	tmp = strings.ToUpper(inp)
	for key = range maps {
		if maps[key] == tmp {
			ret = &methodType{key}
		}
	}

	return ret
}

// Parse string and return interface
func Parse(inp string) Value {
	return new(Implementation).Parse(inp)
}
