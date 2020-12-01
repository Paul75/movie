package myint

import "fmt"

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

func (m MyInt) Divide(i int) MyInt {
	return m / MyInt(i)
}

func test() {
	var i MyInt = 3
	fmt.Println(i.Add(4).Sub(3).Multiply(5))
}
