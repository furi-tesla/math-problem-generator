package main

import (
	"fmt"
	"mathproblemgenerator/problem"
)

func main() {
	//rand.Seed(time.Now().UnixNano()) // Seed the rng
	problem.GenerateInColumns(1, 100, 10, "d")

	problem.PrintAnswerKey()

	var n = problem.Number{}

	n.Value = 27

	n.NameNum()

	fmt.Println(n.Name)

}
