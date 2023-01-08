package main

import (
	"fmt"
	"log"
	"strconv"
)

func hex2Bin(in string) string {
	ui, err := strconv.ParseUint(in, 16, 64)
	if err != nil {
		return ""
	}

	format := fmt.Sprintf("%%0%db", len(in)*4)
	val := fmt.Sprintf(format, ui)

	return val
}

func binToHex(in string) string {
	ui, err := strconv.ParseUint(in, 2, 64)
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}

	return fmt.Sprintf("%x", ui)
}

func binToDec(in string) int64 {
	output, err := strconv.ParseInt(in, 2, 64)
	if err != nil {
		return 00
	}

	return output
}

func decToBin(in int) string {
	output := strconv.FormatInt(int64(in), 2)

	for {
		if len(output) != 4 {
			output = "0" + output
		} else {
			break
		}
	}

	return output
}
