package utils

import (
	"math/rand"
	"strconv"
)

var charsetNumber = []byte("123456789")
var charsetString = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func StringToInt(param string) int {
	i, err := strconv.Atoi(param)
	if err != nil {
		Log(err.Error(), "err", "utils")
	}
	return i
}


func RandCode(n int, types string) string {
	b := make([]byte, n)
	for i := range b {
		if types == "number" {
			b[i] = charsetNumber[rand.Intn(len(charsetNumber))]
		}else{
			b[i] = charsetString[rand.Intn(len(charsetString))]
		}
	}
	return string(b)
}