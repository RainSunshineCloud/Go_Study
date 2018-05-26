package main

import (
	"time"
	"fmt"
)

func main () {
	times := time.Now();
	str := times.UTC();
	nano := times.Nanosecond();
	year,month,day := times.Date();
	timeAfter := time.Now();
	bool1 := times.After(timeAfter);
	bool2 := times.Before(timeAfter);
	fmt.Println(times);
	fmt.Printf("%v \n",str);
	fmt.Println(nano);
	fmt.Println(year,month,day);
	fmt.Println(bool1,bool2);
}
