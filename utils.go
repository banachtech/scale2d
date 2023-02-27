package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// log fatal errors
func log_err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// read in 2d float array from csv file
// csv file has no headers
func array_from_csv(filename string) [][]float64 {
	f, err := os.Open(filename)
	log_err(err)
	defer f.Close()

	r := csv.NewReader(f)
	data, err := r.ReadAll()
	log_err(err)

	array2d := make([][]float64, len(data))
	for i, v := range data {
		array2d[i] = make([]float64, len(v))
		for j, u := range v {
			array2d[i][j], err = strconv.ParseFloat(u, 64)
			log_err(err)
		}
	}
	return array2d
}

// write array to csv file
func array_to_csv(x [][]float64, filename string) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	log_err(err)
	defer f.Close()

	w := csv.NewWriter(f)
	s := make([][]string, len(x))
	for i, v := range x {
		s[i] = make([]string, len(v))
		for j, u := range v {
			s[i][j] = fmt.Sprintf("%v", u)
		}
	}
	w.WriteAll(s)
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

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
			y[i] = v * v
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

// add scalar to vector
func add_scalar(x []float64, a float64) []float64 {
	y := make([]float64, len(x))
	for i, v := range x {
		y[i] = v + a
	}
	return y
}

// count non nans
func value_counts(x []float64) int {
	count := 0
	for _, v := range x {
		if !math.IsNaN(v) {
			count++
		}
	}
	return count
}
