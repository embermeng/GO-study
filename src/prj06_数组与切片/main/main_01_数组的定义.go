package main

import "fmt"

func main() {
	var scores [5]int
	scores[0] = 100
	scores[1] = 200
	scores[2] = 300
	scores[3] = 400
	scores[4] = 500

	var sum int = 0

	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}
	var avg int = sum / len(scores)
	fmt.Println(sum, avg)
}
