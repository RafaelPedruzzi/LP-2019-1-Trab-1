/******************************************************************
Linguagens de Programação - Prof. Flavio Varejão - 2019-1
Primeiro trabalho de implementação

Aluno: Rafael Belmock Pedruzzi

trabIO.go:	módulo responsável pelo tratamento de I/O dos arquivos:
			entrada.txt, distancia.txt, result.txt e saida.txt
*******************************************************************/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 *	Função para leitura dos arquivos entrada.txt e distancia.txt
 *	parâmetros: nenhum.
 *	retorno: a distância limite e um ponteiro para um mapa contendo todos os pontos lidos mapeados pela linha em que foram lidos.
 */
func readEntry() (float64, *map[int]Point) {
	var dist float64 = 0          // dist: distância limite.
	var i int = 1                 // i: contador da linha atual de leitura.
	points := make(map[int]Point) // points: mapa de pontos.

	// Abrindo arquivo distancia.txt para obter a distância limite.
	distancia, err := os.Open("./distancia.txt")
	if err != nil {
		panic(err)
	}
	defer distancia.Close()

	fmt.Fscan(distancia, &dist) // lendo distância limite.

	// Abrindo arquivo entrada.txt para obter os pontos.
	entrada, err := os.Open("./entrada.txt")
	if err != nil {
		panic(err)
	}
	defer entrada.Close()

	fileScanner := bufio.NewScanner(entrada) // scanner do arquivo por linhas.

	// Escaneando cada linha de entrada.txt e armazenando os pontos em points.
	for fileScanner.Scan() {
		var p Point // ponto sendo lido atualmente.

		// Scanner de cada linha por word, ou seja, por strings separadas por espaços.
		lineScanner := bufio.NewScanner(strings.NewReader(fileScanner.Text()))
		lineScanner.Split(bufio.ScanWords)

		// Lendo cada word, convertendo-as para float64 e armazenando em p.
		for lineScanner.Scan() {
			f, err := strconv.ParseFloat(lineScanner.Text(), 64) // leitura e conversão de uma cordenada.
			if err != nil {
				fmt.Println(err)
				panic("Error: can't convert string to float64")
			}
			p = append(p, f) // adicionando a cordenada a p.
		}

		points[i] = p // mapeando o número da linha como chave do ponto.
		i += 1
	}

	return dist, &points
}

/**
 *	Funcão para impressão do arquivo saida.txt
 *	parâmetros: ponteiro para os grupos a serem impressos
 *	pós-condição: estruturas inalteradas.
 */
func writeGroups(g *Groups) {
	// Criando arquivo de escrita
	saida, err := os.Create("saida.txt")
	if err != nil {
		panic(err)
	}
	defer saida.Close()

	// Imprimindo cada grupo. Somente os identifiicadores são impressos, em ordem cressente e separados por espaços. grupos diferentes são separados por uma linha em branco.
	for i := 0; i < len(g.groups); i += 1 {
		for j := 0; j < len(g.groups[i]); j += 1 {
			fmt.Fprintf(saida, "%d ", g.groups[i][j])
		}
		fmt.Fprintf(saida, "\n\n")
	}
}

/**
 *	Funcão para impressão do arquivo result.txt
 *	parâmetros: valor da SSE do agrupamento.
 */
func writeSSE(sse float64) {
	// Criando arquivo de escrita
	saida, err := os.Create("result.txt")
	if err != nil {
		panic(err)
	}
	defer saida.Close()

	// Imprimindo a SSE
	fmt.Fprintf(saida, "%.4f", sse)

}
