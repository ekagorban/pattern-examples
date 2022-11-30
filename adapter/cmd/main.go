package main

import (
	"adapter/random"
	"log"
	"strconv"
)

// generator - интерфейс генератора чисел
type generator interface {
	nextNum() int64 // вернуть следующее в последовательности число
}

// client - клиент для генерации последовательности
type client struct {
	sequenceLen int      // длина последовательности
	buf         []string // буфер для хранения последовательности
}

// generateSequence - генерирует последовательность, используя генератор
func (c *client) generateSequence(g generator) {
	log.Println("start generate sequence")

	c.buf = make([]string, c.sequenceLen)
	for i := 0; i < c.sequenceLen; i++ {
		c.buf[i] = strconv.FormatInt(g.nextNum(), 10)
	}

	log.Println("success generate sequence")
}

// serialNumber - структура для хранения текущего числа в последовательности
type serialNumber struct {
	n int64
}

// nextNum - возвращает следующее по порядку число
func (s *serialNumber) nextNum() int64 {
	s.n++
	return s.n
}

// адаптер к генератору случайного числа
type randomNumberAdapter struct {
	fNum *random.FloatNumber
}

// nextNum - возвращает следующее число в случайной последовательности
func (a randomNumberAdapter) nextNum() int64 {
	return int64(a.fNum.Generate(10))
}

func main() {
	client := client{
		sequenceLen: 5,
	}
	serialNumber := serialNumber{
		n: 5,
	}

	client.generateSequence(&serialNumber)
	log.Printf("sequence: %v", client.buf)

	randomSequence := random.FloatNumber{}
	randomSequenceAdapter := randomNumberAdapter{
		fNum: &randomSequence,
	}

	client.generateSequence(randomSequenceAdapter)
	log.Printf("sequence: %v", client.buf)
}
