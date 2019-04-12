/******************************************************************
Linguagens de Programação - Prof. Flavio Varejão - 2019-1
Primeiro trabalho de implementação

Aluno: Rafael Belmock Pedruzzi

main.go:	módulo main
*******************************************************************/
package main

import "fmt"

func main() {
	dist, points := readEntry()

	fmt.Printf("%v %T\n", dist, dist)
	fmt.Printf("len=%d cap=%d %v\n", len(points), cap(points), points)
}
