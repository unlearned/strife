package strife

import "testing"

func createPredictions() *Predictions {
	return &Predictions{
		prediction{phase{"first", 100, 1.5}, 150.0},
		prediction{phase{"second", 50, 2.0}, 100.0},
		prediction{phase{"third", 25, 3.0}, 75.0},
	}
}

func TestTextSuccessOnPredictions(t *testing.T) {
	ps := createPredictions()
	text := ps.text()
	idealText := "|         phase        | first|second| third|\n"
	idealText += "|         number       |  100 |  50  |  25  |\n"
	idealText += "|average hours required|  1.5 |   2  |   3  |\n"
	idealText += "| total hours required |  150 |  100 |  75  |"

	if text != idealText {
		t.Errorf("it should return \n%v\n but it returned \n%v", idealText, text)
	}
}

func TestLongestLengthOnPredictions(t *testing.T) {
	ps := createPredictions()
	if ps.longestLength() != 6 {
		t.Errorf("it should return 6 but it returned %v", ps.longestLength())
	}
}

func TestPadHoursRequired(t *testing.T) {
	p := &prediction{phase{"first", 100, 1.5}, 150.0}
	str := p.padHoursRequired(7)
	if str != "  150  " {
		t.Errorf("it should return %v but it returned %v", "  150  ", str)
	}
}
