package main

import (
	"fmt"
	"math/big"
)

var stat fmt.State

func main() {
	var int64 int64 = 1

	bigint := big.NewInt(int64)
	bigint.Add(bigint, bigint).Mul(bigint, bigint)
	fmt.Println(bigint)

}
