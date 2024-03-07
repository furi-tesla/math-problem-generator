package problem

import (
	"fmt"
	"math/rand"
	"strings"
)

// Problem type stores information about each problem.
type Problem struct {
	NumberA   Number // randomly generated number
	Operation string // +, -, ÷, x
	NumberB   Number // randomly generated number
	Answer    int    // calculated within functions that generate the problem
	Order     int    // used to create and organize the answer key
}

// Number type's names can be generated by the String func below, up to 4 digits
type Number struct {
	Value int
	Name  string
}

var AnswerKey []int // Slice to store answers

// NumIsValid checks if the number is positive
func (a *Number) NumIsValid() bool {
	return a.Value >= 0
}

// IsDivisible checks whether the input numbers (n1 is larger or equal) are evenly divisible
func IsDivisible(n1, n2 int) bool {
	return n1%n2 == 0
}

// String writes out the name for each number, up to 4 digits (might want to create a version just for the int input)
func (a *Number) String() string {
	ones := a.Value % 10

	tens := (a.Value / 10) % 10

	hunds := (a.Value / 100) % 10

	thous := (a.Value / 1000) % 10

	var o string // ones spelled out

	if tens != 1 { // in every case where tens is not one, ones will be spelled out normally
		switch ones {
		case 1:
			o = "one"
		case 2:
			o = "two"
		case 3:
			o = "three"
		case 4:
			o = "four"
		case 5:
			o = "five"
		case 6:
			o = "six"
		case 7:
			o = "seven"
		case 8:
			o = "eight"
		case 9:
			o = "nine"
		default:
		}
	}

	var t string

	switch tens {
	case 0:
		t = ""
	case 1:
		switch ones { // the ones digit is omitted in the case where tens are 1 (teen numbers)
		case 0:
			t = "ten"
		case 1:
			t = "eleven"
		case 2:
			t = "twelve"
		case 3:
			t = "thirteen"
		case 4:
			t = "fourteen"
		case 5:
			t = "fifteen"
		case 6:
			t = "sixteen"
		case 7:
			t = "seventeen"
		case 8:
			t = "eighteen"
		case 9:
			t = "nineteen"
		default:
		}
	case 2:
		t = "twenty"
	case 3:
		t = "thirty"
	case 4:
		t = "forty"
	case 5:
		t = "fifty"
	case 6:
		t = "sixty"
	case 7:
		t = "seventy"
	case 8:
		t = "eighty"
	case 9:
		t = "ninety"
	default:
	}

	var h string

	switch hunds {
	case 1:
		h = "one"
	case 2:
		h = "two"
	case 3:
		h = "three"
	case 4:
		h = "four"
	case 5:
		h = "five"
	case 6:
		h = "six"
	case 7:
		h = "seven"
	case 8:
		h = "eight"
	case 9:
		h = "nine"
	default:
	}

	if hunds != 0 {
		h += " hundred"
	}

	var th string

	switch thous {
	case 1:
		th = "one"
	case 2:
		th = "two"
	case 3:
		th = "three"
	case 4:
		th = "four"
	case 5:
		th = "five"
	case 6:
		th = "six"
	case 7:
		th = "seven"
	case 8:
		th = "eight"
	case 9:
		th = "nine"
	default:
	}

	if thous != 0 {
		th += " thousand"
	}
	var numName string

	// a little weird concat below, but hopefully I can fix it later
	if th != "" {
		if h != "" {
			if t != "" {
				numName = th + " " + h + " " + t + " " + o
			} else {
				numName = th + " " + h + " " + o
			}
		} else {
			if t != "" {
				numName = th + " " + t + " " + o
			} else {
				numName = th + " " + o
			}
		}
	} else if h != "" {
		if t != "" {
			numName = h + " " + t + " " + o
		} else {
			numName = h + " " + o
		}
	} else {
		if t != "" {
			numName = t + " " + o
		} else {
			numName = o
		}
	}

	a.Name = numName

	return a.Name
}

