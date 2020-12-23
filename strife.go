package strife

import "unicode/utf8"

// func Predict(desiredNum uint16, phases Phases) (*Phases, error) {
// 	rs := phases.rates()
// 	ps, err := rs.predict(desiredNum)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return ps, nil
// }

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
