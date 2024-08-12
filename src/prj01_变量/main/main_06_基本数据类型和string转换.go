package main

import (
	"fmt"
	"strconv"
)

func main() {
	//基本数据类型和string转换

	//基本类型转string
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var c byte = 'h'

	//任意类型转string：用fmt.Sprintf转换，此函数会返回转换后的字符串
	var str string
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%q\n", str, str)
	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%q\n", str, str)
	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%q\n", str, str)
	str = fmt.Sprintf("%c", c)
	fmt.Printf("str type %T str=%q\n", str, str)

	//string转任意类型：使用strconv包的函数
	var str2 string = "true"
	var b2 bool
	//strconv.ParseBool()函数返回两个值(value bool, err error)
	b2, _ = strconv.ParseBool(str2)
	fmt.Printf("b2 type %T b=%v\n", b2, b2)

}
