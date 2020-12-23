package strife

import (
	"errors"
	"math"
	"math/big"
)

type rate struct {
	PhaseBefore string
	PhaseAfter  string
	Rate        *big.Rat
}

type rates struct {
	rates  []*rate
	phases *Phases
}

func (r *rates) predict(num uint16) (*Predictions, error) {
	phs := r.phases
	endIndex := len(*phs) - 1
	ps := make(Predictions, endIndex+1)

	ps[endIndex] = prediction{
		phase:         phase{Name: (*phs)[endIndex].Name, Number: num},
		hoursRequired: float64(num) * (*phs)[endIndex].AverageNumberOfHoursSpent,
	}

	rs := r.rates
	for i := endIndex - 1; i >= 0; i-- {
		rate := rs[i]
		a, err := fractionalDivide(rate.Rate, ps[i+1].Number)
		if err != nil {
			return nil, err
		}
		n := math.Ceil(a)
		ps[i] = prediction{
			phase:         phase{Name: rate.PhaseBefore, Number: uint16(n)},
			hoursRequired: float64(n) * (*phs)[i+1].AverageNumberOfHoursSpent,
		}
	}
	return &ps, nil
}

func fractionalDivide(r *big.Rat, n uint16) (float64, error) {
	den := r.Denom()
	nume := r.Num()
	if nume.Int64() == 0 {
		return 0.0, errors.New("nume MUST be over 0")
	}
	temp := n * uint16(den.Int64())
	return float64(temp) / float64(nume.Int64()), nil
}

func newRateWithPhase(before, after *phase) (*rate, error) {
	rating, err := passRate(before.Number, after.Number)
	if err != nil {
		return nil, err
	}
	return &rate{
		PhaseBefore: before.Name,
		PhaseAfter:  after.Name,
		Rate:        rating,
	}, nil
}

func passRate(numBefore, numAfter uint16) (*big.Rat, error) {
	if numBefore == 0 {
		return nil, errors.New("numBefore MUST be over 0")
	}

	if numAfter > numBefore {
		return nil, errors.New("numBefore MUST be greater than numAfter")
	}
	rate := big.NewRat(int64(numAfter), int64(numBefore))
	return rate, nil
}
