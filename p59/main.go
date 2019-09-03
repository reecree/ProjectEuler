package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// We're given encryption key is only 3 chars long
const EncryptionKeyLen = 3

func main() {
	dat, err := ioutil.ReadFile("./p059_cipher.txt")
	if err != nil {
		fmt.Printf("Couldn't read file: %s\n", err.Error())
		return
	}

	encryptedChars := strings.Split(string(dat), ",")

	spaceASCIICode := 32
	charCountMaps := make([]map[int]int, EncryptionKeyLen)
	mostCommonKey := make([]int, EncryptionKeyLen)
	mostCommonVal := make([]int, EncryptionKeyLen)
	for i := 0; i < EncryptionKeyLen; i++ {
		charCountMaps[i] = make(map[int]int)
	}
	for i, encryptedChar := range encryptedChars {
		encryptedInt, err := strconv.Atoi(encryptedChar)
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		ind := i % EncryptionKeyLen
		charCountMaps[ind][encryptedInt]++
		if charCountMaps[ind][encryptedInt] > mostCommonVal[ind] {
			mostCommonVal[ind] = charCountMaps[ind][encryptedInt]
			mostCommonKey[ind] = encryptedInt
		}
	}

	decryptors := make([]int, EncryptionKeyLen)
	for i, key := range mostCommonKey {
		decryptors[i] = key ^ spaceASCIICode
	}

	var sum int
	for i, encryptedChar := range encryptedChars {
		encryptedInt, _ := strconv.Atoi(encryptedChar)
		asciiVal := decryptors[i%EncryptionKeyLen] ^ encryptedInt

		// Print message
		// fmt.Print(string(asciiVal))
		sum += asciiVal
	}
	fmt.Printf("Total Sum is %d\n", sum)
}
