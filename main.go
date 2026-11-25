package main

import (
	"fmt"
	"os"
	"sort"

	"push-swap/stack"
)

func main() {
	// Pas d'argument : on n'affiche rien.
	if len(os.Args) < 2 {
		return
	}

	nums, err := stack.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	// On remplace chaque nombre par son rang (0 = le plus petit).
	// Trier les rangs revient à trier les nombres, mais c'est plus simple.
	s := &stack.Stacks{A: toRanks(nums)}

	for _, ins := range Solve(s) {
		fmt.Println(ins)
	}
}

// toRanks remplace chaque valeur par sa position dans la liste triée.
func toRanks(nums []int) []int {
	sorted := append([]int{}, nums...)
	sort.Ints(sorted)

	rank := map[int]int{}
	for i, n := range sorted {
		rank[n] = i
	}

	ranks := make([]int, len(nums))
	for i, n := range nums {
		ranks[i] = rank[n]
	}
	return ranks
}
