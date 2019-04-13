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
		panic("Error to open file: distancia.txt")
	}
	defer distancia.Close()

	fmt.Fscan(distancia, &dist) // lendo distância limite.

	// Abrindo arquivo entrada.txt para obter os pontos.
	entrada, err := os.Open("./entrada.txt")
	if err != nil {
		panic("Error to open file: entrada.txt")
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
