package mymath

import (
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	println("Tests are about to run")
	result := m.Run()
	println("Tests done executing")

	os.Exit(result)
}

var addTable = []struct {
	in  []int
	out int
}{
	{[]int{1, 2}, 3},
	{[]int{1, 2, 3, 4}, 10},
}

func TestCanAddNumbers(t *testing.T) {
	t.Parallel()
	// time.Sleep(1 * time.Second)

	for _, entry := range addTable {
		result := Add(entry.in...)
		if result != entry.out {
			t.Error("Failed to add numbers")
		}
	}

	// result := Add(1, 2)
	// if result != 3 {
	// 	t.Log("Failed to add one and two")
	// 	t.Fail()
	// }

	// result = Add(1, 2, 3, 4)
	// if result != 10 {
	// 	t.Error("Failed to add more than two numbers")
	// }
}

func TestCanSubstractNumbers(t *testing.T) {
	t.Parallel()
	time.Sleep(1 * time.Second)

	result := Subtract(10, 5)
	if result != 5 {
		t.Error("Failed to substract two numbers")
	}
}

// func TestCanMultiplyNumbers(t *testing.T) {
// 	t.Skip("Not implemented yet")
// }
