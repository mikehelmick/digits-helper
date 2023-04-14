package digits

import (
	"errors"
	"fmt"

	"golang.org/x/exp/slices"
)

type Solver struct {
	target     int
	digits     []int
	operations []string
	solved     bool
}

var (
	ErrorInvalidOp = errors.New("invalid operation")
)

func New(target int, digits []int) (*Solver, error) {
	if target <= 0 || target > 2000 {
		return nil, fmt.Errorf("target must be >0 and <= 2000")
	}
	if l := len(digits); l != 6 {
		return nil, fmt.Errorf("6 digits are required to start a new puzzle, got: %v", l)
	}
	for _, d := range digits {
		if d <= 0 || d >= 200 {
			return nil, fmt.Errorf("starting digits must be > 0 and <= 200, got: %v", d)
		}
	}

	digs := make([]int, len(digits))
	copy(digs, digits)

	slices.SortFunc(digs, func(i, j int) bool {
		return i > j
	})

	return &Solver{
		target:     target,
		digits:     digs,
		operations: make([]string, 0),
	}, nil
}

func (s *Solver) next(a, b int, op Operation) (*Solver, error) {
	if len(s.digits) < 2 {
		return nil, fmt.Errorf("not enough digits to perform operation")
	}
	if !op.Can(a, b) {
		return nil, ErrorInvalidOp
	}

	foundA, foundB := false, false
	newD := make([]int, 0, len(s.digits)-2)
	for _, d := range s.digits {
		if d == a && !foundA {
			foundA = true
			continue
		}
		if d == b && !foundB {
			foundB = true
			continue
		}
		newD = append(newD, d)
	}
	ops := make([]string, len(s.operations))
	copy(ops, s.operations)

	// Do the operation
	res := op.Op(a, b)
	newD = append(newD, res)
	slices.SortFunc(newD, func(i, j int) bool {
		return i > j
	})

	return &Solver{
		target:     s.target,
		digits:     newD,
		operations: append(ops, op.String(a, b)),
		solved:     s.target == res,
	}, nil
}

func (s *Solver) Solve() ([]string, error) {

	for i := 0; i < len(s.digits)-1; i++ {
		for j := i + 1; j < len(s.digits); j++ {
			for _, op := range AllOperations {
				a := s.digits[i]
				b := s.digits[j]

				n, err := s.next(a, b, op)
				if err != nil {
					continue
				}

				if n.solved {
					return n.operations, nil
				}
				// else try to solve from there
				res, err := n.Solve()
				if res == nil {
					continue
				}
				return res, nil
			}
		}
	}

	return nil, fmt.Errorf("no solutions found")
}
