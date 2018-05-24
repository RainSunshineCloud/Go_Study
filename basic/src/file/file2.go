package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader

	fi, err := os.Open("./file.tar")

	closer, err := gzip.NewReader(fi)

	if err != nil {
		reader = bufio.NewReader(fi)

	} else {
		reader = bufio.NewReader(closer)

	}

	for {

		str, err := reader.ReadString('\n')
		fmt.Println(str)
		if err != nil {
			return
		}

	}
}
