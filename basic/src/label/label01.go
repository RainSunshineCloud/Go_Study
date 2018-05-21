package main

import "fmt"

/*
使用continue/break label 其中label 必须提前定义
使用goto label 其中label不需要提前定义
只有通过以上三种方式label 内代码才会执行
*/
func main() {
	i := 1
	if i == 2 {
		goto LABEL1
	}

	goto LABEL2

LABEL1:
	for ; i < 9; i++ {
		if i == 2 {
			continue LABEL1
		}
		fmt.Println(i)
	}
LABEL2:
	if i < 4 {
		i++
		goto LABEL2
		fmt.Println("走LABEL2")
	} else {
		goto LABEL1
		fmt.Println("走LABEL1")
	}

}
