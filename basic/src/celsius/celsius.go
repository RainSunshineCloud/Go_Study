package main

import (
	"fmt"
	"strconv"
)

type Temperature struct {
	num    int
	symbel string
}

func main() {
	t := new(Temperature)
	t.SetTemperature(8)
	fmt.Println(t)

}

func (t *Temperature) SetTemperature(nums int) {
	t.num = nums
	t.symbel = "℃"
}

func (t *Temperature) String() string {
	return strconv.Itoa(t.num) + t.symbel
}
