package main

import (
	"fmt"
	"time"
)

func main() {
	var time1 = time.Now()
	fmt.Printf("time1 %v\n", time1)
	// time1 2015-09-01 17:59:31.73600891 +0700 WIB

	var time2 = time.Date(2011, 12, 24, 10, 20, 0, 0, time.UTC)
	fmt.Printf("time2 %v\n", time2)
	// time2 2011-12-24 10:20:00 +0000 UTC

	// time.Date(tahun, bulan, tanggal, jam, menit, detik, nanodetik, timezone)

	var now = time.Now()
	fmt.Println("year:", now.Year(), "month:", now.Month())
	// year: 2015 month: 8

	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-01-02 15:04:05"
	value = "2015-09-02 08:04:00"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date.String())
	// 2015-09-02 08:04:00 +0000 UTC

	layoutFormat = "02/01/2006 MST"
	value = "02/09/2015 WIB"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t\t->", date.String())
	// 2015-09-02 00:00:00 +0700 WIB

	var date2, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")
	fmt.Println(date2.String())
	// 2015-09-02 08:00:00 +0700 WIB

	var date3, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")

	var dateS1 = date3.Format("Monday 02, January 2006 15:04 MST")
	fmt.Println("dateS1", dateS1)
	// Wednesday 02, September 2015 08:00 WIB

	var dateS2 = date3.Format(time.RFC3339)
	fmt.Println("dateS2", dateS2)
	// 2015-09-02T08:00:00+07:00

	var date5, err = time.Parse("06 Jan 15", "02 Sep 15 08:00 WIB")

	if err != nil {
		fmt.Println("error", err.Error())
		return
	}

	fmt.Println(date5)
}
