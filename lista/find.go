package lista

import "errors"

func Find(nomeDaLoteria string) (error, int, int) {
	for _, loteria := range ListaDeLoterias {
		if loteria.nomeDaLoteria == nomeDaLoteria {
			return nil, loteria.totalDeNumeros, loteria.numerosSorteados
		}
	}

	err := errors.New("nome do jogo incoreto")

	return err, 0, 0
}
