package sorteio

import "sort"

func Sorteador(totalDeNumeros, numerosSorteados int) []int {
	var listaDeNumerosSorteados []int
	isValid := true

	for isValid {
		n := NumerosAleatorios(totalDeNumeros)

		jaFoiSorteado := Some(listaDeNumerosSorteados, n)

		if !jaFoiSorteado {
			listaDeNumerosSorteados = append(listaDeNumerosSorteados, int(n))
		}

		if len(listaDeNumerosSorteados) == int(numerosSorteados) {
			isValid = false
		}
	}

	sort.Ints(listaDeNumerosSorteados)

	return listaDeNumerosSorteados
}
