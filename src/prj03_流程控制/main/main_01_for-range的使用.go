package main

import "fmt"

func main() {
	var str string = "hello world草了"
	//方式1：普通for
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}
	fmt.Printf("\n")

	//方式2：for range
	for i, val := range str {
		fmt.Printf("索引: %d, 具体: %c\n", i, val)
	}
}
