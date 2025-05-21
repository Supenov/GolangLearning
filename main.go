package main

import "fmt"

func main() {
	arr := [5]int{5, 66, 7, 100, 1}

	for i := 0; i < 5; i++ {
		fmt.Println(string(rune(i))+":", arr[i])
	}
}
