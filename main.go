package main

import (
	"flag"
	"strings"
)

var (
	file    string
	maxiter int
	eps     float64
	verbose bool
	outfile string
)

func init() {
	flag.StringVar(&file, "f", "test.csv", "csv file with matrix data, possibly containing NaN values")
	flag.IntVar(&maxiter, "n", 50, "maximum number of iterations")
	flag.Float64Var(&eps, "e", 0.1, "tolerance for convergence. if residual is less than tolerance, iterations will terminate")
	flag.BoolVar(&verbose, "v", false, "verbose mode: print residuals per iteration")
	flag.StringVar(&outfile, "o", strings.TrimSuffix(file, ".csv")+"_scaled.csv", "output file")
}

func main() {
	flag.Parse()

	// read in matrix
	x := array_from_csv(file)

	// create a Scaler model
	s := NewScaler(x, maxiter, 0.1)

	// fit parameters
	s.Fit(x, verbose)

	// scale 2d array
	y := s.Scale(x)

	// print summary stats
	summary(y)

	// write result to csv
	array_to_csv(y, outfile)
}
