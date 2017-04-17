package optional

import "reflect"

// Tuple is option values which are ordered.
type Tuple struct {
	m map[int]interface{}
	// Len is length of the tuple.
	Len *Int
	// Eq uses comparation of elements.
	// If Eq is nil, == is used for comparation istead of Eq.
	Eq func(v1, v2 interface{}) bool
}

// NewTupple creates new Tupple.
func NewTupple(len *Int, eq func(v1, v2 interface{}) bool) *Tuple {
	return &Tuple{
		m:   make(map[int]interface{}, len.Int()),
		Len: len,
		Eq:  eq,
	}
}

// At returns i-th value.
func (t *Tuple) At(i int) interface{} {
	return t.m[i]
}

// Put puts i-th value and returns itself.
func (t *Tuple) Put(i int, v interface{}) *Tuple {
	if t.m == nil {
		t.m = map[int]interface{}{}
	}
	t.m[i] = v

	return t
}

// PutAll puts i-th value with a map returns itself.
func (t *Tuple) PutAll(vs map[int]interface{}) *Tuple {
	if t.m == nil {
		t.m = map[int]interface{}{}
	}

	for i, v := range vs {
		t.m[i] = v
	}

	return t
}

// Match tests whether v has elements which are in the tupple.
func (t *Tuple) Match(v interface{}) bool {

	if t == nil {
		return true
	}

	var seq Seq

	switch v := v.(type) {
	case Seq:
		seq = v
	default:
		vv := reflect.ValueOf(v)
		s, err := toSeq(vv)
		if err != nil {
			return false
		}
		seq = s
	}

	if !t.Len.Match(seq.Len()) {
		return false
	}

	for i, v := range t.m {
		if i > seq.Len() {
			return false
		}

		if i < 0 {
			i = seq.Len() + i
		}

		if t.Eq != nil {
			if !t.Eq(v, seq.At(i)) {
				return false
			}
		} else if v != seq.At(i) {
			return false
		}
	}

	return true
}
