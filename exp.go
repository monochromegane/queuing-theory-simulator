package queuing_theory_simulator

import (
	"math/rand"
)

type Exper interface {
	Exp(int) float64
}

func NewChangeExp(seed int64, params func(int) float64) *ChangeExp {
	return &ChangeExp{
		rand:   rand.New(rand.NewSource(seed)),
		params: params,
	}
}

type ChangeExp struct {
	params func(int) float64
	rand   *rand.Rand
}

func (e *ChangeExp) Exp(i int) float64 {
	return e.rand.ExpFloat64() / e.params(i)
}
