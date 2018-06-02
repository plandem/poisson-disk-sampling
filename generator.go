package poisson

import (
	"math/rand"
)

type Generator interface {
	RandomFloat() float64
	RandomInt(max int) int
}

type basicGenerator struct { }

func NewBasicGenerator(seed int) (Generator) {
	rand.Seed(int64(seed))
	return basicGenerator{}
}

func (g basicGenerator) RandomFloat() (float64) {
	return rand.Float64()
}

func (g basicGenerator) RandomInt(max int) (int) {
	return rand.Int() * max
}