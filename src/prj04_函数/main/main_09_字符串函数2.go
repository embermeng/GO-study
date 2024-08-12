package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符串替换, n是替换的数量，-1是全部替换
	str1 := strings.Replace("goandjavagogo", "go", "fuck", -1)
	fmt.Println(str1)

	// 切割字符串
	arr := strings.Split("go-python-java", "-")
	fmt.Println(arr)

	// 大小写转换
	fmt.Println(strings.ToLower("GO"))
	fmt.Println(strings.ToUpper("go"))

	// 字符串左右两边空格去掉
	fmt.Println(strings.TrimSpace("   go and java   "))

	// 去掉左右指定的字符（去掉左边TrimLeft，右边TrimRight）
	fmt.Println(strings.Trim("!golang!", "!"))

	// 判断字符串是否以指定字符串开头
	fmt.Println(strings.HasPrefix("https://www.baidu.com", "https"))

	// ...结尾
	fmt.Println(strings.HasSuffix("https://www.baidu.com", ".com"))
}
