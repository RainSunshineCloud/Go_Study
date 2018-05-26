package main

import "fmt"

func main() {

	fmt.Println(fibonacci())

}

func fibonacci() [50]uint {
	var arr = [50]uint{1, 1}
	for i := 2; i < 50; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr
}
