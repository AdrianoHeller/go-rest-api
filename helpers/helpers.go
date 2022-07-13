package helpers

import (
	"errors"
	"math/rand"
	"time"
)

func CreateRandomToken(tokenLength int) (string, error) {
	rand.Seed(time.Now().UnixNano())
	characterList := "abcdefighijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rChars := []rune(characterList)
	var token []rune
	if tokenLength > 10 {
		for i := 0; i <= tokenLength; i++ {
			randomInteger := rand.Intn(len(rChars))
			randomChar := rChars[randomInteger]
			token = append(token, randomChar)
		}
		return string(token), nil
	} else {
		return "", errors.New("token length cannot be lower than 10")
	}
}
