package methods // import "github.com/webdeskltd/transport/methods"

// Int Return method as int constant
func (mt *methodType) Int() int {
	return int(mt.value)
}

// String Return method as string constant
func (mt *methodType) String() string {
	return maps[mt.value]
}

// Type Return method as Type constant
func (mt *methodType) Type() Type {
	return mt.value
}
