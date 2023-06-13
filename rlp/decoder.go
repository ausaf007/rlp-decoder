package rlp

import (
	"encoding/hex"
	"errors"
	log "github.com/sirupsen/logrus"
	"strings"
)

var ItemNotEnclosedInsideList = errors.New("invalid RLP Encoding: multiple items present, " +
	"but not enclosed under a list")

var counter = 0

// validateItems checks if there are more than 1 items, which are not enclosed inside a list
// Throws ItemNotEnclosedInsideList if more than 1 such items are found
func validateItems(depth int) error {
	if depth == 0 {
		counter++
		if counter > 1 {
			return ItemNotEnclosedInsideList
		}
	}
	return nil
}

// handleCase1 decodes single character
// The byte is the RLP encoding of itself
func handleCase1(data []byte, start int, depth int, decodedString string) (string, int) {

	prefix := data[start]
	decodedString += generateSpaces(depth) + "String \"" + string(prefix) + "\"\n"
	start++

	//------------
	return decodedString, start
}

// handleCase2 decodes short string
func handleCase2(data []byte, start int, end int, depth int, decodedString string) (string, int, error) {
	prefix := data[start]
	length := int(prefix - 128)

	if length > end-start-1 {
		err := errors.New("(short string) string length too short than specified in prefix")
		return decodedString, start, err
	}

	decodedString += generateSpaces(depth) + "String \"" + string(data[start+1:start+length+1]) + "\"\n"
	start += 1 + length

	//------------
	return decodedString, start, nil
}

// handleCase3 decodes long string
func handleCase3(data []byte, start int, end int, depth int, decodedString string) (string, int, error) {

	prefix := data[start]
	length := int(prefix - 183) // this is just length of Lengths

	if length > end-start-length-1 {
		err := errors.New("(long string) string length too short than specified in prefix")
		return decodedString, start, err
	}

	actualLength, err := hexToInt(data[start+1 : start+length+1])
	if err != nil {
		return decodedString, start, err
	}
	decodedString += generateSpaces(depth) + "String \"" + string(data[start+length+1:start+length+actualLength+1]) + "\"\n"
	start += 1 + length + actualLength

	//------------
	return decodedString, start, err
}

// handleCase4 decodes short list
func handleCase4(data []byte, start int, depth int, decodedString string) (string, int, error) {

	prefix := data[start]
	length := int(prefix - 192)

	decodedString += generateSpaces(depth) + "List {\n"
	tempDepth := depth
	depth++
	var err error
	decodedString, err = decodeByteArr(data, start+1, start+length+1, depth, decodedString)
	// Returning the function as soon as we get an error
	if err != nil {
		return decodedString, start, err
	}

	decodedString += generateSpaces(tempDepth) + "}\n"
	start += 1 + length
	//depth = tempDepth

	return decodedString, start, nil
}

// handleCase5 decodes long list
func handleCase5(data []byte, start int, depth int, decodedString string) (string, int, error) {

	prefix := data[start]
	length := int(prefix - 247) // this is just length of total payload

	decodedString += generateSpaces(depth) + "List {\n"
	tempDepth := depth

	actualLength, err := hexToInt(data[start+1 : start+length+1])
	if err != nil {
		return decodedString, start, err
	}

	depth++
	decodedString, err = decodeByteArr(data, start+length+1, start+length+actualLength+1, depth, decodedString)
	if err != nil {
		return decodedString, start, err
	}

	decodedString += generateSpaces(tempDepth) + "}\n"

	start += 1 + length + actualLength
	//depth = tempDepth

	return decodedString, start, nil
}

// decodeByteArr takes in byte array data, and decodes it, and returns the decoded string
func decodeByteArr(data []byte, start int, end int, depth int, decodedString string) (string, error) {

	// stores error if it's encountered in any of the codeblocks below
	// the function returns as soon as an error is encountered anywhere
	// by breaking off the for loop
	var err error = nil

	for start < end {

		err = validateItems(depth)
		if err != nil {
			break
		}

		prefix := data[start]
		if prefix <= 127 {
			decodedString, start = handleCase1(data, start, depth, decodedString)

		} else if prefix >= 128 && prefix <= 183 {
			decodedString, start, err = handleCase2(data, start, end, depth, decodedString)

		} else if prefix >= 184 && prefix <= 191 {
			decodedString, start, err = handleCase3(data, start, end, depth, decodedString)

		} else if prefix >= 192 && prefix <= 247 {
			decodedString, start, err = handleCase4(data, start, depth, decodedString)

		} else if prefix >= 248 {
			decodedString, start, err = handleCase5(data, start, depth, decodedString)

		}
		if err != nil {
			break
		}
	}

	return decodedString, err
}

// Decode function assumes that the encoded message contains strings and list of strings only
// For eg, the number 1 has been encoded as "1" (string) and not 1 (number)
// This assumption is made as per the Task requirement
// Decode function decodes the RLP encoded string and returns the decoded output with nested formatting
func Decode(str string) (string, error) {

	counter = 0
	// Converting Input string to byte array
	inputArr, err := hex.DecodeString(str)
	if err != nil {
		// Invalid hex code is detected
		return "", err
	}

	log.Info("Input Array: ", inputArr)

	output, err := decodeByteArr(inputArr, 0, len(inputArr), 0, "")

	return strings.TrimSpace(output), err
}
