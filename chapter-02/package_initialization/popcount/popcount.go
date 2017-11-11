/*Package popcount encapsulates the functionality to determine the population
count (short: popcount, also called "Hamming weight") of a 64-bit number. the
popcount of a number says how many bits are set to "1" in the binary representation
of that number.
*/
package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	// A 64-bit integer consists of 8 octets of 8 bits each. The following computation
	// right-shifts the integer one octet at a time and looks up the popcount for
	// the rightmost octet which must be a number between 0 and 256.
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
