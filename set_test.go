package optional_test

import (
	"testing"

	. "github.com/tenntenn/optional"
)

func TestSet(t *testing.T) {
	var (
		alwaysTrue = func(v1, v2 interface{}) bool {
			return true
		}

		alwaysFalse = func(v1, v2 interface{}) bool {
			return false
		}
	)
	data := []struct {
		S         *Set
		V         interface{}
		IsMatched bool
	}{
		{
			S:         nil,
			V:         nil,
			IsMatched: true,
		},
		{
			S:         NewSet(nil, nil).Put(1).Put(2).Put(3),
			V:         []int{1, 2, 3},
			IsMatched: true,
		},
		{
			S:         NewSet(nil, nil).Put(1).Put(2),
			V:         []int{1, 2, 3},
			IsMatched: true,
		},
		{
			S:         NewSet(nil, nil),
			V:         []int{1, 2, 3},
			IsMatched: true,
		},
		{
			S:         NewSet(nil, nil).Put(1),
			V:         []int{2, 1, 3},
			IsMatched: true,
		},
		{
			S:         NewSet(NewInt(3), nil),
			V:         []int{1, 2, 3},
			IsMatched: true,
		},
		{
			S:         NewSet(NewInt(3), nil).Put(100),
			V:         []int{1, 2, 3},
			IsMatched: false,
		},
		{
			S:         NewSet(NewInt(3), nil),
			V:         []int{1, 2},
			IsMatched: false,
		},
		{
			S:         NewSet(nil, alwaysFalse).Put(1),
			V:         []int{1, 2, 3},
			IsMatched: false,
		},
		{
			S:         NewSet(nil, alwaysTrue).Put(100),
			V:         []int{1, 2, 3},
			IsMatched: true,
		},
		// Seq
		{
			S:         NewSet(nil, nil).Put(1).Put(2).Put(3),
			V:         intSlice([]int{1, 2, 3}),
			IsMatched: true,
		},
		{
			S:         NewSet(nil, nil).Put(1).Put(2),
			V:         intSlice([]int{1, 2, 3}),
			IsMatched: true,
		},
		{
			S:         NewSet(nil, nil),
			V:         intSlice([]int{1, 2, 3}),
			IsMatched: true,
		},
		{
			S:         NewSet(nil, nil).Put(1),
			V:         intSlice([]int{2, 1, 3}),
			IsMatched: true,
		},
		{
			S:         NewSet(NewInt(3), nil),
			V:         intSlice([]int{1, 2, 3}),
			IsMatched: true,
		},
		{
			S:         NewSet(NewInt(3), nil).Put(100),
			V:         intSlice([]int{1, 2, 3}),
			IsMatched: false,
		},
		{
			S:         NewSet(NewInt(3), nil),
			V:         intSlice([]int{1, 2}),
			IsMatched: false,
		},
		{
			S:         NewSet(nil, alwaysFalse).Put(1),
			V:         intSlice([]int{1, 2, 3}),
			IsMatched: false,
		},
		{
			S:         NewSet(nil, alwaysTrue).Put(100),
			V:         intSlice([]int{1, 2, 3}),
			IsMatched: true,
		},
	}

	for i, d := range data {
		if isMatched := d.S.Match(d.V); isMatched != d.IsMatched {
			t.Errorf("data[%d].O.IsMatched expects %v, but %v", i, d.IsMatched, isMatched)
		}
	}
}
