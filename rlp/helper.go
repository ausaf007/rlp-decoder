package rlp

import (
	"encoding/hex"
	"strconv"
	"strings"
)

// generateSpaces creates a given number of 2 space-indentations
func generateSpaces(num int) string {
	var sb strings.Builder
	for i := 0; i < num; i++ {
		sb.WriteString("  ")
	}
	return sb.String()
}

// hexToInt takes hex in the form of byte array and returns integer
// Ex: input = [4 0]
// output = 1024
func hexToInt(data []byte) (int, error) {
	str := hex.EncodeToString(data) // convert byte array to string
	outputInt, err := strconv.ParseInt(str, 16, 64)

	return int(outputInt), err
}
