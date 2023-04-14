package digits

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewBadTarget(t *testing.T) {
	t.Parallel()

	_, err := New(-1, []int{})
	if err == nil {
		t.Fatalf("Expected error, got nil")
	} else if !strings.Contains(err.Error(), "target") {
		t.Fatalf("Wrong error, doesn't reference target.")
	}
}

func TestNewBadDigits(t *testing.T) {
	t.Parallel()

	_, err := New(10, []int{1})
	if err == nil {
		t.Fatalf("expected error, got nil")
	} else if !strings.Contains(err.Error(), "6 digits are required") {
		t.Fatalf("Wrong error, doesn't reference '6 digits are required': got: %v", err.Error())
	}

	_, err = New(10, []int{1, 2, 3, 4, 5, -1})
	if err == nil {
		t.Fatalf("expected error, got nil")
	} else if !strings.Contains(err.Error(), "starting digits") {
		t.Fatalf("Wrong error, doesn't reference 'starting digits': got: %v", err.Error())
	}
}

func TestNew(t *testing.T) {
	t.Parallel()

	got, err := New(59, []int{2, 3, 5, 11, 15, 25})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := &Solver{
		target:     59,
		digits:     []int{25, 15, 11, 5, 3, 2},
		operations: []string{},
	}

	if diff := cmp.Diff(want, got, cmp.AllowUnexported(Solver{})); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestNext(t *testing.T) {
	t.Parallel()

	target := 59
	digits := []int{2, 3, 5, 11, 15, 25}

	cases := []struct {
		Name          string
		A             int
		B             int
		Op            Operation
		WantDigits    []int
		WantOperation string
	}{
		{
			Name:          "Multiplication",
			A:             25,
			B:             2,
			Op:            OpMultiplication,
			WantDigits:    []int{50, 15, 11, 5, 3},
			WantOperation: "25 * 2",
		},
		{
			Name:          "Addition",
			A:             25,
			B:             2,
			Op:            OpAddition,
			WantDigits:    []int{27, 15, 11, 5, 3},
			WantOperation: "25 + 2",
		},
		{
			Name:          "Subtraction",
			A:             25,
			B:             2,
			Op:            OpSubtraction,
			WantDigits:    []int{23, 15, 11, 5, 3},
			WantOperation: "25 - 2",
		},
		{
			Name:          "Division",
			A:             15,
			B:             3,
			Op:            OpDivision,
			WantDigits:    []int{25, 11, 5, 5, 2},
			WantOperation: "15 / 3",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			s, _ := New(target, digits)

			got, err := s.next(c.A, c.B, c.Op)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			want := &Solver{
				target:     target,
				digits:     c.WantDigits,
				operations: []string{c.WantOperation},
			}

			if diff := cmp.Diff(want, got, cmp.AllowUnexported(Solver{})); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestSol(t *testing.T) {
	t.Parallel()

	s, err := New(59, []int{2, 3, 5, 11, 15, 25})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got, err := s.Solve()

	want := []string{
		"25 + 15",
		"40 + 11",
		"51 + 5",
		"56 + 3",
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}
