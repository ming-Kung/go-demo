package main

import (
	"fmt"
	"sync"
)

var global *int

func f() {
	var x int = 1
	global = &x
}

func main() {
	f()
	fmt.Println("*global:", *global)
	a := make([]int, 2)
	b := make([]int, 2, 10)
	var c [2]int = [2]int{1, 2}
	fmt.Println(a, b) //打印：[0,0] [0,0]
	fmt.Println(len(a), len(b))
	fmt.Println(c, len(c))

	//声明
	var scene sync.Map
	//将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	//从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))

	//根据键删除对应的键值对
	scene.Delete("london")

	//遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("interate:", k, v)
		return true
	})
}
