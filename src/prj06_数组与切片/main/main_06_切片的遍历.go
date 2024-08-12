package main

import "fmt"

func main() {
	slice := make([]int, 4, 20)
	slice[0] = 66
	slice[1] = 88
	slice[2] = 99
	slice[3] = 110

	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v] = %v\t", i, slice[i])
	}
	fmt.Println("\n------------------------------")
	for i, v := range slice {
		fmt.Printf("slice[%v] = %v\t", i, v)
	}
}
