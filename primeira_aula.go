package main

import (
	"fmt"
	"math"
	"strings"
)


const pi float64 = 3.141592

//funções são declaradas com func
	func hello(message string) (string, string) {
		return message, strings.ToUpper(message)
	}

	func sum(nums ...int) (int, int) {
		total := 0

		for _, num := range nums {
			total += num
		}

		return len(nums), total
	}

func main() {
	idade := 21 //int
	var nome string = "Gabriela"
	ehLinda := true // bool

	//var não precisam ser inicializados no momento
	//no entanto, todos as variaveis devem ser usadas, se não não compila

	fmt.Print("Olá Mundo!\n")
	fmt.Printf("A %v tem %v anos e é muito linda %v\n", nome, idade, ehLinda)

	//Operadores aritméticos: + - * / % em golang são iguais em C
	//Golang não faz casting automático de dados igual em C, se dividir int por float64, é necessário fazer o casting manualmente

	x, y := 5.0, 2.0
	z := x * y + math.Pow(x, 3)

	//Operadores lógicos também são iguais em C: && || !

	//Condicionais também são iguais em C, mas não precisa de parênteses

	if (z > 0 && math.Sqrt(5*5) == 5) {
		fmt.Print("Oi\n")
	}

	//É possível iterar sobre arrays, slices, maps e strings com for e range
	//Go não tem while, nem do while, só for.
	//Para o for se comportar como while, basta escrever só for

	for i := range 10{
		fmt.Println(i)
	}

	fmt.Println(hello("oii"))

	fmt.Println((sum(1, 2, 4, 8, 21, 9)))

	//arrays podem ser declarados vazios
	//arrays tem capacidade fixa e limitada, não da pra redimensionar um array

	var nums[5] int
	nums[0] = 10
	nums[1] = 20
	fmt.Println(nums)

	//declaração e inicialização direta
	letters := [3]string{"a", "b", "c"}
	fmt.Println(letters)

	//slices são mais flexíveis que arrays, podem crescer e diminuir de tamanho
	s := []int{1, 2, 3}
	s = append(s, 4, 5)
	fmt.Println(s)

	//Da pra criar um slice a partir de um array
	arr := [5]int{10, 20, 30, 40, 50}
	slice := arr[1:4]
	fmt.Println(slice)
	
	//maps são como dicionários em python, ou objetos em JS
	m := make(map[string]int)
	m["Alice"] = 25
	m["Bob"] = 30
	fmt.Println(m)

	//acessar valor
	age := m["Alice"]
	fmt.Println("Alice tem", age, "anos")

	//verificar se chave existe
	if val, ok := m["Charlie"]; ok {
		fmt.Println("Charlie existe: ", val)
	} else {
		fmt.Println("Charlie não existe")
	}

	//deletar chave
	delete(m, "Bob")
	fmt.Println(m)

}