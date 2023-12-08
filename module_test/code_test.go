package module_test

import (
	"github.com/btnguyen2k/mymodule"
	"reflect"
	"testing"
)

// sample tests

func TestClone_int(t *testing.T) {
	testName := "TestClone_int"
	input := []int{1, 2, 3}
	output := mymodule.Clone(input)
	if !reflect.DeepEqual(input, output) {
		t.Fatalf("%s failed: expected %v but received %v", testName, input, output)
	}
}

func TestClone_uint(t *testing.T) {
	testName := "TestClone_uint"
	input := []uint{4, 5, 6}
	output := mymodule.Clone(input)
	if !reflect.DeepEqual(input, output) {
		t.Fatalf("%s failed: expected %v but received %v", testName, input, output)
	}
}

func TestClone_float(t *testing.T) {
	testName := "TestClone_float"
	input := []float64{7.8, 8.9, 9.0}
	output := mymodule.Clone(input)
	if !reflect.DeepEqual(input, output) {
		t.Fatalf("%s failed: expected %v but received %v", testName, input, output)
	}
}

func TestClone_string(t *testing.T) {
	testName := "TestClone_string"
	input := []string{"string 1", "string 2", "string 3"}
	output := mymodule.Clone(input)
	if !reflect.DeepEqual(input, output) {
		t.Fatalf("%s failed: expected %v but received %v", testName, input, output)
	}
}
