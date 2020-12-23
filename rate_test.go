package strife

import (
	"math/big"
	"reflect"
	"testing"
)

func TestPredictSuccess(t *testing.T) {
	r := &rates{
		rates: []*rate{
			&rate{PhaseBefore: "First", PhaseAfter: "Second", Rate: big.NewRat(2, 5)},
			&rate{PhaseBefore: "Second", PhaseAfter: "Third", Rate: big.NewRat(1, 3)},
			&rate{PhaseBefore: "Third", PhaseAfter: "Fourth", Rate: big.NewRat(1, 4)},
		},
		phases: &Phases{
			phase{Name: "First", Number: 3600, AverageNumberOfHoursSpent: 1.1},
			phase{Name: "Second", Number: 1440, AverageNumberOfHoursSpent: 1.2},
			phase{Name: "Third", Number: 480, AverageNumberOfHoursSpent: 1.3},
			phase{Name: "Fourth", Number: 120, AverageNumberOfHoursSpent: 1.4},
		},
	}

	ps, err := r.predict(100)
	if err != nil {
		t.Errorf("it should return err as nil but it returned err as %v", err)
	}

	if reflect.TypeOf(ps).String() != "*strife.Predictions" {
		t.Fatalf("it should return *strife.Phases but %v", reflect.TypeOf(ps).String())
	}

	if len(*ps) != 4 {
		t.Fatalf("it should return 4 length but %v length", len(*ps))
	}

	p1 := prediction{phase: phase{"First", 3000, 1.1}, hoursRequired: 3300}
	if (*ps)[0] != p1 {
		t.Errorf("it should return {{First 3000 1.1} 3300} but it returned %v", (*ps)[0])
	}

	p2 := prediction{phase: phase{"Second", 1200, 1.2}, hoursRequired: 1440}
	if (*ps)[1] != p2 {
		t.Errorf("it should return  {{Second 1200 1.2} 1440} but it returned %v", (*ps)[1])
	}

	p3 := prediction{phase: phase{"Third", 400, 1.3}, hoursRequired: 520}
	if (*ps)[2] != p3 {
		t.Errorf("it should return {{Third 400 1.3} 520} but it returned %v", (*ps)[2])
	}

	p4 := prediction{phase: phase{"Fourth", 100, 1.4}, hoursRequired: 140}
	if (*ps)[3] != p4 {
		t.Errorf("it should return {{Fourth 100 1.4} 140} but it returned %v", (*ps)[3])
	}
}

func TestPredictFailed(t *testing.T) {
	r := &rates{
		rates: []*rate{
			&rate{PhaseBefore: "First", PhaseAfter: "Second", Rate: big.NewRat(0, 5)},
		},
		phases: &Phases{
			phase{Name: "First", Number: 0},
			phase{Name: "Second", Number: 0},
		},
	}

	ps, err := r.predict(100)

	if ps != nil {
		t.Errorf("it should return ps as nil but it returned ps as %v", ps)
	}

	if err == nil {
		t.Error("it should return err but it returned err as nil")
	}
}

func TestNewRateWithPhaseSuccess(t *testing.T) {
	before := &phase{"before", 100, 1.5}
	after := &phase{"after", 50, 1.5}
	rate, err := newRateWithPhase(before, after)

	if err != nil {
		t.Errorf("it should return err as nil but it returned err as %v", err)
	}

	if reflect.TypeOf(rate).String() != "*strife.rate" {
		t.Errorf("it should return *strife.rate but %v", reflect.TypeOf(rate).String())
	}

	if rate.PhaseBefore != before.Name {
		t.Errorf("it should return %v but it returned %v", before.Name, rate.PhaseBefore)
	}

	if rate.PhaseAfter != after.Name {
		t.Errorf("it should return %v but it returned %v", after.Name, rate.PhaseAfter)
	}

	if rate.Rate.String() != big.NewRat(1, 2).String() {
		t.Errorf("it should return 1/2 but it returned %v", rate.Rate.String())
	}
}

func TestNewRateWithPhaseFailed(t *testing.T) {
	before := &phase{"before", 50, 1.5}
	after := &phase{"after", 100, 1.5}
	rate, err := newRateWithPhase(before, after)

	if rate != nil {
		t.Errorf("it should return rate as nil but it returned rate as %v", rate)
	}

	if err == nil {
		t.Errorf("it should return err")
	}
}

func TestPassRateFailedNumberOfInterviewsBeforeIs0(t *testing.T) {
	result, err := passRate(0, 1)
	if err == nil {
		t.Fatal("failed test")
	}

	if result != nil {
		t.Fatal("failed test")
	}
}

func TestPassRateFailedNumberOfInterviewsAfterIsOverNumberOfInterviewsBefore(t *testing.T) {
	res, err := passRate(1, 3)
	if err == nil {
		t.Errorf("it should return err as nil but it returned err as %v", err)
	}

	if res != nil {
		t.Errorf("it should return res as nil but it returned res as %v", res)
	}
}

func TestPassRateSuccess(t *testing.T) {
	result, err := passRate(10, 2)
	if err != nil {
		t.Errorf("it should return err as nil but it returned %v", err)
	}

	if result.String() != big.NewRat(2, 10).String() {
		t.Errorf("it should return 1/5 but it returned %v", result)
	}
}

func TestFractionalDivideSuccess(t *testing.T) {
	r := big.NewRat(2, 4)
	n, err := fractionalDivide(r, 5)

	if err != nil {
		t.Errorf("it should return err as nil but it returned %v", err)
	}

	if n != 10.0 {
		t.Errorf("it should return 10 but it returned %v", n)
	}
}

func TestFractionalDivideFailed(t *testing.T) {
	r := big.NewRat(0, 4)
	n, err := fractionalDivide(r, 5)

	if err == nil {
		t.Error("it should return err but it returned err as nil")
	}

	if n != 0.0 {
		t.Errorf("it should return 0.0 but it returned %v", n)
	}
}
