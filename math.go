package main

import "math"

var answerSum float64 = 0

func mainSum() {
	for i = 0; i <= n; i++ {
		answerSum = answerSum + (Funcs[i] * compositionFiX(i))
	}
}

func compositionFiX(i int) float64 {
	var compos float64 = 1
	for j := 0; j <= n; j++ {
		if i != j {
			compos *= (a - Knots[j]) / (Knots[i] - Knots[j])
		}
	}
	return compos
}

//Оценка погрешности
func errorRating() float64 {
	return derivative() / fact(float64(n)+1) * math.Abs(errorRatingCompose())
}

func derivative() float64 {
	return m
}

//Факториал
func fact(a float64) float64 {

	var i, sum float64 = 2, 1
	for i = 2; i <= a; i++ {
		sum *= i
	}
	return sum
}

func errorRatingCompose() float64 {
	var compose float64 = 1
	for i = 0; i <= n; i++ {
		compose *= a - Knots[i]
	}
	return compose
}
