package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

func sha256of(input string) []byte {

	alg := sha256.New()
	alg.Write([]byte(input))
	return alg.Sum(nil)
}

func base58Encoode(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

func GenerateShortLink(initialLink string, userId string) string {
	urlHashBytes := sha256of(initialLink + userId)
	generateNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoode([]byte(fmt.Sprintf("%d", generateNumber)))
	return finalString[:8]
}
