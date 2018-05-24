package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./file1.go")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	dst, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY, 0666)
	defer dst.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(dst, file)
}
