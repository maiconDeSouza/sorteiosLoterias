package sorteio

import (
	"math/rand"
	"time"
)

func NumerosAleatorios(totalDeNumeros int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	n := r.Intn(int(totalDeNumeros)) + 1

	return n
}
