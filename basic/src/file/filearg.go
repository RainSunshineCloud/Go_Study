package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println(strings.Trim(os.Args[1], "-"))
	}

}
