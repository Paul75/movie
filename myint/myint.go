package myint

import "errors"

var ErrorDivideByZero = errors.New("try to divide by zero")

type MyInt int32

func (m MyInt) Add(i int) MyInt {
	return m + MyInt(i)
}

func (m MyInt) Sub(i int) MyInt {
	return m - MyInt(i)
}

func (m MyInt) Multiply(i int) MyInt {
	return m * MyInt(i)
}

func (m MyInt) Divide(i int) (MyInt, error) {
	if i == 0 {
		return 0, ErrorDivideByZero
	}
	return m / MyInt(i), nil
}
