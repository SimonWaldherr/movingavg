package movingavg

import (
	"math"
	"sort"
)

type movingAvgFloat struct {
	window int
	values []float64
}

func NewFloat(window int) *movingAvgFloat {
	return &movingAvgFloat{
		window: window,
		values: make([]float64, 0),
	}
}

func (ma *movingAvgFloat) Values() []float64 {
	return ma.values
}

func (ma *movingAvgFloat) Add(value float64) {
	if len(ma.values) == ma.window {
		ma.values = append((ma.values)[1:ma.window], value)
	} else {
		ma.values = append(ma.values, value)
	}
}

func (ma *movingAvgFloat) Min() float64 {
	min := ma.values[0]
	for _, value := range ma.values {
		if value < min {
			min = value
		}
	}
	return min
}

func (ma *movingAvgFloat) Max() float64 {
	max := ma.values[0]
	for _, value := range ma.values {
		if value > max {
			max = value
		}
	}
	return max
}

func (ma *movingAvgFloat) Sum() float64 {
	var sum float64
	values := ma.values
	if values == nil {
		return 0
	}

	for _, value := range values {
		sum += value
	}
	return sum
}

func (ma *movingAvgFloat) Median() float64 {
	values := ma.values
	c := len(values)

	if c == 0 {
		return 0
	}

	sort.Float64s(values)
	if c%2 == 1 {
		return values[c/2]
	}
	return (values[c/2] + values[c/2-1]) / 2
}

func (ma *movingAvgFloat) Arithmetic() float64 {
	values := ma.values
	c := float64(len(values))

	if c == 0 {
		return 0
	}

	sum := 0.00

	for _, value := range ma.values {
		sum += value
	}
	return sum / c
}

func (ma *movingAvgFloat) Geometric() float64 {
	values := ma.values
	c := len(values)

	var m float64 = 1
	for _, value := range values {
		m = m * float64(value)
	}
	return math.Pow(m, 1/float64(c))
}
