package utils

import (
	"log"
	"strconv"
)

// HexToIntConversion converts hex into int format
func HexToIntConversion(hex string) (int, error) {
	val := hex[2:]
	n, err := strconv.ParseUint(val, 16, 64)
	if err != nil {
		log.Printf("Error while converting hex to int : %v", err)
		return int(n), err
	}
	n2 := int(n)

	return n2, err
}
