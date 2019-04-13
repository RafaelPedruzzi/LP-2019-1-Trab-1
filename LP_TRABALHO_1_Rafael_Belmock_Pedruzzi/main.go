/******************************************************************
Linguagens de Programação - Prof. Flavio Varejão - 2019-1
Primeiro trabalho de implementação

Aluno: Rafael Belmock Pedruzzi

main.go:	módulo main
*******************************************************************/
package main

import "fmt"

func main() {
	// Obtendo a distância mínima e a lista de pontos dos arquivos de entrada
	dist, p := readEntry()

	points := *p
	fmt.Printf("%v %T\n", dist, dist)
	fmt.Printf("len=%d %v\n", len(points), points)
	// fmt.Printf("%f\n", points[1].Dist(points[1]))
	// fmt.Printf("%f\n", points[1].Dist(points[3]))
	// fmt.Printf("%f\n", points[3].Dist(points[4]))
	// fmt.Printf("%f\n", points[1].Dist(points[4]))

	// Realizando oo algoritimo de agrupamento
	g := makeGroups(dist, p)

	fmt.Printf("%v\n", g.groups)

	// Calculando o SSE do agrupamento
	sse := g.sse()

	// Imprimindo arquivos de saida
	writeSSE(sse)
	writeGroups(g)
}
