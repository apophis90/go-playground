package popcount

/*PopCount is yet another implementation which determines the population
count by resetting the rightmost non-zero bit to zero and counting the
number of iterations required until x is 0, i.e. all non-zero bits have
been cleared.*/
func PopCount(x uint64) int {
	bitsCleared := 0
	for x != 0 {
		x = x & (x - 1)
		bitsCleared++
	}
	return bitsCleared
}
