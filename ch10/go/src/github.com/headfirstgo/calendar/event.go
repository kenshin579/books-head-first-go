package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string
	Date  //Date (anonymous field)는 여기에 embedded가 됨 (outer struct, Event에 inner struct로 인식됨)
}

func (e *Event) Title() string {
	return e.title
}
func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}
