package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	cepParam := "32450000"
	// Validando se o CEP informado é válido.
	if validarFormatoCEP(cepParam) {
		fmt.Printf("O CEP %s é válido.\n\n", cepParam)
	} else {
		fmt.Printf("O CEP %s é inválido. Deve ter exatamente 8 caracteres e ser composto apenas por números inteiros.\n\n", cepParam)
		os.Exit(1)
	}

}

// Função que valida o CEP informado por parâmetro
func validarFormatoCEP(parametro string) bool {
	// Verifica se o parâmetro tem exatamente 8 caracteres
	if len(parametro) != 8 {
		return false
	}

	// Verifica se todos os caracteres são números inteiros
	_, err := strconv.Atoi(parametro)
	return err == nil
}
