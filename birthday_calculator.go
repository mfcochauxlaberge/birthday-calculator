package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	ns := []uint64{0, 1, 2, 5, 10, 20, 23, 30, 40, 50, 60, 70, 80, 90, 100, 200, 300, 365}

	d := float64(365)
	ps := map[uint64]float64{}

	for _, n := range ns {
		ps[n] = 1
		for i := float64(0); i < float64(n); i++ {
			ps[n] *= (d - i) / d
		}

		ps[n] = (1 - ps[n]) * 100
	}

	fmt.Printf("p=%.0f\n", d)
	fmt.Printf("+------------+------------+\n")
	for _, n := range ns {
		nStr := strconv.FormatUint(n, 10) + " "
		for len(nStr) < 12 {
			nStr = " " + nStr
		}

		pStr := fmt.Sprintf("%3.4f%% ", ps[n])
		for len(pStr) < 12 {
			pStr = " " + pStr
		}

		fmt.Printf("|%s|%s|\n", nStr, pStr)
	}
	fmt.Printf("+------------+------------+\n")
	fmt.Printf("Calculated in %.3f seconds.\n", time.Now().Sub(start).Seconds())
}
