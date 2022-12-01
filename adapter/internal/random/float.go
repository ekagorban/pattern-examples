package random

import (
	"crypto/rand"
	"math/big"
)

// FloatNumber - структура для хранения текущего числа в последовательности
type FloatNumber struct {
	f float64
}

// Generate - генерирует случайное число
func (num *FloatNumber) Generate(max uint) float64 {
	if max == 0 {
		return 0
	}
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		num.f = 0
		return num.f
	}
	f, _ := new(big.Float).SetInt(n).Float64()
	num.f = f
	return num.f
}
