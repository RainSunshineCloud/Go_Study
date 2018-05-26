package main

import "fmt"

func main() {
	//扩len (最大为cap)
	slice2 := make([]int, 2, 2)
	fmt.Println(len(slice2), slice2)
	slice2 = slice2[:cap(slice2)]
	fmt.Println(len(slice2))
	//扩cap

	slice3 := make([]int, 3, 15)
	n := copy(slice3, slice2) //copy个数
	fmt.Println(cap(slice3), n)

	//append (产生新的slice
	fmt.Println(&slice2[0])
	slice4 := append(slice2, 1)
	fmt.Println(slice4, &slice4[0])
}
