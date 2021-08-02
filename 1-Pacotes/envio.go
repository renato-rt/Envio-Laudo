package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

//Validando formato do e-mail
func main() {

	validate := checkmail.ValidateFormat("tomikawa.renato@gmail.com")
	if validate == nil {
		fmt.Println("Formato de e-mail Válido")
	} else {
		fmt.Println("Formato de e-mail Inválido")
	}
}
