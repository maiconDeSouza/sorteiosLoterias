package sorteio

func Some(listaDeNumerosSorteados []int, n int) bool {
	for _, numero := range listaDeNumerosSorteados {
		if numero == n {
			return true
		}
	}

	return false
}
