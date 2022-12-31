package lista

type Loterias struct {
	nomeDaLoteria    string
	totalDeNumeros   int
	numerosSorteados int
}

var ListaDeLoterias = []Loterias{
	{
		nomeDaLoteria:    "mega-sena",
		totalDeNumeros:   60,
		numerosSorteados: 6,
	},
	{
		nomeDaLoteria:    "lotofacil",
		totalDeNumeros:   25,
		numerosSorteados: 15,
	},
}
