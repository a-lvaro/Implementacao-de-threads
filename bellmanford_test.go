package main

import (
	"math"
	"testing"
)

func TestBellmanFord(t *testing.T) {
	inf := math.Inf(1)

	w := [][]float64{
		{inf, 6, inf, inf, 7},
		{inf, inf, 5, -4, 8},
		{inf, -2, inf, inf, inf},
		{2, inf, 7, inf, inf},
		{inf, inf, -3, 9, inf}}

	p := []int{0, 0, 0, 0, 0}
	d := []float64{0, 0, 0, 0, 0}

	n := len(w)

	expected_p := []int{-1, 2, 4, 1, 0}
	expected_d := []float64{0, 2, 4, -2, 7}

	if !BellmanFord(w, 0, d, p) {
		t.Errorf("expected true")
	}

	for i := 0; i < n; i++ {
		if expected_d[i] != d[i] {
			t.Errorf("p: expected %f at vertex %d but got %f", expected_d[i], i, d[i])
		}

		if expected_p[i] != p[i] {
			t.Errorf("p: expected %d at vertex %d but got %d", expected_p[i], i, p[i])
		}
	}
}

func TestParallelBellmanFord(t *testing.T) {
	inf := math.Inf(1)

	w := [][]float64{
		{inf, 6, inf, inf, 7},
		{inf, inf, 5, -4, 8},
		{inf, -2, inf, inf, inf},
		{2, inf, 7, inf, inf},
		{inf, inf, -3, 9, inf}}

	p := []int{0, 0, 0, 0, 0}
	d := []float64{0, 0, 0, 0, 0}

	n := len(w)

	expected_p := []int{-1, 2, 4, 1, 0}
	expected_d := []float64{0, 2, 4, -2, 7}

	if !ParallelBellmanFord(w, 0, d, p) {
		t.Errorf("expected true")
	}

	for i := 0; i < n; i++ {
		if expected_d[i] != d[i] {
			t.Errorf("p: expected %f at vertex %d but got %f", expected_d[i], i, d[i])
		}

		if expected_p[i] != p[i] {
			t.Errorf("p: expected %d at vertex %d but got %d", expected_p[i], i, p[i])
		}
	}
}

const size int = 1000

var w = buildRandomMatrix(size)
var d = make([]float64, size)
var p = make([]int, size)

func BenchmarkBellmanFord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BellmanFord(w, 0, d, p)
	}
}

func BenchmarkParallelBellmanFord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParallelBellmanFord(w, 0, d, p)
	}
}
