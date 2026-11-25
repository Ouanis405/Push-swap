package main

import (
	"math/rand"
	"testing"

	"push-swap/stack"
)

// Pour plusieurs tailles, on trie des nombres au hasard
// puis on rejoue les instructions : la pile doit être triée.
func TestSolve(t *testing.T) {
	for _, n := range []int{0, 1, 2, 3, 4, 5, 6, 10, 100, 500} {
		for try := 0; try < 20; try++ {
			nums := rand.Perm(n)
			s := &stack.Stacks{A: toRanks(nums)}
			ins := Solve(s)

			check := &stack.Stacks{A: toRanks(nums)}
			for _, i := range ins {
				check.Exec(i)
			}
			if !check.IsSorted() {
				t.Fatalf("n=%d : pas triée avec %v", n, nums)
			}
		}
	}
}
