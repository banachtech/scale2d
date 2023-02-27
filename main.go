package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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
