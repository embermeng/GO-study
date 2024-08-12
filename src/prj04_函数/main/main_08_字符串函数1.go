package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 计算长度，按字节进行统计，汉字是3个字节
	str := "golang你好"
	fmt.Println(len(str))

	// 遍历
	// 方式1：for-range

	// 方式2：字符串转切片
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c\n", r[i])
	}

	// 字符串转整数
	num1, _ := strconv.Atoi("666")
	fmt.Println(num1)

	// 整数转字符串
	str1 := strconv.Itoa(888)
	fmt.Println(str1)

	// 统计一个字符串中有几个指定的子串
	count := strings.Count("golangandjava", "ga")
	fmt.Println(count)

	// 不分大小写的字符串比较
	isEqual := strings.EqualFold("hello", "HELLO")
	fmt.Println(isEqual)

	// 区分大小写的字符串比较
	fmt.Println("hello" == "HELLO")

	// 返回子串在字符串中第一次出现的索引值
	fmt.Println(strings.Index("golangandjavaga", "ga"))

}
