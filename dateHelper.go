package main

import (
	"fmt"
	"strconv"
	"strings"
)

var Month = []string{
	"",
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

type Date struct {
	year  int
	month int
	day   int
}

func (d *Date) String() string {
	return fmt.Sprintf("%d-%d-%d", d.year, d.month, d.day)
}

func (d *Date) GetYear() int {
	return d.year
}

func (d *Date) GetMonth() int {
	return d.month
}

func (d *Date) GetDay() int {
	return d.day
}

func (d *Date) GetMonthString() string {
	return Month[d.month]
}

func (d *Date) IsLeapYear() bool {
	return (d.year%4 == 0 || (d.year%400 == 0 && d.year%100 != 0))
}

func (d *Date) DayInYear() int {
	var days = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if d.IsLeapYear() {
		days[1] = 29
	}

	var iRet = d.day
	for i := 0; i < d.month-1; i++ {
		iRet += days[i]
	}
	return iRet
}

func DaysDiff(d *Date, other *Date) int {
	if Comapre(d, other) < 0 {
		return DaysDiff(other, d)
	}

	d1 := d.DayInYear()
	d2 := other.DayInYear()

	diff := (d.year-other.year)*365 + d1 - d2 + (d.year / 4) - (other.year / 4) - (d.year / 100) + (other.year / 100)
	return diff
}

func Comapre(d *Date, other *Date) int {
	if d.year > other.year {
		return 1
	} else if d.year < other.year {
		return -1
	}

	if d.month > other.month {
		return 1
	} else if d.month < other.month {
		return -1
	}

	if d.day > other.day {
		return 1
	} else if d.day < other.day {
		return -1
	}

	return 0
}

func ChangeStringToDate(date string) (Date, error) {
	s := strings.Split(date, "-")

	y, err := strconv.Atoi(s[0])
	if err != nil {
		return Date{0, 0, 0}, err
	}

	m, err := strconv.Atoi(s[1])
	if err != nil {
		return Date{0, 0, 0}, err
	}

	d, err := strconv.Atoi(s[2])
	if err != nil {
		return Date{0, 0, 0}, err
	}

	return Date{y, m, d}, nil
}

func main() {
	d1 := Date{2018, 10, 15}
	d2 := Date{2018, 10, 10}

	days := DaysDiff(&d1, &d2)
	fmt.Println(days)

	dd, _ := ChangeStringToDate("2018-12-12")
	fmt.Println(dd)

	days = DaysDiff(&d1, &dd)
	fmt.Println(days)

	fmt.Println(dd.GetMonthString())

	d3 := Date{2018,1,1}
	d4 := Date{2018,1,1}
	fmt.Println(DaysDiff(&d3, &d4))
}

