package main

import (
	"fmt"
	"math"
)

type Scaler struct {
	// correspond to alpha, beta, tau and gamma
	// of the scale 2d parametrization scheme
	a, b, t, g []float64
	maxiter    int
	eps        float64
}

func NewScaler(x [][]float64, maxiter int, eps float64) *Scaler {
	var s Scaler
	s.a, s.t = row_stats(x)
	s.b, s.g = col_stats(x)
	for i, v := range s.t {
		s.t[i] = math.Sqrt(v)
	}
	for i, v := range s.g {
		s.g[i] = math.Sqrt(v)
	}
	s.maxiter = maxiter
	s.eps = eps
	return &s
}

func (s *Scaler) Scale(x [][]float64) [][]float64 {
	y := make([][]float64, len(x))
	for i := range x {
		y[i] = make([]float64, len(x[i]))
		for j := range x[i] {
			if !math.IsNaN(x[i][j]) {
				y[i][j] = (x[i][j] - s.a[i] - s.b[j]) / s.t[i] / s.g[j]
			} else {
				y[i][j] = x[i][j]
			}
		}
	}
	return y
}

// compute scale 2d parameters
func (s *Scaler) Fit(x [][]float64, verbose bool) {
	for k := 0; k < s.maxiter; k++ {
		// update alpha and tau
		for i := range s.a {
			s.a[i] = ab_calc(x[i], s.b, s.g)
			s.t[i] = tg_calc(x[i], add_scalar(s.b, s.a[i]), s.g)
		}
		// update beta and gamma
		for j := range s.b {
			col := make([]float64, len(x))
			for i := range col {
				col[i] = x[i][j]
			}
			s.b[j] = ab_calc(col, s.a, s.t)
			s.g[j] = tg_calc(col, add_scalar(s.a, s.b[j]), s.t)
		}
		// compute residual
		y := s.Scale(x)
		res := residual(y)
		if verbose {
			fmt.Printf("\niteration %d, residual %f\n", k, res)
		}
		if res < s.eps {
			break
		}
	}
}

// compute residual
func residual(x [][]float64) float64 {
	res := 0.0
	for _, v := range x {
		tmp := skip_nan(v)
		res += math.Pow(mean(tmp), 2.0) + math.Pow(math.Log(mean(square(tmp))), 2)
	}
	for j := range x[0] {
		col := make([]float64, len(x))
		for i := range col {
			col[i] = x[i][j]
		}
		tmp := skip_nan(col)
		res += math.Pow(mean(tmp), 2.0) + math.Pow(math.Log(mean(square(tmp))), 2)
	}
	return res
}

// alpha / beta computation
func ab_calc(x []float64, mu, sigma []float64) float64 {
	s1 := 0.0 // running sum of de-meaned x's
	s2 := 0.0 // running sum of inverse of sigma's
	for i, v := range x {
		if !math.IsNaN(v) {
			s1 += (v - mu[i]) / sigma[i]
			s2 += 1.0 / sigma[i]
		}
	}
	return s1 / s2
}

// tau / gamma computation
func tg_calc(x []float64, mu, sigma []float64) float64 {
	s := 0.0
	n := value_counts(x)
	for i, v := range x {
		if !math.IsNaN(v) {
			tmp := (v - mu[i]) / sigma[i]
			s += tmp * tmp
		}
	}
	return math.Sqrt(s / float64(n))
}
