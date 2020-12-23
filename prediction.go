package strife

import (
	"fmt"
	"unicode/utf8"
)

type prediction struct {
	phase
	hoursRequired float64
}

type Predictions []prediction

func (ps *Predictions) text() string {
	var l1 string
	var l2 string
	var l3 string
	var l4 string
	longestLen := ps.longestLength()

	for _, p := range *ps {
		l1 += fmt.Sprintf("%v|", p.padName(longestLen))
		l2 += fmt.Sprintf("%v|", p.padNumber(longestLen))
		l3 += fmt.Sprintf("%v|", p.padAvarageNumberOfHoursSpent(longestLen))
		l4 += fmt.Sprintf("%v|", p.padHoursRequired(longestLen))
	}

	format := "|         phase        |%v\n"
	format += "|         number       |%v\n"
	format += "|average hours required|%v\n"
	format += "| total hours required |%v"

	return fmt.Sprintf(format, l1, l2, l3, l4)
}

func (p *prediction) padHoursRequired(num uint32) string {
	return pad(fmt.Sprintf("%v", p.hoursRequired), num)
}

func (ps *Predictions) longestLength() uint32 {
	length := 0
	for _, p := range *ps {
		nameLen := utf8.RuneCountInString(p.Name)
		numberLen := utf8.RuneCountInString(fmt.Sprint(uint64(p.Number)))
		hoursLen := utf8.RuneCountInString(fmt.Sprint(p.AverageNumberOfHoursSpent))
		reqLen := utf8.RuneCountInString(fmt.Sprint(p.hoursRequired))
		tempLen := 0

		if nameLen > numberLen {
			tempLen = nameLen
		} else {
			tempLen = numberLen
		}

		if hoursLen > tempLen {
			tempLen = hoursLen
		}

		if reqLen > tempLen {
			tempLen = reqLen
		}

		if tempLen > length {
			length = tempLen
		}
	}
	return uint32(length)
}
