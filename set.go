package optional

import "reflect"

// Set is option values which are not ordered.
type Set struct {
	m map[interface{}]struct{}
	// Len is length of the set.
	Len *Int
	// Eq uses comparation of elements.
	// If Eq is nil, == is used for comparation istead of Eq.
	Eq func(v1, v2 interface{}) bool
}

// NewSet creates new Set.
func NewSet(len *Int, eq func(v1, v2 interface{}) bool) *Set {
	return &Set{
		m:   make(map[interface{}]struct{}, len.Int()),
		Len: len,
		Eq:  eq,
	}
}

// Put adds v into the set.
func (s *Set) Put(v interface{}) *Set {
	if s.m == nil {
		s.m = map[interface{}]struct{}{}
	}
	s.m[v] = struct{}{}
	return s
}

// Delete deletes v from the set.
func (s *Set) Delete(v interface{}) *Set {
	delete(s.m, v)
	return s
}

func (s *Set) contains(v interface{}, seq Seq) bool {
	for i := 0; i < seq.Len(); i++ {
		if s.Eq != nil {
			if s.Eq(v, seq.At(i)) {
				return true
			}
		} else if v == seq.At(i) {
			return true
		}
	}
	return false
}

// Match tests whether slice satisfies this set.
func (s *Set) Match(v interface{}) bool {

	if s == nil {
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

	if !s.Len.Match(seq.Len()) {
		return false
	}

	for v := range s.m {
		if !s.contains(v, seq) {
			return false
		}
	}

	return true
}
