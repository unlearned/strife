package strife

import (
	"math/big"
	"reflect"
	"testing"
)

func createPhases() *Phases {
	return &Phases{
		phase{"first", 100, 1.5},
		phase{"second", 50, 2.0},
		phase{"third", 25, 3.0},
	}
}

func TestRatesSuccess(t *testing.T) {
	ps := createPhases()
	rt := ps.rates()

	if reflect.TypeOf(rt).String() != "*strife.rates" {
		t.Errorf("it should return *strife.rates but %v", reflect.TypeOf(rt).String())
	}

	if len(rt.rates) != 2 {
		t.Errorf("it should return length 2 of slice but it returned %v", len(rt.rates))
	}

	if rt.rates[0].PhaseBefore != "first" {
		t.Errorf("it should return first but it returned %v", rt.rates[0].PhaseBefore)
	}

	if rt.rates[0].PhaseAfter != "second" {
		t.Errorf("it should return second but it returned %v", rt.rates[0].PhaseAfter)
	}

	if rt.rates[0].Rate.String() != big.NewRat(50, 100).String() {
		t.Errorf("it should return 1/2 but it returned %v", rt.rates[0].Rate)
	}

	if rt.rates[1].PhaseBefore != "second" {
		t.Errorf("it should return first but it returned %v", rt.rates[1].PhaseBefore)
	}

	if rt.rates[1].PhaseAfter != "third" {
		t.Errorf("it should return third but it returned %v", rt.rates[1].PhaseAfter)
	}

	if rt.rates[1].Rate.String() != big.NewRat(25, 50).String() {
		t.Errorf("it should return 1/2 but it returned %v", rt.rates[1].Rate)
	}
}

func TestLongestLengthOnPhases(t *testing.T) {
	ps := createPhases()
	if ps.longestLength() != 6 {
		t.Errorf("it should return 6 but it returned %v", ps.longestLength())
	}
}

func TestTextSuccessOnPhases(t *testing.T) {
	ps := createPhases()
	text := ps.text()
	idealText := "|      phase        | first|second| third|\n"
	idealText += "|      number       |  100 |  50  |  25  |\n"
	idealText += "|average hours spent|  1.5 |   2  |   3  |"
	if text != idealText {
		t.Errorf("it should return \n%v\n but it returned \n%v", idealText, text)
	}
}

func TestPadName(t *testing.T) {
	p := &phase{"first", 100, 1.5}
	str := p.padName(7)
	if str != " first " {
		t.Errorf("it should return %v but it returned %v", " first ", str)
	}
}

func TestPadNumber(t *testing.T) {
	p := &phase{"first", 100, 1.5}
	str := p.padNumber(8)
	if str != "   100  " {
		t.Errorf("it should return %v but it returned %v", "   100  ", str)
	}
}

func TestAvarageNumberOfHoursSpent(t *testing.T) {
	p := &phase{"first", 100, 1.5}
	str := p.padAvarageNumberOfHoursSpent(7)
	if str != "  1.5  " {
		t.Errorf("it should return %v but it returned %v", "  1.5  ", str)
	}
}
