package main

import (
	"errors"
	"fmt"
	"log"
)

type PathError struct {
	path  string
	error error
}

func main() {
	var errs interface{}
	str := "j"
	errs = PathError{path: "/", error: errors.New(str)}
	switch errs.(type) {
	case PathError:
		fmt.Println(1)

	case int:
		fmt.Println(2)

	}

	//panic('3')

	defer func() {
		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
		}
		err := []int{1, 2}
		l := len(err)
		fmt.Println("12", l)
		return
	}()

	panic('3')

	fmt.Println("23")

}
