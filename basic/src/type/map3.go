package main

import "fmt"

func main() {
	slice1 := []map[string]int{{"key": 1}}
	for _, val := range slice1 {
		if val, ok := val["key"]; ok {
			fmt.Println(val)
		}
	}

	slice2 := make([]map[string]int, 10)
	slice2[0] = make(map[string]int)
	slice2[0] = map[string]int{"key": 1}
	fmt.Println(slice2)
}
