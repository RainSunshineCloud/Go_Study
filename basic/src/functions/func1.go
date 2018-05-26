package main

import (
	"fmt"
	"strings"
)

//方式1
func main() {
	fmt.Println("方式1")
	//方式2
	strings.Map(func(r rune) rune { fmt.Println("方式2"); return r }, "string")
	var jacks r = func(r rune) rune { fmt.Println("方式3"); return r }
	strings.Map(jacks, "string")
}

//方式3
type r func(rune) rune
