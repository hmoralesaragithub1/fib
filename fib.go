// el paquete fib contiene dos funciones para calcular la seccion fibonacci usalo de forma confiable

package fib

/*llamamos a los Fib*/
func Fib() int {
	return FibCiclo(10)
}

/*FibCiclo obtiene la secuencia fibonacci de la forma original
mas optima en go*/
func FibCiclo(n int) int {
	a, b := 0, 1

	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

// FibRecursivo obtiene la secuencia fibonacci de manera recursiva
// menos optima en go
func FibRecursivo(n int) int {
	if n < 2 {
		return n
	}
	return FibRecursivo(n-1) + FibRecursivo(n-2)
}

func main() {}
