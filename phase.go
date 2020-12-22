package strife

import (
	"encoding/json"
	"io/ioutil"
)

type phase struct {
	Name                      string  `json:"name"`
	Number                    uint16  `json:"number"`
	AvarageNumberOfHoursSpent float64 `json:"average_number_of_hours_spent"`
}

type phases []phase

type recruting struct {
	Phases *phases `json:"phases"`
}

func (p *phases) Rates() *rates {
	var rs rates
	for i, ph := range *p {
		if i == 0 {
			continue
		}
		rate, _ := newRateWithPhase(&((*p)[i-1]), &ph)
		rs = append(rs, rate)
	}
	return &rs
}

func readJSON(jsonPath string) (*recruting, error) {
	bytes, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		return nil, err
	}

	var recruting recruting
	if err := json.Unmarshal(bytes, &recruting); err != nil {
		return nil, err
	}
	return &recruting, nil
}
