// package name is suggested to match the folder name
package math

// the "Average" is in capital, means that other packages/programs are able to see it
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}
