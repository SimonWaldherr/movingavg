package movingavg_test

import (
	movingavg "../movingavg"
	"fmt"
)

func ExampleMovingAverage() {
	values := movingavg.New(5)
	values.Add(1)
	values.Add(2)
	values.Add(3)

	fmt.Printf("Values: %#v\nMin: %d, Max: %d, Sum: %d, Median: %d, Arithmetic: %d, Geometric: %.1f\n\n",
		values.Values(),
		values.Min(),
		values.Max(),
		values.Sum(),
		values.Median(),
		values.Arithmetic(),
		values.Geometric(),
	)

	values.Add(4)
	values.Add(5)
	values.Add(6)
	values.Add(7)

	fmt.Printf("Values: %#v\nMin: %d, Max: %d, Sum: %d, Median: %d, Arithmetic: %d, Geometric: %.1f\n\n",
		values.Values(),
		values.Min(),
		values.Max(),
		values.Sum(),
		values.Median(),
		values.Arithmetic(),
		values.Geometric(),
	)

	values2 := movingavg.New(5)
	values2.Add(1)
	values2.Add(9)
	fmt.Printf("Values: %#v\nMin: %d, Max: %d, Sum: %d, Median: %d, Arithmetic: %d, Geometric: %.1f\n\n",
		values2.Values(),
		values2.Min(),
		values2.Max(),
		values2.Sum(),
		values2.Median(),
		values2.Arithmetic(),
		values2.Geometric(),
	)

	// Output:
	// Values: []int{1, 2, 3}
	// Min: 1, Max: 3, Sum: 6, Median: 2, Arithmetic: 2, Geometric: 1.8
	//
	// Values: []int{3, 4, 5, 6, 7}
	// Min: 3, Max: 7, Sum: 25, Median: 5, Arithmetic: 5, Geometric: 4.8
	//
	// Values: []int{1, 9}
	// Min: 1, Max: 9, Sum: 10, Median: 5, Arithmetic: 5, Geometric: 3.0
}
