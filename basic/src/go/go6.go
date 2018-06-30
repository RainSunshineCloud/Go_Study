package main

import (
	"fmt"
	"math"
)

func main () {
	channel1 := make( chan [2]float64);
	channel2 := make( chan [2]float64);
	go getPolar(channel1);
	go getQuader(channel1,channel2);
	res := <-channel2;
	fmt.Printf("x轴%f,y轴%f",res[0],res[1]);
}

func getPolar(channel1 chan [2]float64) {

	var polar [2] float64;
	fmt.Println("请输入角度和值（用,分隔）");
	fmt.Scanf("%f,%f",&polar[0],&polar[1]);
	channel1 <- polar;
}

func getQuader(channel1 chan [2]float64,channel2 chan [2]float64) {
	polar := <-channel1;
	var Quader [2]float64;

	Quader[0] = math.Cos(polar[0]) * polar[1];
	Quader[1] = math.Sin(polar[0]) * polar[1];

	channel2 <- Quader;
}

