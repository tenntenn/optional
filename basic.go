package optional

// String is an optional string value.
type String string

// NewString creates new a String.
func NewString(s string) *String {
	v := String(s)
	return &v
}

// String returns string value.
// When s is nil it returns empty string.
func (s *String) String() string {
	if s != nil {
		return string(*s)
	}
	return ""
}

// IsSet returns whether s is set as v or not.
func (s *String) IsSet(v string) bool {
	if s == nil {
		return false
	}
	return string(*s) == v
}

// Float32 is an optional float32 value.
type Float32 float32

// NewFloat32 creates new a Float32.
func NewFloat32(f float32) *Float32 {
	v := Float32(f)
	return &v
}

// Float32 returns float32 value.
// When f is nil it returns 0.
func (f *Float32) Float32() float32 {
	if f != nil {
		return float32(*f)
	}
	return 0
}

// IsSet returns whether f is set as v or not.
func (f *Float32) IsSet(v float32) bool {
	if f == nil {
		return false
	}
	return float32(*f) == v
}

// Float64 is an optional float64 value.
type Float64 float64

// NewFloat64 creates new a Float64.
func NewFloat64(f float64) *Float64 {
	v := Float64(f)
	return &v
}

// Float64 returns float64 value.
// When f is nil it returns 0.
func (f *Float64) Float64() float64 {
	if f != nil {
		return float64(*f)
	}
	return 0
}

// IsSet returns whether f is set as v or not.
func (f *Float64) IsSet(v float64) bool {
	if f == nil {
		return false
	}
	return float64(*f) == v
}

// Int is an optional int value.
type Int int

// NewInt creates new a Int.
func NewInt(n int) *Int {
	v := Int(n)
	return &v
}

// Int returns int value.
// When n is nil it returns 0.
func (n *Int) Int() int {
	if n != nil {
		return int(*n)
	}
	return 0
}

// IsSet returns whether n is set as v or not.
func (n *Int) IsSet(v int) bool {
	if n == nil {
		return false
	}
	return int(*n) == v
}

// Bool is an optional bool value.
type Bool bool

// NewBool creates new a Bool.
func NewBool(b bool) *Bool {
	v := Bool(b)
	return &v
}

// IsSet returns whether b is set as v or not.
func (b *Bool) IsSet(v bool) bool {
	if b == nil {
		return false
	}
	return bool(*b) == v
}
