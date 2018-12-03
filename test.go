package main

import (
	"fmt"
	"math"
	"time"
)

const DATE_FORMAT = "2006-01-02"

func main() {
	year, month, _ := time.Now().Date()

	// 上个月
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	start := thisMonth.AddDate(0, -1, 0).Format(DATE_FORMAT)
	end := thisMonth.AddDate(0, 0, -1).Format(DATE_FORMAT)
	fmt.Println(fmt.Sprintf("%s~%s", start, end))

	now := time.Now()
	if now.Weekday() == time.Monday {
		s3 := now.AddDate(0, 0, -7).Format(DATE_FORMAT)
		e3 := now.AddDate(0, 0, -1).Format(DATE_FORMAT)
		fmt.Println(fmt.Sprintf("上个星期: %s~%s", s3, e3))
	}

	t1, _ := time.ParseInLocation(DATE_FORMAT, "2018-11-19", time.Local)
	t2, _ := time.ParseInLocation(DATE_FORMAT, "2018-11-20", time.Local)
	fmt.Println(math.Abs(t1.Sub(t2).Hours() / 24))

	t := time.Date(year-1, 1, 1, 0, 0, 0, 0, time.Local)
	s4 := t.Format(DATE_FORMAT)
	e4 := t.AddDate(0, 12, -1).Format(DATE_FORMAT)
	fmt.Println(s4, e4)

	tt, _ := time.ParseInLocation(DATE_FORMAT, "2018-01-19", time.Local)
	fmt.Println(tt)


}