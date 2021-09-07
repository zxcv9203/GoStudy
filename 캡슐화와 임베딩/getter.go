package main

import (
	"fmt"
	"github.com/GoStudy/캡슐화와 임베딩/calendar"
	"log"
)

func main() {
	date := calendar.Date{}
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
	fmt.Println(date.Year)
	fmt.Println(date.Month)
	fmt.Prinltn(date.Day)
}