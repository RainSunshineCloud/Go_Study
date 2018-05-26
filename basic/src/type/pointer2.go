package main

import "fmt"

func main() {
	i1 := 1
	i2 := 2
	i3 := &i1
	i4 := &i2
	i5 := *i3 + *i4
	i6 := *i3 * *i4
	i7 := 1 << uint(*i3)
	fmt.Println(i5, i6, i7)
}
