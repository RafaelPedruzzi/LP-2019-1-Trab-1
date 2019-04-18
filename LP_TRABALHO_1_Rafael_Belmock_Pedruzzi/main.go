/******************************************************************
Linguagens de Programação - Prof. Flavio Varejão - 2019-1
Primeiro trabalho de implementação

Aluno: Rafael Belmock Pedruzzi

main.go:	módulo main
*******************************************************************/
package main

func main() {
	// Obtendo a distância mínima e a lista de pontos dos arquivos de entrada
	dist, p := readEntry()

	// Realizando oo algoritimo de agrupamento
	g := makeGroups(dist, p)

	// Calculando o SSE do agrupamento
	sse := g.sse()

	// Imprimindo arquivos de saida
	writeSSE(sse)
	writeGroups(g)
}
