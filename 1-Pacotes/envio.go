package main

import (
	"fmt"
	
	"github.com/badoux/checkmail"
)

//Validando formato do e-mail
func main() {

	validate := checkmail.ValidateFormat("tomikawa.renatomail.com")
	if validate == "invalid format" {
		fmt.Println("Formato de e-mail inválido")
	} else {
		fmt.Println("Formato de e-mail válido")
	}
}
