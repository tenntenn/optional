package optional_test

import (
	"testing"

	. "github.com/tenntenn/optional"
)

type intSlice []int

var _ Seq = intSlice(nil)

func (s intSlice) At(i int) interface{} {
	return s[i]
}

func (s intSlice) Len() int {
	return len(s)
}

func TestTuple(t *testing.T) {
	var (
		alwaysTrue = func(v1, v2 interface{}) bool {
			return true
		}

		alwaysFalse = func(v1, v2 interface{}) bool {
			return false
		}
	)
	data := []struct {
		T         *Tuple
		V         interface{}
		IsMatched bool
	}{
		{
			T:         nil,
			V:         nil,
			IsMatched: true,
		},
		{
			T:         NewTupple(nil, nil).Put(0, 1).Put(1, 2).Put(2, 3),
			V:         []int{1, 2, 3},
			IsMatched: true,
		},
		{
			T:         NewTupple(nil, nil).Put(0, 1).Put(2, 3),
			V:         []int{1, 2, 3},
			IsMatched: true,
		},
		{
			T:         NewTupple(nil, nil),
			V:         []int{1, 2, 3},
			IsMatched: true,
		},
		{
			T:         NewTupple(nil, nil).Put(0, 1),
			V:         []int{2, 1, 3},
			IsMatched: false,
		},
		{
			T:         NewTupple(NewInt(3), nil),
			V:         []int{1, 2, 3},
			IsMatched: true,
		},
		{
			T:         NewTupple(NewInt(3), nil).Put(0, 100),
			V:         []int{1, 2, 3},
			IsMatched: false,
		},
		{
			T:         NewTupple(NewInt(3), nil),
			V:         []int{1, 2},
			IsMatched: false,
		},
		{
			T:         NewTupple(nil, alwaysFalse).Put(0, 1),
			V:         []int{1, 2, 3},
			IsMatched: false,
		},
		{
			T:         NewTupple(nil, alwaysTrue).Put(0, 100),
			V:         []int{1, 2, 3},
			IsMatched: true,
		},
		// Seq
		{
			T:         NewTupple(nil, nil).Put(0, 1).Put(1, 2).Put(2, 3),
			V:         intSlice([]int{1, 2, 3}),
			IsMatched: true,
		},
		{
			T:         NewTupple(nil, nil).Put(0, 1).Put(2, 3),
			V:         intSlice([]int{1, 2, 3}),
			IsMatched: true,
		},
		{
			T:         NewTupple(nil, nil),
			V:         intSlice([]int{1, 2, 3}),
			IsMatched: true,
		},
		{
			T:         NewTupple(nil, nil).Put(0, 1),
			V:         intSlice([]int{2, 1, 3}),
			IsMatched: false,
		},
		{
			T:         NewTupple(NewInt(3), nil),
			V:         intSlice([]int{1, 2, 3}),
			IsMatched: true,
		},
		{
			T:         NewTupple(NewInt(3), nil).Put(0, 100),
			V:         intSlice([]int{1, 2, 3}),
			IsMatched: false,
		},
		{
			T:         NewTupple(NewInt(3), nil),
			V:         intSlice([]int{1, 2}),
			IsMatched: false,
		},
		{
			T:         NewTupple(nil, alwaysFalse).Put(0, 1),
			V:         intSlice([]int{1, 2, 3}),
			IsMatched: false,
		},
		{
			T:         NewTupple(nil, alwaysTrue).Put(0, 100),
			V:         intSlice([]int{1, 2, 3}),
			IsMatched: true,
		},
	}

	for i, d := range data {
		if isMatched := d.T.Match(d.V); isMatched != d.IsMatched {
			t.Errorf("data[%d].O.IsMatched expects %v, but %v", i, d.IsMatched, isMatched)
		}
	}
}
