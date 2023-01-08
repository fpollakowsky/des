package main

import (
	"fmt"
	"strings"
)

func main() {
	pt := "123456ABCD132536"
	key := "AABB09182736CCDD"

	println("Key Generation")
	key, rkb, rk := keyGen(key)
	println(key)

	println("Encryption")
	cipherText := binToHex(encrypt(pt, rkb, rk))
	fmt.Println("Final: ", cipherText)
}

func encrypt(s string, rkb []string, rk []string) string {
	pt := hex2Bin(s)

	// initial permutation
	pt = permute(pt, initialPerm, 64)
	ptHex := binToHex(pt)
	fmt.Println("Initial Permutation: " + ptHex)

	// split
	left, right := splitStringInMiddle(pt)

	for i := 0; i < 16; i++ {
		// Expansion D-box: Expanding the 32 bits data into 48 bits
		rightExpanded := permute(right, expD, 48)

		// XOR RoundKeyMatrix[i] and right_expanded
		xorX := xor(rightExpanded, rkb[i])

		// S-boxex: substituting the value from s-box table by calculating row and column
		sboxString := ""

		for j := 0; j < 8; j++ {
			row := binToDec(xorX[j*6] + xorX[j*6+5])
			col := binToDec(xorX[j*6+1] + xorX[j*6+2] + xorX[j*6+3] + xorX[j*6+4])

			val := sbox[j][row][col]
			sboxString = sboxString + decToBin(val)
		}

		// Straight D-box: After substituting rearranging the bits
		sboxString = permute(sboxString, per, 32)

		// XOR left and sbox_str
		result := xor(left, sboxString)
		left = strings.Join(result, "") // check

		if i != 15 {
			left, right = right, left
		}

		println("Round: ", i+1, " ", binToHex(left), " ", binToHex(right), " ", rk[i])
	}

	combination := left + right

	cipherText := permute(combination, final_perm, 64)
	return cipherText
}
