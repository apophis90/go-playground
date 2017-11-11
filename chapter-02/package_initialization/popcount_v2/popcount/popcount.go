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

// PopCount implements another approach which uses a loop for the bit shift
// operations.
func PopCount(x uint64) int {
	var popCount byte
	for i := 0; i < 8; i++ {
		popCount += pc[byte(x>>uint64(i*8))]
	}
	return int(popCount)
}

// PopCountv3 is another version of PopCount which shifts the input number by one
// bit at a time and checks the rightmost bit instead of doing a per-octet check.
func PopCountv3(x uint64) int {
	var popCount byte
	for i := 0; i < 64; i++ {
		popCount += byte((x >> uint64(i)) & 1)
	}
	return int(popCount)
}
