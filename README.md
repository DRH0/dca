# dca
## Decline Curve Analysis Package

This package contains functions useful for fast decline curve analysis of typical oil and gas production data.

dca.Hyperbolic(result [][]float64, Qi float64, Di float64, n float64, Dt float64, length int, delay int32)

This function modifies a previously allocated 2d array in-place for speed, populating it with a table of the following useful columns:
[index][monthly production][start][end][nom month][nom year][effective %]
