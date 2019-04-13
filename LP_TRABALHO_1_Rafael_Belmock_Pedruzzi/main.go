/******************************************************************
Linguagens de Programação - Prof. Flavio Varejão - 2019-1
Primeiro trabalho de implementação

Aluno: Rafael Belmock Pedruzzi

main.go:	módulo main
*******************************************************************/
package main

import "fmt"

func main() {
	dist, p := readEntry()
	points := *p

	fmt.Printf("%v %T\n", dist, dist)
	fmt.Printf("len=%d %v\n", len(points), points)
	// fmt.Printf("%f\n", points[1].Dist(points[1]))
	fmt.Printf("%f\n", points[1].Dist(points[3]))
	fmt.Printf("%f\n", points[3].Dist(points[4]))
	// fmt.Printf("%f\n", points[1].Dist(points[4]))

	g := makeGroups(dist, p)
	fmt.Printf("%v\n", g.groups)
}
