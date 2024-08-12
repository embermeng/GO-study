package main

import "fmt"

func main() {
	//string的使用
	var address string = "北京长城 110 hello!"
	fmt.Println(address)

	//go中不能修改字符串中的字符，只能重新赋值字符串
	//var str = "hello"
	//str[0] = 'a'
	address = "哈哈哈哈哈"
	fmt.Println(address)

	//字符串的两种表示形式
	//双引号
	str2 := "abc\ndef"
	fmt.Println(str2)
	//反引号
	str3 := `
	package main

	import "fmt"
	
	func main() {
	}
	`
	fmt.Println(str3)
	//当拼接操作很长时，可以分行写，要将+保留在上一行
	str4 := "hello" + "world" + "hello" + "world" +
		"hello" + "world"
	fmt.Println(str4)
}
