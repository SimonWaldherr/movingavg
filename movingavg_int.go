package movingavg

import (
	"math"
	"sort"
)

type movingAvgInt struct {
	window int
	values []int
}

func New(window int) *movingAvgInt {
	return &movingAvgInt{
		window: window,
		values: make([]int, 0),
	}
}

func NewInt(window int) *movingAvgInt {
	return &movingAvgInt{
		window: window,
		values: make([]int, 0),
	}
}

func (ma *movingAvgInt) Values() []int {
	return ma.values
}

func (ma *movingAvgInt) Add(value int) {
	if len(ma.values) == ma.window {
		ma.values = append((ma.values)[1:ma.window], value)
	} else {
		ma.values = append(ma.values, value)
	}
}

func (ma *movingAvgInt) Min() int {
	min := ma.values[0]
	for _, value := range ma.values {
		if value < min {
			min = value
		}
	}
	return min
}

func (ma *movingAvgInt) Max() int {
	max := ma.values[0]
	for _, value := range ma.values {
		if value > max {
			max = value
		}
	}
	return max
}

func (ma *movingAvgInt) Sum() int {
	var sum int
	values := ma.values
	if values == nil {
		return 0
	}

	for _, value := range values {
		sum += value
	}
	return sum
}

func (ma *movingAvgInt) Median() int {
	values := ma.values
	c := len(values)

	if c == 0 {
		return 0
	}

	sort.Ints(values)
	if c%2 == 1 {
		return values[c/2]
	}
	return (values[c/2] + values[c/2-1]) / 2
}

func (ma *movingAvgInt) Arithmetic() int {
	values := ma.values
	c := len(values)

	if c == 0 {
		return 0
	}

	sum := 0

	for _, value := range ma.values {
		sum += value
	}
	return sum / c
}

func (ma *movingAvgInt) Geometric() float64 {
	values := ma.values
	c := len(values)

	var m float64 = 1
	for _, value := range values {
		m = m * float64(value)
	}
	return math.Pow(m, 1/float64(c))
}
