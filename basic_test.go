package optional_test

import (
	"testing"

	. "github.com/tenntenn/optional"
)

func TestString(t *testing.T) {
	data := []struct {
		O         *String
		V         string
		IsMatched bool
	}{
		{O: NewString("test"), V: "test", IsMatched: true},
		{O: NewString("test"), V: "foo", IsMatched: false},
		{O: nil, V: "test", IsMatched: true},
	}

	for i, d := range data {
		if isMached := d.O.Match(d.V); isMached != d.IsMatched {
			t.Errorf("data[%d].O.Match expects %v, but %v", i, d.IsMatched, isMached)
		}
	}
}

func TestFloat64(t *testing.T) {
	data := []struct {
		O         *Float64
		V         float64
		IsMatched bool
	}{
		{O: NewFloat64(1.0), V: 1.0, IsMatched: true},
		{O: NewFloat64(1.0), V: 0.5, IsMatched: false},
		{O: nil, V: 1.0, IsMatched: true},
	}

	for i, d := range data {
		if isMached := d.O.Match(d.V); isMached != d.IsMatched {
			t.Errorf("data[%d].O.Match expects %v, but %v", i, d.IsMatched, isMached)
		}
	}
}

func TestFloat32(t *testing.T) {
	data := []struct {
		O         *Float32
		V         float32
		IsMatched bool
	}{
		{O: NewFloat32(1.0), V: 1.0, IsMatched: true},
		{O: NewFloat32(1.0), V: 0.5, IsMatched: false},
		{O: nil, V: 1.0, IsMatched: true},
	}

	for i, d := range data {
		if isMached := d.O.Match(d.V); isMached != d.IsMatched {
			t.Errorf("data[%d].O.Match expects %v, but %v", i, d.IsMatched, isMached)
		}
	}
}

func TestInt(t *testing.T) {
	data := []struct {
		O         *Int
		V         int
		IsMatched bool
	}{
		{O: NewInt(1), V: 1, IsMatched: true},
		{O: NewInt(1), V: 2, IsMatched: false},
		{O: nil, V: 1, IsMatched: true},
	}

	for i, d := range data {
		if isMached := d.O.Match(d.V); isMached != d.IsMatched {
			t.Errorf("data[%d].O.Match expects %v, but %v", i, d.IsMatched, isMached)
		}
	}
}

func TestBool(t *testing.T) {
	data := []struct {
		O         *Bool
		V         bool
		IsMatched bool
	}{
		{O: NewBool(true), V: true, IsMatched: true},
		{O: NewBool(true), V: false, IsMatched: false},
		{O: nil, V: false, IsMatched: true},
	}

	for i, d := range data {
		if isMached := d.O.Match(d.V); isMached != d.IsMatched {
			t.Errorf("data[%d].O.Match expects %v, but %v", i, d.IsMatched, isMached)
		}
	}
}
