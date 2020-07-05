package fib

import "testing"

/**test de funcionalidad*/
func TestFibCiclo(t *testing.T) {

	pruebas := []struct {
		n        int
		esperado int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
	}

	for _, prueba := range pruebas {
		if resultado := FibCiclo(prueba.n); resultado != prueba.esperado {
			t.Errorf("fibCiclo(%d)==%d , se espera %d", prueba.n, resultado, prueba.esperado)
		}
	}
}

func TestFibRecursivo(t *testing.T) {
	pruebas := []struct {
		n        int
		esperado int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
	}

	for _, prueba := range pruebas {
		if resultado := FibRecursivo(prueba.n); resultado != prueba.esperado {
			t.Errorf("fibRecursivo(%d)==%d , se espera %d", prueba.n, resultado, prueba.esperado)
		}
	}
}

/**test de rendimiento*/

func BenchmarkFibCiclo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibCiclo(25)
	}
}

func BenchmarkFibRecursivo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibCiclo(25)
	}
}
