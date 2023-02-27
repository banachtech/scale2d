package main

import (
	"fmt"
	"math"
)

// skip nan values
func skip_nan(x []float64) []float64 {
	var y []float64
	for _, v := range x {
		if !math.IsNaN(v) {
			y = append(y, v)
		}
	}
	return y
}

// compute square of an array
func square(x []float64) []float64 {
	y := make([]float64, len(x))
	for i, v := range x {
		if math.IsNaN(v) {
			y[i] = v
		} else {
			y[i] = v*v
		}
	}
	return y
}

// compute sum of an array
func sum(x []float64) float64 {
	sum := 0.0
	for _, v := range x {
		sum += v
	}
	return sum
}

// Compute sample mean
func mean(x []float64) float64 {
	return sum(x) / float64(len(x))
}

// compute variance
func variance(x []float64) float64 {
	m := mean(x)
	sxx := 0.0
	for _, v := range x {
		x0 := v - m
		sxx += x0 * x0
	}
	return sxx / float64(len(x))
}

// row means and vars
func row_stats(x [][]float64) ([]float64, []float64) {
	n := len(x)
	row_means := make([]float64, n)
	row_vars := make([]float64, n)
	for i, row := range x {
		tmp := skip_nan(row)
		row_means[i] = mean(tmp)
		row_vars[i] = variance(tmp)
	}
	return row_means, row_vars
}

// col means and vars
func col_stats(x [][]float64) ([]float64, []float64) {
	m := len(x[0])
	col_means := make([]float64, m)
	col_vars := make([]float64, m)
	for j := 0; j < m; j++ {
		col := make([]float64, len(x))
		for i := range x {
			col[i] = x[i][j]
		}
		tmp := skip_nan(col)
		col_means[j] = mean(tmp)
		col_vars[j] = variance(tmp)
	}
	return col_means, col_vars
}

// summary stats
func summary(x [][]float64) {
	row_means, row_vars := row_stats(x)
	col_means, col_vars := col_stats(x)
	// summarise stats
	fmt.Printf("mean of row means = %f\n", mean(row_means))
	fmt.Printf("mean of col means = %f\n", mean(col_means))
	fmt.Printf("mean of row vars = %f\n", mean(row_vars))
	fmt.Printf("mean of col vars = %f\n", mean(col_vars))
}
