package main

import "fmt"

func main() {
	convertInt64ToInt(10000000000000000000000)
}

type error interface{}

const (
	UINT_MIN uint = 0
	INT_MAX       = ^UINT_MIN >> 1
	INT_MIN       = ^int(INT_MAX)
)

func convertInt64ToInt(i64 int64) (int, err error) {
	defer func() {
		if err = recover(); err != nil {
			err = fmt.Errorf("Err:", err)
		}
	}()

	if i64 > int64(INT_MAX) {
		panic("超出INT最大取值范围无法转化")
	} else if i64 < int64(INT_MIN) {
		panic("超出INT最小取值范围无法转化")
	}

	return i64, nil
}