func Generate(min, max, num int, operation string) {

	p := Problem{}

	op := strings.ToLower(operation)

	var opSign string

	switch op {
	case "m", "multiply", "times", "multiplication":
		opSign = "x"
	case "d", "divide", "divided by", "division":
		opSign = "÷"
	case "a", "add", "plus", "addition":
		opSign = "+"
	case "s", "minus", "subtract", "subtraction":
		opSign = "-"
	default:
		fmt.Println("Please enter a valid operation. You may enter 'm' for multiplication, 'd' for division, 'a' to add, and 's' to subtract.")
		return // exit the function if the operation is invalid
	}

	p.Operation = opSign

	for i := 1; i <= num; { // increment is provided below only when the problem is successfully generated
		p.Order = i

		r1 := rand.Intn(max-min+1) + min // no need to seed random numbers anymore
		r2 := rand.Intn(max-min+1) + min

		if opSign == "+" || opSign == "x" {
			fmt.Printf("%d. %d %s %d =\n", i, r1, opSign, r2)
			p.NumberA.Value = r1
			p.NumberB.Value = r2

			if opSign == "+" {
				p.Answer = r1 + r2
			} else if opSign == "x" {
				p.Answer = r1 * r2
			}
			AnswerKey = append(AnswerKey, p.Answer)
			i++
		} else if opSign == "-" {
			if r1 >= r2 {
				fmt.Printf("%d. %d %s %d =\n", i, r1, opSign, r2)

				p.NumberA.Value = r1
				p.NumberB.Value = r2
				p.Answer = r1 - r2

			} else {
				fmt.Printf("%d. %d %s %d =\n", i, r2, opSign, r1)

				p.NumberA.Value = r2
				p.NumberB.Value = r1
				p.Answer = r2 - r1
			}
			AnswerKey = append(AnswerKey, p.Answer)
			i++
		} else if opSign == "÷" {
			if r1 >= r2 {
				if IsDivisible(r1, r2) {
					fmt.Printf("%d. %d %s %d =\n", i, r1, opSign, r2)

					p.NumberA.Value = r1
					p.NumberB.Value = r2
					p.Answer = r1 / r2
					AnswerKey = append(AnswerKey, p.Answer)
					i++
				}
			}
		}
	}
}

func PrintAnswerKey() {
	fmt.Println("Answer Key:")
	for order, answer := range AnswerKey {
		fmt.Printf("%d. %d", order+1, answer)
		fmt.Println()
	}
}

func GenerateInColumns(min, max, num int, operation string) {

	p := Problem{}

	op := strings.ToLower(operation)

	var opSign string

	switch op {
	case "m", "multiply", "times", "multiplication":
		opSign = "x"
	case "d", "divide", "divided by", "division":
		opSign = "÷"
	case "a", "add", "plus", "addition":
		opSign = "+"
	case "s", "minus", "subtract", "subtraction":
		opSign = "-"
	default:
		fmt.Println("Please enter a valid operation. You may enter 'm' for multiplication, 'd' for division, 'a' to add, and 's' to subtract.")
		return // exit the function if the operation is invalid
	}

	p.Operation = opSign

	for i := 1; i <= num; { // increment is provided below only when the problem is successfully generated
		p.Order = i

		r1 := rand.Intn(max-min+1) + min // no need to seed random numbers anymore
		r2 := rand.Intn(max-min+1) + min

		if opSign == "+" || opSign == "x" {
			if i < 10 {
				fmt.Printf("%d. %6d\n%4s   %2d\n━━━━━━━━━\n", i, r1, opSign, r2)
			} else if i >= 10 {
				fmt.Printf("%d. %5d\n%4s   %2d\n━━━━━━━━━\n", i, r1, opSign, r2)
			}
			p.NumberA.Value = r1
			p.NumberB.Value = r2

			if opSign == "+" {
				p.Answer = r1 + r2
			} else if opSign == "x" {
				p.Answer = r1 * r2
			}
			AnswerKey = append(AnswerKey, p.Answer)
			i++
		} else if opSign == "-" {
			if r1 >= r2 {
				if i < 10 {
					fmt.Printf("%d. %6d\n%4s   %2d\n━━━━━━━━━\n", i, r1, opSign, r2)
				} else if i >= 10 {
					fmt.Printf("%d. %5d\n%4s   %2d\n━━━━━━━━━\n", i, r1, opSign, r2)
				}

				p.NumberA.Value = r1
				p.NumberB.Value = r2
				p.Answer = r1 - r2

			} else {
				if i < 10 {
					fmt.Printf("%d. %6d\n%4s   %2d\n━━━━━━━━━\n", i, r2, opSign, r1)
				} else if i >= 10 {
					fmt.Printf("%d. %5d\n%4s   %2d\n━━━━━━━━━\n", i, r2, opSign, r1)
				}

				p.NumberA.Value = r2
				p.NumberB.Value = r1
				p.Answer = r2 - r1
			}
			AnswerKey = append(AnswerKey, p.Answer)
			i++
		} else if opSign == "÷" {
			if r1 >= r2 {
				if IsDivisible(r1, r2) {
					if i < 10 {
						if r2 < 10 {
							fmt.Printf("%d. ______\n  %d│%5d\n", i, r2, r1)
						} else {
							fmt.Printf("%d. ______\n %d│%5d\n", i, r2, r1)
						}
					} else if i >= 10 {
						if r2 < 10 {
							fmt.Printf("%d. ______\n  %d│%5d\n", i, r2, r1)
						} else {
							fmt.Printf("%d. ______\n  %d│%6d\n", i, r2, r1)
						}
					}

					p.NumberA.Value = r1
					p.NumberB.Value = r2
					p.Answer = r1 / r2
					AnswerKey = append(AnswerKey, p.Answer)
					i++
				}
			}
		}
	}
}
