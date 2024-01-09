package main

import (
	//bufio se usa para la entrada de datos
	"bufio"
	//os se usa para interactuar con el sistema operativo
	"os"
	//fmt para imprimir en consola y formatear a string
	"fmt"
	//para convertir strings a numeros

	//Manipulacion de strings
	"strings"
)

func main() {
	//creacion de lector para las entradas del usuario
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Calculadora Simple en Go")
	fmt.Println("------------------------")

	for {
		//Solicitud de entrada
		fmt.Println("Ingrese operacion [formato: a + b]: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		fmt.Println(strings.TrimSpace(input))
	}
}
