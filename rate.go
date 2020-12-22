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

type rates []*rate

func (r *rates) predict(num uint16) (*phases, error) {
	rLen := len(*r)
	ps := make(phases, rLen+1)
	ps[rLen] = phase{Name: (*r)[rLen-1].PhaseAfter, Number: num}

	for i := rLen - 1; i >= 0; i-- {
		rate := (*r)[i]
		den := rate.Rate.Denom()
		nume := rate.Rate.Num()
		if nume.Int64() == 0 {
			return nil, errors.New("nume MUST be over 0")
		}
		tempNumber := ps[i+1].Number * uint16(den.Int64())
		ceiled := math.Ceil(float64(tempNumber) / float64(nume.Int64()))
		ps[i] = phase{Name: rate.PhaseBefore, Number: uint16(ceiled)}
	}
	return &ps, nil
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
