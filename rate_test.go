package strife

import (
	"math/big"
	"reflect"
	"testing"
)

func TestPredictSuccess(t *testing.T) {
	r := &rates{
		&rate{PhaseBefore: "First", PhaseAfter: "Second", Rate: big.NewRat(2, 5)},
		&rate{PhaseBefore: "Second", PhaseAfter: "Third", Rate: big.NewRat(1, 3)},
		&rate{PhaseBefore: "Third", PhaseAfter: "Fourth", Rate: big.NewRat(1, 4)},
	}

	ps, err := r.predict(100)
	if err != nil {
		t.Errorf("it should return err as nil but it returned err as %v", err)
	}

	if reflect.TypeOf(ps).String() != "*strife.phases" {
		t.Fatalf("it should return *strife.phases but %v", reflect.TypeOf(ps).String())
	}

	if len(*ps) != 4 {
		t.Fatalf("it should return 4 length but %v length", len(*ps))
	}

	p1 := phase{"First", 3000, 0}
	if (*ps)[0] != p1 {
		t.Errorf("it should return {First, 3000, 0} but it returned %v", (*ps)[0])
	}

	p2 := phase{"Second", 1200, 0}
	if (*ps)[1] != p2 {
		t.Errorf("it should return {Second, 1200, 0} but it returned %v", (*ps)[1])
	}

	p3 := phase{"Third", 400, 0}
	if (*ps)[2] != p3 {
		t.Errorf("it should return {Third, 400, 0} but it returned %v", (*ps)[2])
	}

	p4 := phase{"Fourth", 100, 0}
	if (*ps)[3] != p4 {
		t.Errorf("it should return {Fourth, 100, 0} but it returned %v", (*ps)[3])
	}
}

func TestPredictFailed(t *testing.T) {
	r := &rates{
		&rate{PhaseBefore: "First", PhaseAfter: "Second", Rate: big.NewRat(0, 5)},
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
