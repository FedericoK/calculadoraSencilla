package main

import (
	//bufio se usa para la entrada de datos
	"bufio"
	//os se usa para interactuar con el sistema operativo
	"os"
	//fmt para imprimir en consola y formatear a string
	"fmt"
	//para convertir strings a numeros
	"strconv"
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

		//parte en secciones el input del usuario dependiendo de los espacios, en este caso en 3 partes "numero[0] operador[1] numero[2]"
		partes := strings.Fields(input)

		//Verifica que se haya escrito correctamente la operacion
		if len(partes) != 3 {
			//si es incorrecto - mensaje de error y vuelve al incio
			fmt.Println("Formato incorrecto. Asegurece de usar el formato 'a + b'")
			continue
		}

		//definimos los operandos y el operador y convertimos las partes de los operandos [0] y [2] de Sting a Float64
		operando1, err1 := strconv.ParseFloat(partes[0], 64)
		operando2, err2 := strconv.ParseFloat(partes[2], 64)
		operador := partes[1]

		//Verifica la conversion de los operandos cuando se convierten letras a float se hacen 0
		if err1 != nil || err2 != nil {
			//si hay error - muestra mensaje y vuelve al inicio
			fmt.Println("Error: Asegurese de que ambos operandos sean numeros.")
			continue
		}

		//Usamos un switch para realizar la operacion dependiendo del mismo
		switch operador {
		case "+":
			fmt.Printf("Resultado: %v\n", operando1+operando2)
		case "-":
			fmt.Printf("Resultado: %v\n", operando1-operando2)
		case "*":
			fmt.Printf("Resultado: %v\n", operando1*operando2)
		case "/":
			//si es division hay que verificar que no divida entre 0
			if operando2 == 0 {
				fmt.Println("Error: Division por cero.")
			} else {
				fmt.Printf("Resultado: %v\n", operando1/operando2)
			}
		default:
			fmt.Printf("Operador desconocido. Utilice +, -, *, /.")
		}
	}
}
