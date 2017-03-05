package optional

import (
	"errors"
	"reflect"
)

// Seq represents sequential values.
type Seq interface {
	// At returns i-th value.
	At(i int) interface{}
	Len() int
}

type seqValue struct {
	v reflect.Value
}

func toSeq(v reflect.Value) (Seq, error) {
	if v.Kind() != reflect.Slice &&
		v.Kind() != reflect.Array {
		err := errors.New("v must be slice or array")
		return nil, err
	}
	return &seqValue{v: v}, nil
}

func (v *seqValue) At(i int) interface{} {
	return v.v.Index(i).Interface()
}

func (v *seqValue) Len() int {
	return v.v.Len()
}
