# movingavg

In statistics, a moving average is a calculation to analyze data points by creating a series of averages of different subsets of the full data set. 

## usage

```golang
package main

import (
    "simonwaldherr.de/go/movingavg"
    "fmt"
)

func main() {
    ma := movingavg.New(3)
    ma.Add(1)
    ma.Add(2)
    ma.Add(3)
    ma.Add(5)
    fmt.Printf("Avg: %d", ma.Arithmetic())
}
```



