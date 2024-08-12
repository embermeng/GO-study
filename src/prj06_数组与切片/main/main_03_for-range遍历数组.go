package main

import "fmt"

func main() {
	arr1 := [...]string{"tom", "jerry", "allen"}
	for index, value := range arr1 {
		fmt.Println(index, value)
	}
}
