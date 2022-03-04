package random_util

import "math/rand"

const hexStr = "0123456789abcdef"
const allStr = "0123456789abcdefghijklmnopqrstuvwxyz"

func GenerateRandomHexStr(length int) string {
	b := make([]rune, 0, length)
	for i := range b {
		b[i] = rune(hexStr[rand.Intn(len(hexStr))])
	}
	return string(b)
}

func GenerateRandomStr(length int) string {
	b := make([]rune, 0, length)
	for i := range b {
		b[i] = rune(allStr[rand.Intn(len(allStr))])
	}
	return string(b)
}
