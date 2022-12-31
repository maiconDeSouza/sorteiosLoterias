package sorteio

import "sorteiosLoterias/lista"

func Sorteio(nomeDaLoteria string) (error, []int) {
	err, totalDeNumeros, numerosSorteados := lista.Find(nomeDaLoteria)
	if err != nil {
		return err, nil
	}

	var sorteio = Sorteador(totalDeNumeros, numerosSorteados)

	return nil, sorteio

}
