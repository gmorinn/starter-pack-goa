package utils

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var stringRunes = []rune("abcdefghijklmnopqrstuvwxyz123456789")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = stringRunes[rand.Intn(len(stringRunes))]
	}
	return string(b)
}

var numberRunes = []rune("0123456789")

func RandNumberRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = numberRunes[rand.Intn(len(numberRunes))]
	}
	return string(b)
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomAttribut(tab []string) string {
	n := len(tab)
	return tab[rand.Intn(n)]
}

func RandomEmail() string {
	firstname := RandStringRunes(8)
	lastname := RandStringRunes((8))
	return strings.ToLower(firstname) + "." + strings.ToLower(lastname) + "@example.com"
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxy")

func RandLetterRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
