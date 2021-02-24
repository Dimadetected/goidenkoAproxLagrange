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
	time.Sleep(time.Second * 3)
}
