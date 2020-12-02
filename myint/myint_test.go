package myint_test

import (
	"movie/myint"
	"testing"
)

type testData struct {
	title  string
	start  myint.MyInt
	input  int
	output myint.MyInt
	err    error
}

func TestAdd(t *testing.T) {
	data := []testData{
		{"a", 1, 1, 2, nil},
		{"b", 1, 0, 1, nil},
	}
	for _, d := range data {
		if res := d.start.Add(d.input); res != d.output {
			t.Errorf("for test %v waiting for %v got %v", d.title, d.output, res)
		}
	}
}

func TestSub(t *testing.T) {
	data := []testData{
		{"a", 1, 1, 0, nil},
		{"b", 10, 1, 9, nil},
	}
	for _, d := range data {
		if res := d.start.Sub(d.input); res != d.output {
			t.Errorf("for test %v waiting for %v got %v", d.title, d.output, res)
		}
	}
}

func TestMultiply(t *testing.T) {
	data := []testData{
		{"a", 1, 1, 1, nil},
		{"b", 2, 2, 4, nil},
	}
	for _, d := range data {
		if res := d.start.Multiply(d.input); res != d.output {
			t.Errorf("for test %v waiting for %v got %v", d.title, d.output, res)
		}
	}
}

func TestDivide(t *testing.T) {
	data := []testData{
		{"a", 1, 1, 1, nil},
		{"b", 10, 2, 5, nil},
		{"c", 2, 0, 0, myint.ErrorDivideByZero},
	}
	for _, d := range data {
		t.Logf("try to execute test %v", d.title)
		if res, err := d.start.Divide(d.input); res != d.output || err != d.err {
			t.Errorf("for test %v waiting for %v got %v", d.title, d.output, res)
		} else {
			t.Logf("test %v passed succefully", d.title)
		}
	}
}
