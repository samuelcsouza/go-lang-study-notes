package main

import "fmt"

func main() {

	var arr [10]int

	for i := 0; i < 10; i++ {
		arr[i] = i + 1
		fmt.Println(arr[i])
	}

	fmt.Println("\n************************\n")

	arrNum := [5]int{0, 2, 4, 6, 8}

	for i := 0; i < 5; i++ {
		fmt.Println(arrNum[i])
	}
}
