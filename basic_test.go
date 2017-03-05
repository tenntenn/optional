package optional_test

import (
	"testing"

	. "github.com/tenntenn/optional"
)

func TestString(t *testing.T) {
	data := []struct {
		O     *String
		V     string
		IsSet bool
	}{
		{O: NewString("test"), V: "test", IsSet: true},
		{O: NewString("test"), V: "foo", IsSet: false},
		{O: nil, V: "test", IsSet: false},
	}

	for i, d := range data {
		if isSet := d.O.IsSet(d.V); isSet != d.IsSet {
			t.Errorf("data[%d].O.IsSet expects %v, but %v", i, d.IsSet, isSet)
		}
	}
}

func TestFloat64(t *testing.T) {
	data := []struct {
		O     *Float64
		V     float64
		IsSet bool
	}{
		{O: NewFloat64(1.0), V: 1.0, IsSet: true},
		{O: NewFloat64(1.0), V: 0.5, IsSet: false},
		{O: nil, V: 1.0, IsSet: false},
	}

	for i, d := range data {
		if isSet := d.O.IsSet(d.V); isSet != d.IsSet {
			t.Errorf("data[%d].O.IsSet expects %v, but %v", i, d.IsSet, isSet)
		}
	}
}

func TestFloat32(t *testing.T) {
	data := []struct {
		O     *Float32
		V     float32
		IsSet bool
	}{
		{O: NewFloat32(1.0), V: 1.0, IsSet: true},
		{O: NewFloat32(1.0), V: 0.5, IsSet: false},
		{O: nil, V: 1.0, IsSet: false},
	}

	for i, d := range data {
		if isSet := d.O.IsSet(d.V); isSet != d.IsSet {
			t.Errorf("data[%d].O.IsSet expects %v, but %v", i, d.IsSet, isSet)
		}
	}
}

func TestInt(t *testing.T) {
	data := []struct {
		O     *Int
		V     int
		IsSet bool
	}{
		{O: NewInt(1), V: 1, IsSet: true},
		{O: NewInt(1), V: 2, IsSet: false},
		{O: nil, V: 1, IsSet: false},
	}

	for i, d := range data {
		if isSet := d.O.IsSet(d.V); isSet != d.IsSet {
			t.Errorf("data[%d].O.IsSet expects %v, but %v", i, d.IsSet, isSet)
		}
	}
}

func TestBool(t *testing.T) {
	data := []struct {
		O     *Bool
		V     bool
		IsSet bool
	}{
		{O: NewBool(true), V: true, IsSet: true},
		{O: NewBool(true), V: false, IsSet: false},
		{O: nil, V: false, IsSet: false},
	}

	for i, d := range data {
		if isSet := d.O.IsSet(d.V); isSet != d.IsSet {
			t.Errorf("data[%d].O.IsSet expects %v, but %v", i, d.IsSet, isSet)
		}
	}
}
