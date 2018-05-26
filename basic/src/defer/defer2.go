package main

import "fmt"

/*
	1.defer 对于控制语句 只有对应的defer会执行
	2.defer 对于goto 只有对应的defer会执行
	3.defer 对于label 所有的label里面的defer都会执行

*/
func main() {
	defer fmt.Println("defer 1")

	for i := 1; i < 10; i++ {
		if i == 2 {
			continue
		}

		if i == 5 {
			goto label1
		}

		if i == 7 {
			goto label2
		}
		defer fmt.Println("defer", i)
	}

	if false {
		defer fmt.Println("defer false")
	}

	switch {
	case 1 == 2:
		goto label3
		defer fmt.Println("defer switch")
	}

label1:
	defer fmt.Println("defer label")
label2:
	defer fmt.Println("defer label")
label3:
	defer fmt.Println("defer label")
}
