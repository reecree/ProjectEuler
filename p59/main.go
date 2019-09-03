package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("./p059_cipher.txt")
	if err != nil {
		fmt.Printf("Couldn't read file: %s\n", err.Error())
		return
	}

	encryptedChars := strings.Split(string(dat), ",")

	// We're given encryption key is only 3 chars long
	encryptionKeyLen := 3
	spaceASCIICode := 32
	charCountMaps := make([]map[int]int, encryptionKeyLen)
	mostCommonKey := make([]int, encryptionKeyLen)
	mostCommonVal := make([]int, encryptionKeyLen)
	for i := 0; i < encryptionKeyLen; i++ {
		charCountMaps[i] = make(map[int]int)
	}
	for i, encryptedChar := range encryptedChars {
		encryptedInt, err := strconv.Atoi(encryptedChar)
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		charCountMaps[i%3][encryptedInt]++
		if charCountMaps[i%3][encryptedInt] > mostCommonVal[i%3] {
			mostCommonVal[i%3] = charCountMaps[i%3][encryptedInt]
			mostCommonKey[i%3] = encryptedInt
		}
	}

	decryptors := make([]int, encryptionKeyLen)
	for i, key := range mostCommonKey {
		decryptors[i] = key ^ spaceASCIICode
	}

	var sum int
	for i, encryptedChar := range encryptedChars {
		encryptedInt, _ := strconv.Atoi(encryptedChar)
		asciiVal := decryptors[i%3] ^ encryptedInt

		// Print message
		// fmt.Print(string(asciiVal))
		sum += asciiVal
	}
	fmt.Printf("Total Sum is %d\n", sum)
}
