package main

import "fmt"

func main() {

	// for k := 0; k < 10; k++ {
	// 	for l := 0; l < 10; l++ {
	// 		if k == l {
	// 			mat[k][l] = 1
	// 		} else {
	// 			mat[k][l] = 0
	// 		}
	// 	}

	// }
	//fmt.Println(mat)
	elegirImpresion()

}

func elegirImpresion() {

	var nombreFiguraImprimir string
	fmt.Println("Elija la forma que desea imprimr la figura: Derecha, Izquierda, X con matriz (Con) o en X sin matrices (Sin)")
	fmt.Scan(&nombreFiguraImprimir)
	fmt.Println("Usted eligio la impresion de la figura ", nombreFiguraImprimir)

	switch nombreFiguraImprimir {
	case "Derecha":
		impresionDerecha()
	case "Izquierda":
		impresionIzquierda()
	case "Con":
		impresionXConMatriz()
	case "Sin":
		impresionXSinMatriz()
	}

}

func impresionDerecha() {

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == j {
				fmt.Print(" 1 ")
			} else {
				fmt.Print(" 0 ")
			}
		}
		fmt.Println()
	}
}

func impresionIzquierda() {
	for c := 0; c < 10; c++ {
		for v := 9; v >= 0; v-- {
			if c == v {
				fmt.Print(" 1 ")
			} else {
				fmt.Print(" 0 ")
			}
		}
		fmt.Println()
	}
}

func impresionXConMatriz() {
	var mat [10][10]int

	for h := 0; h < 10; h++ {
		mat[h][h] = 1
		mat[h][9-h] = 1

	}

	for a := 0; a < 10; a++ {
		for b := 0; b < 10; b++ {
			fmt.Printf(" %v ", mat[a][b])
		}
		fmt.Println()
	}
}

func impresionXSinMatriz() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == j || i+j == 9 {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}

}

// imprimir por consola una diagonal (figura gometrica) del numero 1 que tenga un tamaño de 10 por 10.
