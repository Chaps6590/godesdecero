package main

import (
	"fmt"

	"github.com/Chaps6590/godesdecero/variables"
)

func main() {

	estado, texto := variables.ConviertoaTexto(1098)
	fmt.Println(estado)
	fmt.Println(texto)

}
