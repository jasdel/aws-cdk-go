package jsii

import "time"

func S(v string) *string {
	return &v
}
func SVal(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func B(v bool) *bool {
	return &v
}
func BVal(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}

// String represents a JSII defined string type.
type String string

// JSIIString identifies the type as a JSII string for interfaces.
func (String) JSIIString() {}

// Number represents a JSII defined Number type.
type Number float64

// JSIINumber identifies the type as a JSII number for interfaces.
func (Number) JSIINumber() {}

// Date represents a JSII defined Date type.
type Date time.Time

// JSIIDate identifies the type as a JSII date for interfaces.
func (Date) JSIIDate() {}

// Bool represents a JSII defined Boolean type.
type Bool bool

// JSIIBool identifies the type as a JSII boolean for interfaces.
func (Bool) JSIIBool() {}

// JSON represents a JSII defined Json type.
type JSON struct{}

// JSIIJSON indentifies the type as a JSII Json for interfaces.
func (JSON) JSIIJSON() {}

// Any represents a JSII defined Any type.
type Any interface{}
