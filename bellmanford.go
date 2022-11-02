package main

import "math"

/*
Inicializa todos os vértices do grafo com seus
predecessores e estimativas de pes do caminho mínimo
partindo da fonte

w -> matriz de adjacência
d -> estimativas de peso
p -> predecessores
*/
func InitSingleSource(w [][]float64, source int, d []float64, p []int) {
	n := len(w)

	for i := 0; i < n; i++ {
		p[i] = -1
		d[i] = math.Inf(1)
	}

	d[source] = 0 //fonte
}

/*
Relaxamento de arestas: testa se é possível melhorar
o caminho mínimo da fonte até v, analisando o peso
da aresta (u, v)
*/
func Relax(w [][]float64, d []float64, p []int, u, v int) {
	if d[v] > d[u]+w[u][v] {
		d[v] = d[u] + w[u][v]
		p[v] = u
	}
}

/*
Algoritmo de Bellman-Ford: encontra o caminho mínimo da fonte
até todos os vértices do grafo. Retorna false se existe um ciclo
de peso negativo que pode ser alcançado da fonte. Caso contrário
retorna true e produz os caminhos mínimos e seus pesos.

w -> matriz de adjacência
*/
func BellmanFord(w [][]float64, source int, d []float64, p []int) bool {
	n := len(w)

	InitSingleSource(w, source, d, p)

	// Relaxa cada aresta |V| - 1 vezes
	for i := 0; i < n-1; i++ {
		// Itera por todas as arestas
		for u := 0; u < n; u++ {
			for v := 0; v < n; v++ {
				if !math.IsInf(w[u][v], 0) {
					Relax(w, d, p, u, v)
				}
			}
		}
	}

	// Verifica se existe um ciclo de peso negativo
	for u := 0; u < n; u++ {
		for v := 0; v < n; v++ {
			if !math.IsInf(w[u][v], 0) && d[v] > d[u]+w[u][v] {
				return false
			}
		}
	}

	return true
}

func RelaxBF(k int, w [][]float64, d []float64, p []int) {
	n := len(w)

	for i := 0; i < k; i++ {
		for u := 0; u < n; u++ {
			for v := 0; v < n; v++ {
				if !math.IsInf(w[u][v], 0) {
					Relax(w, d, p, u, v)
				}
			}
		}
	}
}

func ParallelBellmanFord(w [][]float64, source int, d []float64, p []int) bool {
	n := len(w)

	m := int(math.Floor(float64(n) / 2))

	InitSingleSource(w, source, d, p)

	// Relaxa cada aresta |V| - 1 vezes
	go RelaxBF(m, w, d, p)
	RelaxBF(n-m, w, d, p)

	// Verifica se existe um ciclo de peso negativo
	for u := 0; u < n; u++ {
		for v := 0; v < n; v++ {
			if !math.IsInf(w[u][v], 0) && d[v] > d[u]+w[u][v] {
				return false
			}
		}
	}

	return true
}
