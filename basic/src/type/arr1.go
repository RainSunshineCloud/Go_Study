package main

import "fmt"

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := arr1
	for key, val := range arr2 {
		arr2[key] = val * 2
	}

	fmt.Println(arr1, arr2)
}
