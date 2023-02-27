# Centering and Scaling 2D Arrays

## Problem

Let $X = \{X_{ij}\} \in \mathrm{R}^{m \times n}$ be a 2d array, with some elements missing. Let $\Omega$ be the set of indices $(i,j)$ for which values are present. Each row and column of $X$ has atleast one value present.

The objective is to standardize rows _and_ columns of $X$ to zero mean and unit variance _simultaneously_. In other words, if $X_{i.}$ is the $i^{th}$ row and $X_{.j}$ is the $j^{th}$ column, then $X_{i.}$ has zero mean and unit variance and $X_{.j}$ has zero mean and unit variance for all $i, j$ (among the observed values).


## Methodology

We parametrise centering and scaling and compute the normaised $\tilde{X}_{ij}$ (the observed entries) as:

$$ \tilde{X}_{ij} = \frac{X_{ij} - \alpha_i - \beta_j}{\tau_i\gamma_j} $$

So there are a total of $2m+2n$ parameters. Estimation of parameters starts with a guess of parameters and iteratively refines them till a convergence criterion is met. The updating formulas are:

$$ \alpha_i = \frac{\sum_{j \in \Omega_i} \frac{1}{\gamma_j}(X_{ij}-\beta_j)}{\sum_{j \in \Omega_i}\frac{1}{\gamma_j}}, i = 1,\ldots,m $$

$$ \beta_j = \frac{\sum_{i \in \Omega^j} \frac{1}{\tau_i}(X_{ij}-\alpha_i)}{\sum_{i \in \Omega^j}\frac{1}{\tau_i}}, j = 1,\ldots,n $$

$$ \tau_i^2 = \frac{1}{n_i}\sum_{j \in \Omega_i}\frac{(X_{ij}-\alpha_i -\beta_j)^2}{\gamma_j^2}, i = 1, \ldots, m $$

$$ \gamma_j^2 = \frac{1}{m_j}\sum_{i \in \Omega^j}\frac{(X_{ij}-\alpha_i -\beta_j)^2}{\tau_i^2}, i = 1, \ldots, m $$

where $n_i$ is the number number of observed values in row $i$ and $m_j$ is the number of observed values in column $j$.

$\Omega_i = \{k: (i,k) \in \Omega\}$ and $\Omega^j = \{k: (k,j) \in \Omega\}$. Therefore, $n_i = \vert \Omega_i \vert, m_j = \vert \Omega^j$.


We stop the iteration when residual as defined below is suffiently close to zero.

$$
R &= \sum_{i=1}^m \left[\frac{1}{n_i}\sum_{j \in \Omega_i}\tilde{X}_{ij}\right]^2 + \sum_{j=1}^n \left[\frac{1}{m_j}\sum_{i \in \Omega^j}\tilde{X}_{ij}\right]^2 \\ &+ \sum_{i=1}^m \log^2\left(\frac{1}{n_i}\sum_{j \in \Omega_i}\tilde{X}_{ij}^2\right) + \sum_{j=1}^n \log^2\left(\frac{1}{m_j}\sum_{i \in \Omega^j}\tilde{X}_{ij}^2\right)
$$

## Usage
The executable scale2d is compiled for MacOS.

```bash
scale2d > ./scale2d -h
Usage of ./scale2d:
  -e float
    	tolerance for convergence. if residual is less than tolerance, iterations will terminate (default 0.1)
  -f string
    	csv file with matrix data, possibly containing NaN values (default "test.csv")
  -n int
    	maximum number of iterations (default 50)
  -o string
    	output file (default "test_scaled.csv")
  -v	verbose mode: print residuals per iteration
```

To compile for Windows-64, build as follows:
```bash
GOOS=windows GOARCH=amd64 go build .
```