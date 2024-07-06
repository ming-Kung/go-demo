package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "http World!!"
	fmt.Println(str)
	strs := "gm is best rd!!"
	fmt.Println(strs)
	fmt.Println(len(str))
	fmt.Println(len(strs))
	fmt.Println(utf8.RuneCountInString(strs))
	fmt.Println(strings.Index(strs, " "))

	profile := &struct {
		Name string
		HP   int
	}{
		Name: "rat",
		HP:   150,
	}
	fmt.Printf("使用'%%v' %v\n", profile)
	fmt.Printf("使用'%%+v' %+v\n", profile)
	fmt.Printf("使用'%%#v' %#v\n", profile)
	fmt.Printf("使用'%%T' %T\n", profile)

}
