package main

import (
	"flag"
	"fmt"
)

var name, usage string = "n", "sd"
var value bool = true
var Bool = flag.Bool(name, value, usage)

func main() {
	fmt.Println(*Bool)
	flag.Parse()
	for i := 0; i < flag.NArg(); i++ {
		fmt.Println(flag.Arg(i))
		fmt.Println(*Bool)
	}
}
