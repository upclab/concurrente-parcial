package main

import "fmt"

func sync() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Iteracion %d: \n", i)
		for j := 1; j <= 8; j++ {
			var a int
			for k := 1; k <= 10; k++ {
				a += j * k
			}
			fmt.Printf("Punto Suma %d - %d: %d\n", i, j, a)
		}
		fmt.Printf("Fin Puntos Suma de %d\n\n", i)
	}
}

func async() {
	for i := 1; i <= 5; i++ {
		ch := make(chan int)

		fmt.Printf("Iteracion %d: \n", i)

		for j := 1; j <= 8; j++ {
			go func(iC int, jC int) {
				var a int

				for k := 0; k <= 10; k++ {
					a += jC * k
				}
				ch <- a
			}(i, j)
		}

		for j := 1; j <= 8; j++ {
			fmt.Printf("Punto Suma %d - %d: %d\n", i, j, <-ch)
		}

		fmt.Printf("Fin Puntos Suma de %d\n\n", i)
	}
}

func main() {
	sync()
}
