package main

import (
	"fmt"
	"mathproblemgenerator/problem"
)

func main() {

	problem.Generate(1, 100, 10, "d")

	problem.PrintAnswerKey()

	var n = problem.Number{}

	n.Value = 27

	n.String()

	fmt.Println(n.Name)

}
