package main

import (
	"fmt"
)

const maxStep = 10

var k, f, a float64
var i, n int
var Knots []float64
var Funcs []float64

func Enter() {
	enterN()
	enterKnot()
	enterFunc()
	enterA()
}

/**
Вводим степень n, до тех пор пока не будет удовлетворять нашим условиям
*/
func enterN() {

	fmt.Println("Введите степень: n <= 10")
	for n <= maxStep {
		if fmt.Scan(&n); n <= maxStep {
			fmt.Println("n была успешно введена")
			break
		} else {
			fmt.Println("Вы ввели неподходящее n")
		}
	}
}

/**
Вводим узлы
*/
func enterKnot() {

	Knots = make([]float64, n, n)

	fmt.Println("Введите узлы:")
	for i = 0; i < n; i++ {
		fmt.Scan(&k)
		Knots[i] = k
	}
}

/**
Вводим значения функции в узлах
*/
func enterFunc() {
	Funcs = make([]float64, n, n)
	fmt.Println("Введите значения функций в узлах:")
	for i = 0; i < n; i++ {
		fmt.Scan(&f)
		Funcs[i] = f
	}

}

func enterA() {
	fmt.Println("Введите точку а:")
	fmt.Scan(&a)
}
