package calendar

import "errors"

type Date struct {
	year int
	month int
	day int
}

func (d *Date) SetYear(year int) error {
	if year <= 0 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month <= 0 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day <= 0 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}