package strife

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"unicode/utf8"
)

func PredictText(n uint16, jsonPath string) (string, error) {
	rec, err := readJSON(jsonPath)
	if err != nil {
		return "", err
	}

	rs := rec.Phases.rates()
	ps, err := rs.predict(n)
	if err != nil {
		return "", err
	}

	text := "## Past Performance ##\n"
	text += rs.phases.text() + "\n"
	text += "\n"
	text += "## Prediction ##\n"
	text += ps.text() + "\n"

	return text, nil
}

func readJSON(jsonPath string) (*Recruting, error) {
	bytes, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		return nil, err
	}

	var recruting Recruting
	if err := json.Unmarshal(bytes, &recruting); err != nil {
		return nil, err
	}
	return &recruting, nil
}

func pad(s string, num uint32) string {
	len := uint32(utf8.RuneCountInString(s))
	if len > num {
		return s
	}
	padNum := num - len

	str := s
	for i := padNum / 2; i > 0; i-- {
		str = " " + str + " "
	}

	if padNum%2 != 0 {
		str = " " + str
	}
	return str
}

func round2(num float64) float64 {
	output := math.Pow(10, float64(2))
	n := num * output
	i := int(n + math.Copysign(0.5, n))
	return float64(i) / output
}
