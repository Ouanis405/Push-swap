package main

import "push-swap/stack"

// Solve trie la pile a (qui contient des rangs 0..n-1)
// et renvoie la liste des instructions utilisées.
func Solve(s *stack.Stacks) []string {
	var out []string

	// do exécute l'instruction sur les piles ET la note dans out.
	do := func(ins string) {
		s.Exec(ins)
		out = append(out, ins)
	}

	n := len(s.A)
	switch {
	case s.IsSorted():
		// rien à faire
	case n <= 5:
		sortSmall(s, do)
	default:
		sortBig(s, do)
	}
	return out
}

// sortSmall (2 à 5 nombres) : on envoie les plus petits dans b
// jusqu'à ce qu'il en reste 3, on trie ces 3, puis on remet b dans a.
func sortSmall(s *stack.Stacks, do func(string)) {
	for len(s.A) > 3 {
		rotateToTop(s.A, minIndex(s.A), "ra", "rra", do)
		do("pb")
	}
	sortThree(s, do)
	for len(s.B) > 0 {
		do("pa")
	}
}

// sortThree trie 2 ou 3 nombres en regardant les cas un par un.
func sortThree(s *stack.Stacks, do func(string)) {
	a := s.A
	if len(a) == 2 {
		if a[0] > a[1] {
			do("sa")
		}
		return
	}
	x, y, z := a[0], a[1], a[2]
	switch {
	case x > y && y > z: // 2 1 0
		do("sa")
		do("rra")
	case x > y && x < z: // 1 0 2
		do("sa")
	case x > y: // 2 0 1
		do("ra")
	case x < z && y > z: // 0 2 1
		do("sa")
		do("ra")
	case y > z: // 1 2 0
		do("rra")
	}
}

// sortBig (plus de 5 nombres) : algorithme par "tranches".
//  1. On envoie dans b les nombres par tranches de la taille "chunk"
//     (d'abord les chunk plus petits, puis les chunk suivants, etc.).
//  2. On ramène dans a toujours le plus grand de b : a se remplit du plus
//     grand au plus petit, donc à la fin a est triée.
func sortBig(s *stack.Stacks, do func(string)) {
	n := len(s.A)
	chunk := 15
	if n > 100 {
		chunk = 30
	}

	// Étape 1 : a -> b
	pushed := 0
	for len(s.A) > 0 {
		top := s.A[0]
		switch {
		case top < pushed: // petit nombre : on le range tout en bas de b
			do("pb")
			do("rb")
			pushed++
		case top < pushed+chunk: // nombre de la tranche : il reste en haut de b
			do("pb")
			pushed++
		default: // trop grand pour l'instant : on le laisse passer
			do("ra")
		}
	}

	// Étape 2 : b -> a, toujours le plus grand en premier.
	// Les rangs restants dans b sont 0..len(b)-1, donc le plus grand est len(b)-1.
	for len(s.B) > 0 {
		rotateToTop(s.B, maxIndex(s.B), "rb", "rrb", do)
		do("pa")
	}
}

// rotateToTop amène l'élément d'index i au sommet par le chemin le plus court :
// vers le haut (rot) si il est dans la 1re moitié, vers le bas (rev) sinon.
func rotateToTop(x []int, i int, rot, rev string, do func(string)) {
	if i <= len(x)/2 {
		for k := 0; k < i; k++ {
			do(rot)
		}
	} else {
		for k := 0; k < len(x)-i; k++ {
			do(rev)
		}
	}
}

func minIndex(x []int) int {
	best := 0
	for i, v := range x {
		if v < x[best] {
			best = i
		}
	}
	return best
}

func maxIndex(x []int) int {
	best := 0
	for i, v := range x {
		if v > x[best] {
			best = i
		}
	}
	return best
}
