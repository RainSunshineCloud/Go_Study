package main

import (
	"encoding/json"
	"fmt"
)

type s struct {
	inter string
	data  interface{}
}

func main() {
	str := [3]int{1, 2, 4}
	f := func(s) {

	}
	fmt.Println(str)
	if bt, err := json.Marshal(str); err == nil {
		fmt.Println(bt)
		var v interface{}
		err = json.Unmarshal(bt, &v)
		fmt.Println(bt, err, v)
	} else {
		fmt.Println(err)
	}
}
