package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	month := now.Month()         //5月
	year := now.Year()           //2018
	yearday := now.YearDay()     //141天
	day := now.Day()             //21号
	weekday := now.Weekday()     //星期几
	year1, day1 := now.ISOWeek() //2018年的第21周
	format := now.Format("15")

	thisLocation := time.Local
	thisMonth := time.Date(2018, 5, 1, 12, 0, 0, 0, thisLocation)
	nextTime := thisMonth.Add(1e15)
	tick := time.Tick(1e11)
	time.Sleep(1e11) //睡眠按微纳秒计算
	fmt.Println(now, month, year, yearday, day, weekday, year1, day1, format, thisMonth, nextTime, tick)

}
