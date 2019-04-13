/******************************************************************
Linguagens de Programação - Prof. Flavio Varejão - 2019-1
Primeiro trabalho de implementação

Aluno: Rafael Belmock Pedruzzi

point.go:	módulo responsável pela implementação dos calculos e
			estruturas feitos com pontos muldidimensionais
*******************************************************************/
package main

import "math"

type Point []float64

/**
 *	Struct que define um grupo de pontos
 *	groups: slice representando o conjunto de grupos. Cada grupo é representado por uma slice de inteiros positivos que são os ids dos pontos mapesdos em points. O primeiro ponto de cada grupo é o lider do grupo.
 *	points: mapa de pontos.
 */
type Groups struct {
	groups [][]int
	points *map[int]Point
}

/**
 *	Método de Point que calcula a distância euclidiana entre dois pontos
 *	parâmetros: um ponto p2.
 *	retorno: a distância euclidiana entre p1 e p2.
 *	condição: p1 e p2 devem ter o mesmo número de dimensões.
 */
func (p1 Point) Dist(p2 Point) float64 {
	var sum, sub float64 = 0, 0
	for i := 0; i < len(p1); i++ {
		sub = (p1[i] - p2[i])
		sum += sub * sub
	}
	return math.Sqrt(sum)
}

/**
 *	Função que monta os grupos segundo o algoritimo de agrupamento por líder
 *	parâmetros: a distância maxima entre um ponto e seu líder e u ponteiro para o mapa de pontos.
 *	retorno: um struct Groups contendo os grupos formados.
 *	condição: todos os pontos devem ter o mesmo número de dimensões.
 */
func makeGroups(dist float64, p *map[int]Point) *Groups {
	g := Groups{points: p} // inicializando g com o ponteiro para o mapa de pontos.
	pts := *(g.points)     // mapa de pontos.
	var lider bool         // variável auxiliar usada para reconhecer novos líderes.

	// Montando os grupos.
	// Criando o primeiro grupo e adicionando o primeiro ponto como seu líder.
	g.groups = append(g.groups, make([]int, 1))
	g.groups[0][0] = 1

	// Adicionando/criando os demais pontos/grupos.
	for i := 2; i <= len(pts); i += 1 { // para cada ponto i no mapa de pontos.
		for j := 0; j < len(g.groups); j += 1 { // para cada grupo j em g.

			p := g.groups[j][0] // posição do líder do grupo j no mapa.
			lider = true

			// Verificando se a distância do ponto i ao lider do grupo j é menor ou igual a dist. Caso verdadeiro, i é adicionado a j.
			if pts[i].Dist(pts[p]) <= dist {
				g.groups[j] = append(g.groups[j], i)
				lider = false
				break
			}
		}
		// Caso i seja líder, um novo grupo será criado para i.
		if lider {
			g.groups = append(g.groups, make([]int, 1))
			g.groups[len(g.groups)-1][0] = i
		}
	}

	return &g
}
