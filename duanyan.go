package main

import (
	"fmt"
	"io"
	"os"
)

// switch-types 断言升级版
func PrintType(any interface{}) {
	switch any.(type) {
	case int:
		fmt.Println("any is int")
	case string:
		fmt.Println("any is string")
	case bool:
		fmt.Println("any is bool")
	default:
		fmt.Println("any is other")
	}
}

func main() {
	var a interface{}
	a = "10"
	//T为具体类型，x类型不等于具体类型T
	b, err1 := a.(int)
	fmt.Printf("%v %v\n", b, err1)
	//T为具体类型，x类型等于具体类型T
	c, err2 := a.(string)
	fmt.Printf("%v %v\n", c, err2)
	//T为接口类型，x类型不满足具体类型T
	d, err3 := a.(io.ReadWriter)
	fmt.Printf("%v %v\n", d, err3)

	a = os.Stdout
	//T为接口类型，x类型满足具体类型T
	e, err4 := a.(io.ReadWriter)
	fmt.Printf("%T %v %v\n", e, e, err4)

	any := "1"
	PrintType(any)
}
