package wsa_lib_utils

import (
	"math/rand"
	"strconv"
	"time"
)

var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numberRunes = []rune("0123456789")

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandNumberRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = numberRunes[rand.Intn(len(numberRunes))]
	}
	return string(b)
}

func GenerateExternalID() string {
	return RandNumberRunes(4) + strconv.Itoa(int(time.Now().UnixMilli()))[5:]
}

func GenerateRandomBalance() float64 {
	val, _ := strconv.ParseFloat(RandNumberRunes(4)+"0000", 64)
	return val
}
