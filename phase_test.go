package strife

import (
	"math/big"
	"reflect"
	"testing"
)

func createPhases() *phases {
	return &phases{
		phase{"first", 100, 1.0},
		phase{"second", 50, 2.0},
		phase{"third", 25, 3.0},
	}
}

func TestRatesSuccess(t *testing.T) {
	ps := createPhases()
	rt := ps.Rates()

	if reflect.TypeOf(rt).String() != "*strife.rates" {
		t.Errorf("it should return *strife.rates but %v", reflect.TypeOf(rt).String())
	}

	if len(*rt) != 2 {
		t.Errorf("it should return length 2 of slice but it returned %v", len(*rt))
	}

	if (*rt)[0].PhaseBefore != "first" {
		t.Errorf("it should return first but it returned %v", (*rt)[0].PhaseBefore)
	}

	if (*rt)[0].PhaseAfter != "second" {
		t.Errorf("it should return second but it returned %v", (*rt)[0].PhaseAfter)
	}

	if (*rt)[0].Rate.String() != big.NewRat(50, 100).String() {
		t.Errorf("it should return 1/2 but it returned %v", (*rt)[0].Rate)
	}

	if (*rt)[1].PhaseBefore != "second" {
		t.Errorf("it should return first but it returned %v", (*rt)[1].PhaseBefore)
	}

	if (*rt)[1].PhaseAfter != "third" {
		t.Errorf("it should return third but it returned %v", (*rt)[0].PhaseAfter)
	}

	if (*rt)[1].Rate.String() != big.NewRat(25, 50).String() {
		t.Errorf("it should return 1/2 but it returned %v", (*rt)[1].Rate)
	}
}

func TestReadJSONSuccess(t *testing.T) {
	p, err := readJSON("./testdata/interview.json")

	if err != nil {
		t.Errorf("it should return err as nil but it returned %v", err)
	}

	if reflect.TypeOf(*p).String() != "strife.recruting" {
		t.Errorf("it should return strife.recruting but it returned %v", reflect.TypeOf(*p))
		t.Fatal(err)
	}
}
