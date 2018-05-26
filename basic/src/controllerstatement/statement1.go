package main

import "fmt"

func main() {
	//ifs();
	//switchs()
	fors()

}

func ifs() {

	i := 1
	if i < 10 {
		fmt.Println(i)
	} else if i < 5 {
		fmt.Println(i)
	}
	//报错，condition 只能是bool类型的。php可以是因为隐型转换
	//if i {
	//	fmt.Println(i)
	//}

	if x := 0; x < 10 {
		fmt.Println(2)
	}
	//不行的，只能包含一条初始化语句
	//if x := 0; y := 1; x < y {
	//	fmt.Println(y)
	//}
}

func switchs() {
	i := 1

	/*
		switch 不能使用continue (这个和php 不一样）
		switch php 默认向下执行（c也一样） go 默认break
	*/
	switch i {

	case 1:
		i = 2
		fallthrough
	case 2:
		i = 3
	case 3:
		i = 4
	}

	/*
		不行的，只能包含一条初始化语句 或者 一个值 或者不包含
	*/
	//switch b := 1; i {
	//case 1:
	//	i = 2
	//	fallthrough
	//case 2:
	//	i = 3
	//case 3:
	//	i = 4
	//}

	/*
		以下两条语句等价
	*/
	switch true {
	case 1 == 1:
		i = 6
	case 1 == 3:
		i = 7
	}

	switch {
	case 1 == 1:
		i = 6
	case 1 == 3:
		i = 7
	}

	fmt.Println(i)
}

func fors() {
	//用法1（和php的for 用法一样）
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
	//用法2 （记住作用域的影响）相当于dowhile
	i := 1
	for {
		fmt.Println(i)
		i++
		if i >= 0 {
			break
		}

	}
	//用法3相当于while
	for i < 9 {
		fmt.Println(i)
		i++
	}

	//用法4 和range 配合 相当于 foreach
	slices := [5]int{1, 2, 3, 4, 5}
	for key, val := range slices {
		fmt.Printf("键是%v,值是%v\n", key, val)
	}

	for key := range slices {
		fmt.Printf("键是%v\n", key)
	}
	for _, val := range slices {
		fmt.Printf("值是%v\n", val)
	}

	for _, c := range "sd" {
		fmt.Printf("值是%v\n", c)
	}

	//使用指针的修改值的方法
	for key, val := range slices {
		fmt.Printf("指针是%v\n", &slices[key])
		fmt.Printf("value临时指针是%v\n", &val)
	}
}
