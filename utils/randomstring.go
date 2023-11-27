package utils

import "math/rand"

var letters = []rune("abcdefghijklmnopqrstuvwxyQWERTYUIOPASDFGHJKLZXCVBNM1234567890")

func RandomString(n int) string {
	b := make([]rune, n)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)

}
