package strife

import (
	"fmt"
	"unicode/utf8"
)

type phase struct {
	Name                      string  `json:"name"`
	Number                    uint16  `json:"number"`
	AverageNumberOfHoursSpent float64 `json:"average_number_of_hours_spent"`
}

type Phases []phase

func (p *Phases) rates() *rates {
	var rs []*rate
	for i, ph := range *p {
		if i == 0 {
			continue
		}
		rate, _ := newRateWithPhase(&((*p)[i-1]), &ph)
		rs = append(rs, rate)
	}
	return &rates{
		rates:  rs,
		phases: p,
	}
}

func (p *phase) padName(num uint32) string {
	return pad(p.Name, num)
}

func (p *phase) padNumber(num uint32) string {
	return pad(fmt.Sprintf("%v", p.Number), num)
}

func (p *phase) padAvarageNumberOfHoursSpent(num uint32) string {
	return pad(fmt.Sprintf("%v", p.AverageNumberOfHoursSpent), num)
}

func (ps *Phases) text() string {
	var l1 string
	var l2 string
	var l3 string
	longestLen := ps.longestLength()

	for _, p := range *ps {
		l1 += fmt.Sprintf("%v|", p.padName(longestLen))
		l2 += fmt.Sprintf("%v|", p.padNumber(longestLen))
		l3 += fmt.Sprintf("%v|", p.padAvarageNumberOfHoursSpent(longestLen))
	}
	return fmt.Sprintf("|      phase        |%v\n|      number       |%v\n|average hours spent|%v", l1, l2, l3)
}

func (ps *Phases) longestLength() uint32 {
	length := 0
	for _, p := range *ps {
		nameLen := utf8.RuneCountInString(p.Name)
		numberLen := utf8.RuneCountInString(fmt.Sprint(uint64(p.Number)))
		hoursLen := utf8.RuneCountInString(fmt.Sprint(p.AverageNumberOfHoursSpent))
		tempLen := 0

		if nameLen > numberLen {
			tempLen = nameLen
		} else {
			tempLen = numberLen
		}

		if hoursLen > tempLen {
			tempLen = hoursLen
		}

		if tempLen > length {
			length = tempLen
		}
	}
	return uint32(length)
}
