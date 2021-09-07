package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	Year int
	Month int
	Day int
}

func (d *Date) SetYear(year int) error {
	if year <= 0 {
		return errors.New("invalid year")
	}
	d.Year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month <= 0 || month > 12 {
		return errors.New("invalid month")
	}
	d.Month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day <= 0 {
		return errors.New("invalid day")
	}
	d.Day = day
	return nil
}
func main() {
	date := Date{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(9)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(7)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
}