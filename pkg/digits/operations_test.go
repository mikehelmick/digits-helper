package digits

import "testing"

func TestAddition(t *testing.T) {
	op := &Addition{}
	if !op.Can(1, 2) {
		t.Fatalf("addition always can...")
	}

	res := op.Op(1, 2)
	if res != 3 {
		t.Fatalf("addition got wrong result 1+2 = %v", res)
	}
}

func TestSubtraction(t *testing.T) {
	op := &Subtraction{}
	if !op.Can(10, 6) {
		t.Fatalf("subtraction always can...")
	}

	res := op.Op(10, 6)
	if res != 4 {
		t.Fatalf("subtraction got wrong result 10-6 = %v", res)
	}
}

func TestMultiplication(t *testing.T) {
	op := &Multiplication{}
	if !op.Can(10, 6) {
		t.Fatalf("multiplication always can...")
	}

	res := op.Op(10, 6)
	if res != 60 {
		t.Fatalf("subtraction got wrong result 10*6 = %v", res)
	}
}

func TestDivision(t *testing.T) {
	op := &Division{}
	if op.Can(10, 3) {
		t.Fatalf("division shouldn't be able to 10/3")
	}

	if !op.Can(9, 3) {
		t.Fatalf("division should be able to 9/3")
	}

	res := op.Op(9, 3)
	if res != 3 {
		t.Fatalf("division got wrong result 9/3 = %v", res)
	}
}
