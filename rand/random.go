package rand

import (
	stdrand "crypto/rand"
	"math/big"
	"unicode"
)

func GenerateString(length int) (string, error) {
	var result string
	for len(result) < length {
		num, err := stdrand.Int(stdrand.Reader, big.NewInt(int64(127)))
		if err != nil {
			return "", err
		}
		n := num.Int64()
		if unicode.IsLetter(rune(n)) {
			result += string(n)
		}
	}
	return result, nil
}
