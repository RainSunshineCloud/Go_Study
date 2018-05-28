package main

import "fmt"

type obj interface{}
type mf func(obj) obj

/*
注意：空接口虽然通过reflect.TypeOf返回的是值本身的类型，但是其还是一个接口。
接口转换时需要使用接口断言，不能使用强制转换。而其他类型转化为接口是隐形的。
*/
func main() {
	o := []obj{1, "23"}
	res := mapFunc(f1, o)
	fmt.Println(res)
}

func mapFunc(f mf, o []obj) []obj {
	res := make([]obj, len(o))
	for k, v := range o {
		res[k] = f(v)
	}
	return res
}

func f1(o obj) obj {
	var res obj
	switch o.(type) {
	case string:
		res = o.(string) + o.(string)
	case int:
		res = o.(int) * 2
	}
	return res
}
