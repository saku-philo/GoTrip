package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { // v はifスコープ内のみで有効, math.Powはべき乗関数
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		math.Pow(3, 4),
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
