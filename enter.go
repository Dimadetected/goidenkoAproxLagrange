package main

import (
	"fmt"
	"math"
)

const maxStep = 10

var k, f, a, m float64
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

	Knots = make([]float64, n+1, n+1)

	fmt.Print("Введите узлы:")
	for i = 0; i <= n; i++ {
		fmt.Scan(&k)
		Knots[i] = k
	}
}

func testF1(x float64) float64 {
	return x * x * x
}
func testF2(x float64) float64 {
	return math.Sin(x)
}

/**
Вводим значения функции в узлах
*/
func enterFunc() {
	Funcs = make([]float64, n+1, n+1)
	fmt.Println("Значения функций в узлах:")
	for i = 0; i <= n; i++ {
		//fmt.Scan(&f)
		//Funcs[i] = testF1(Knots[i])
		Funcs[i] = testF2(Knots[i])
	}
	fmt.Println(Funcs)
}

func enterA() {
	fmt.Print("Введите точку а:")
	fmt.Scan(&a)
}

func enterM() {
	fmt.Print("Введите константу для оценки n:")
	fmt.Scan(&m)
}
