package digits

import "fmt"

type Operation interface {
	Can(a, b int) bool
	Op(a, b int) int
	String(a, b int) string
}

var (
	AllOperations = []Operation{
		&Addition{},
		&Subtraction{},
		&Multiplication{},
		&Division{},
	}

	OpAddition       = &Addition{}
	OpSubtraction    = &Subtraction{}
	OpMultiplication = &Multiplication{}
	OpDivision       = &Division{}
)

type Addition struct{}

func (op *Addition) Can(a, b int) bool { return true }

func (op *Addition) Op(a, b int) int {
	return a + b
}

func (op *Addition) String(a, b int) string {
	return fmt.Sprintf("%v + %v", a, b)
}

type Subtraction struct{}

func (op *Subtraction) Can(a, b int) bool { return true }

func (op *Subtraction) Op(a, b int) int {
	return a - b
}

func (op *Subtraction) String(a, b int) string {
	return fmt.Sprintf("%v - %v", a, b)
}

type Multiplication struct{}

func (op *Multiplication) Can(a, b int) bool { return true }

func (op *Multiplication) Op(a, b int) int {
	return a * b
}

func (op *Multiplication) String(a, b int) string {
	return fmt.Sprintf("%v * %v", a, b)
}

type Division struct{}

func (op *Division) Can(a, b int) bool {
	if b == 0 {
		return false
	}
	return a%b == 0
}

func (op *Division) Op(a, b int) int {
	return a / b
}

func (op *Division) String(a, b int) string {
	return fmt.Sprintf("%v / %v", a, b)
}
