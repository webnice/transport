package methods

// Int constants as named type
const (
	optionsMethod Type = 1 // OPTIONS
	getMethod     Type = 2 // GET
	headMethod    Type = 3 // HEAD
	postMethod    Type = 4 // POST
	putMethod     Type = 5 // PUT
	deleteMethod  Type = 6 // DELETE
	traceMethod   Type = 7 // TRACE
	connectMethod Type = 8 // CONNECT
)

// maps String constants
var maps = map[Type]string{
	optionsMethod: `OPTIONS`,
	getMethod:     `GET`,
	headMethod:    `HEAD`,
	postMethod:    `POST`,
	putMethod:     `PUT`,
	deleteMethod:  `DELETE`,
	traceMethod:   `TRACE`,
	connectMethod: `CONNECT`,
}

// Type Type of methods
type Type int

// Value Value is an interface of method
type Value interface {
	Int() int
	String() string
	Type() Type
}

// methodType is an implementation of Value
type methodType struct {
	value Type
}

// Interface is an methods interface
type Interface interface {
	// Options Return HTTP method OPTIONS
	Options() Value

	// Get Return HTTP method GET
	Get() Value

	// Head Return HTTP method GET
	Head() Value

	// Post Return HTTP method POST
	Post() Value

	// Put Return HTTP method PUT
	Put() Value

	// Delete Return HTTP method DELETE
	Delete() Value

	// Trace Return HTTP method TRACE
	Trace() Value

	// Connect Return HTTP method CONNECT
	Connect() Value

	// Parse string and return value interface
	Parse(inp string) Value
}

// Implementation is an methods implementation
type Implementation struct {
}
