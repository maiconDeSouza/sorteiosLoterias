package main

import (
	"fmt"
	"sorteiosLoterias/sorteio"
)

func main() {
	err, n := sorteio.Sorteio("lotofacil")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n)
}
