package main

import "fmt"

type copy struct {
	name    string
	sex     string
	hobbies []string
}

// 浅拷贝/深拷贝
// 浅拷贝：创建一个新的对象，只复制对象的引用，不复制对象本身，底层内存共享
// 深拷贝：新对象是一个完全独立的副本，新老对象不共享内存
func main() {
	p1 := copy{
		name:    "Gm",
		sex:     "男",
		hobbies: []string{"aa", "bb"},
	}
	p2 := p1
	p1.sex = "女"
	p1.hobbies[0] = "cc"
	fmt.Println(p2)

	var str1, str2 string
	str1 = "111"
	str2 = str1
	str1 = "222"
	fmt.Println(str2)

	var n1, n2 int
	n1 = 111
	n2 = n1
	n1 = 222
	fmt.Println(n2)

	var b1, b2 bool
	b1 = true
	b2 = b1
	b1 = false
	fmt.Println(b2)
}
