package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
	"math"
	"sort"
)

func main() {
	i := 1
	slice := make([]int, 10)
	fmt.Println("\nENTER AN ARRAY (PRESS 0 TO STOP) : ")
	for i != 0 {
		fmt.Scan(&i)
		slice = append(slice, i)
	}
	sort.Ints(slice)
	mean := stat.Mean(slice, nil)
	variance := stat.Variance(slice, nil)
	stddev := math.Sqrt(variance)
	median := stat.Quantile(0.5, stat.Empirical, slice, nil)

	fmt.Println("\nMEAN :", mean)
	fmt.Println("\nVARIANCE :", variance)
	fmt.Println("\nMEDIAN :", median)
	fmt.Println("\nSTANDARD DEVIATION :", stddev)
}
