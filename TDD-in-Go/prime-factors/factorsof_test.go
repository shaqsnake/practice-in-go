package primefactors

import (
	"testing"
	"fmt"
	"reflect"
)
func TestFactorsOf(t *testing.T) {
	type checkFunc func([]int) error

	// SECTION 1: checkers
	isEmptyList := func(have []int) error {
		if len(have) > 0 {
			return fmt.Errorf("Expected empty list, found %v.", have)
		}
		return nil
	}

	is := func(want ...int) checkFunc {
		return func(have []int) error {
			if !reflect.DeepEqual(have, want) {
				return fmt.Errorf("Expect list %v, found %v.", want, have)
			}
			return nil
		}
	}

	// SECTION 2: test cases
	tests := [...]struct {
		in int
		check checkFunc
	}{
		// factors of n
		{1, isEmptyList},
		{2, is(2)},
	}

	// SECTION 3: test logic
	for _, tc := range tests {
		t.Run(fmt.Sprintf("Factors of %d", tc.in), func(t *testing.T) {
			factors := factorsOf(tc.in)
			if err := tc.check(factors); err != nil {
				t.Error(err)
			}
		})
	}
}