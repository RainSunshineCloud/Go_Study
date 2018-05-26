package main

import "fmt"

func main() {
	//方式1
	var map1 map[string]string
	map1 = map[string]string{"one": "1", "two": "2", "three": "3"}
	each(map1)

	//方式2
	map2 := map[string]string{"one": "one", "two": "two", "three": "three"}
	each(map2)

	//方式3
	map3 := make(map[string]string)
	map3 = map[string]string{"one": "1", "two": "2", "three": "3"}
	each(map3)
}

func each(map1 map[string]string) {
	for key := range map1 {
		fmt.Println(key, "=>", map1[key])
	}
}
