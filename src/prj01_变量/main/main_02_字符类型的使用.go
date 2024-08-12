package main

import "fmt"

func main() {
	// 字符类型的使用
	var c1 byte = 'a'
	var c2 byte = '0'

	// 若直接输出byte值，就是输出对应字符的ASCII码值
	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	// 使用格式化输出才能输出字符
	fmt.Printf("c1=%c c2=%c\n", c1, c2)

	// ASCII码值大于255，就会溢出，一个汉字3个字节
	//var c3 byte = '东'
	// 可以用int保存
	var c3 int = '东'
	fmt.Printf("c3=%c c3对应码值=%d \n", c3, c3)

	// 可以给变量赋一个数字，用%c输出一个字符
	var c4 int = 22269
	fmt.Printf("c4=%c\n", c4)

	// 字符相当于一个整数，可以按照码值运算
	var c5 = 'a' + 10
	fmt.Printf("c5=%c\n", c5)
}
