package poisson

import (
	"math/rand"
)

//RandomGenerator is interface for any random generator that can be used with poisson disk sampling
type RandomGenerator interface {
	Float() float64
	Integer(max int) int
}

type basicGenerator struct{}

//NewBasicGenerator returns a basic random generator
func NewBasicGenerator(seed int) RandomGenerator {
	rand.Seed(int64(seed))
	return basicGenerator{}
}

func (g basicGenerator) Float() float64 {
	return rand.Float64()
}

func (g basicGenerator) Integer(max int) int {
	return rand.Int() * max
}
