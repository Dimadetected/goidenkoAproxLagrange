package main

var answerSum float64 = 0

func mainSum() {
	for i = 0; i < n; i++ {
		answerSum += (Funcs[i] * compositionFiX(i))
	}
}

func compositionFiX(i int) float64 {
	var compos float64 = 1
	for j := 0; j < n; j++ {
		if i != j {
			compos *= (a - Knots[j]) / (Knots[i] - Knots[j])
		}
	}
	return compos
}
