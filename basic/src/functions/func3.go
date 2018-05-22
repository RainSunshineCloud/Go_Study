package main

import "fmt"

func main() {
	fmt.Println(fibocci(6))
	fmt.Println(even(10))
}

func fibocci(num uint) (uint, uint) {

	if num >= 2 {
		res1, _ := fibocci(num - 1)
		res2, _ := fibocci(num - 2)
		return res1 + res2, num
	}
	return num, num
}

/*
嵌套循环，通过nr -2n == 0 来判断是否是偶数
每次返回重新调用为nr-2
*/
func even(nr int) bool {
	if nr == 0 {
		return true
	}
	return odd(RevSign(nr) - 1)
}

func RevSign(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}

func odd(nr int) bool {
	if nr == 0 {
		return false
	}
	return even(RevSign(nr) - 1)
}
