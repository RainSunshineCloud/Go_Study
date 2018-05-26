package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var cl cal
	cl.getInput()
	res, _ := cl.operate()
	fmt.Println(res)
}

type cal struct {
	first    float32
	operator string
	second   float32
}

func (cl *cal) getInput() {
	reader := bufio.NewReader(os.Stdin)
	var string string
	var err error

	if string, err = reader.ReadString('\n'); err != nil {
		panic(err)
	}

	if _, err = fmt.Sscanf(string, "%f %s %f", &cl.first, &cl.operator, &cl.second); err != nil {
		panic(err)
	}
}

type error interface{}

func (cl *cal) operate() (float32, error) {
	switch cl.operator {
	case "*":
		return cl.first * cl.second, nil
	case "-":
		return cl.first - cl.second, nil
	case "/":
		return cl.first / cl.second, nil
	case "+":
		return cl.first + cl.second, nil
	default:
		return 0, "请输入加减乘除操作符"
	}
}
