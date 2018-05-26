package main

import "fmt"

func main() {
	map1 := map[string][]int{"key": []int{1, 1, 2}}
	for key, val := range map1 {
		for key1, val1 := range val {
			fmt.Println(key, "=>", key1, val1)
		}
	}

	map2 := make(map[string]int)
	if val, ok := map2["key"]; !ok {
		fmt.Println(val)
	}

	delete(map1, "key")
	if val, ok := map1["key"]; ok {
		fmt.Println(val)
	}
}
