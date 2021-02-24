package main

import (
	"fmt"
	"time"
)

func main() {
	Enter()
	mainSum()
	fmt.Print("L_", n, "(a) = ", answerSum, "\n")

	enterM()
	fmt.Println("Оценка погрешности: ", errorRating())
	fmt.Println("Оценка погрешности sin(x): ", errorRatingSin())
	time.Sleep(time.Second * 30)
}
