package math

import (
	"testing"
)

func TestAdd(t *testing.T) {
	total := Add(2, 2)
	if total != 4 {
		t.Errorf("Sum was incorrect, Actual: %d, Expected: %d", total, 4)
	}
	t.Log("running TestAdd")
}

func TestSub(t *testing.T) {
	total := Subtract(2, 2)
	if total != 0 {
		t.Errorf("Sub was incorrect, Actual: %d, Expected: %d", total, 0)
	}
	t.Log("running TestSubtract")
}

// func TestDivide(t *testing.T) {
// 	total := Divide(4, 2)
// 	if total != float32(2) {
// 		t.Errorf("Divide was incorrect, Actual: %.2f, Expected: %.2f", total, float32(2))
// 	}
// 	t.Log("running TestDivide")
// }
