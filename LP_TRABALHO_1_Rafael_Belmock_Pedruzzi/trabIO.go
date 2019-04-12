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
 *	parametros: nenhum
 *	retorno: a distância limite e uma slice contendo todos os pontos lidos. Cada ponto é representado em uma slice, sendo que a primeira posição dela contem o número da linha que foi lido e as posições seguintes as cordenadas do ponto em ordem de leitura.
 *
 */
func readEntry() (float64, [][]float64) {
	var dist, i float64 = 0, 1 // dist: distância limite ; i: contador da linha atual de leitura
	var points [][]float64     // points: slice de pontos

	// Abrindo arquivo distancia.txt para obter a distância limite
	distancia, err := os.Open("./distancia.txt")
	if err != nil {
		panic("Error to open file: distancia.txt")
	}
	defer distancia.Close()

	fmt.Fscan(distancia, &dist) // lendo distância limite

	// Abrindo arquivo entrada.txt para obter os pontos
	entrada, err := os.Open("./entrada.txt")
	if err != nil {
		panic("Error to open file: entrada.txt")
	}
	defer entrada.Close()

	fileScanner := bufio.NewScanner(entrada) // scanner do arquivo por linhas

	// Escaneando cada linha de entrada.txt e armazenando os pontos em points
	for fileScanner.Scan() {
		var p []float64  // slice auxiliar para armazenar o ponto sendo lido atualmente
		p = append(p, i) // adicionando o número da linha na primeira posição
		i += 1

		// Scanner de cada linha por word, ou seja, por strings separadas por espaços
		lineScanner := bufio.NewScanner(strings.NewReader(fileScanner.Text()))
		lineScanner.Split(bufio.ScanWords)

		// Lendo cada word, convertendo-as para float64 e armazenando em p
		for lineScanner.Scan() {
			f, err := strconv.ParseFloat(lineScanner.Text(), 64) // leitura e conversão e uma cordenada
			if err != nil {
				fmt.Println(err)
				panic("Error: can't convert string to float64")
			}
			p = append(p, f) // adicionando a cordenada ao fim de p
		}

		points = append(points, p) // adicionando p a lista de pontos
	}

	return dist, points
}
