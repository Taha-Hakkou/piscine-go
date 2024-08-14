package piscine

func ActiveBits(n int) int {
	activeBits := 0
	for n > 1 {
		activeBits += n % 2
		n /= 2
	}
	if n == 1 {
		activeBits++
	}
	return activeBits
}
