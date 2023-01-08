package main

func splitStringInMiddle(s string) (string, string) {
	var left string
	var right string

	for i := 0; i < len(s)/2; i++ {
		left = left + string(s[i])
	}

	for i := len(s) / 2; i < len(s); i++ {
		right = right + string(s[i])
	}

	return left, right
}

func permute(k string, arr []int, n int) string {
	var permutation = ""

	for i := 0; i < n; i++ {
		permutation = permutation + string(k[arr[i]-1])
	}
	return permutation
}

// shifting the bits towards left by nth shifts
func shiftLeft(k string, nthShifts int) string {
	var s = ""

	for i := 0; i < nthShifts; i++ {
		for j := 1; j < len(k); j++ {
			s = s + string(k[j])
		}

		s = s + string(k[0])
		k = s
		s = ""
	}

	return k
}

// calculating xor of two strings of binary number a and b
func xor(a, b string) []string {
	var ans []string
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			ans = append(ans, "0")
		} else {
			ans = append(ans, "1")
		}
	}

	return ans
}

func keyGen(key string) (string, []string, []string) {
	var rkb []string
	var rk []string

	key = hex2Bin(key)

	// getting 56 bit key from 64 bit using the parity bits
	key = permute(key, keyP, 56)

	// splitting
	left, right := splitStringInMiddle(key)

	for i := 0; i < 16; i++ {
		left = shiftLeft(left, shiftTable[i])
		right = shiftLeft(right, shiftTable[i])

		combineStr := left + right
		roundKey := permute(combineStr, keyComp, 48)

		rkb = append(rkb, roundKey)
		rk = append(rk, binToHex(roundKey))
	}

	return key, rkb, rk
}
