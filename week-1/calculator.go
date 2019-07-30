package calculator

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "errors"
    "strconv"
    "log"
)

func parse(text string) (a float64, b float64, op string, err error) {
    expr := strings.Fields(text)
    if len(expr) != 3 {
        err = errors.New("Invalid format")
        return
    }
    // fmt.Println(expr)

    op = expr[1]

	a, err = strconv.ParseFloat(expr[0], 10)
	if err != nil {
		return
	}

	b, err = strconv.ParseFloat(expr[2], 10)

	if err != nil {
		return
	}

    return a, b, op, err
}

type Calculator interface {
    Add(x, y float64) float64
    Subtract(x, y float64) float64
    Multiply(x, y float64) float64
    Divide(x, y float64) float64
}

type calc struct {}

func newCalculator() Calculator {
	return calc{}
}

func(c calc) Add(x float64, y float64) float64 {
    return x + y
}

func(c calc) Subtract(x float64, y float64) float64 {
    return x - y
}

func(c calc) Multiply(x float64, y float64) float64 {
    return x * y
}

func(c calc) Divide(x float64, y float64) float64 {
    return x / y
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("enter your formula : ")
    text, _ := reader.ReadString('\n')
    // fmt.Print(text)
    a, b, op, _ := parse(text)

    calc := newCalculator()

    var result float64
    switch op {
    case "+":
       result = calc.Add(a,b)
    case "-":
       result = calc.Subtract(a,b)
    case "*":
       result = calc.Multiply(a,b)
    case "/":
       result = calc.Divide(a,b)
	default:
		log.Fatal("Invalid operator. Allowed operators: + - * /")
    }

    fmt.Println(result)
}
