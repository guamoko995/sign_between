package main

import (
	"testing"
)

func TestPow(t *testing.T) {
	// 0 ^ 0 = 1
	added := Pow(0, 0)
	exp := 1
	if added != exp {
		t.Errorf("Expected: %v, added: %v.", exp, added)
	}

	// 0 ^ 1 = 0
	added = Pow(0, 1)
	exp = 0
	if added != exp {
		t.Errorf("Expected: %v, added: %v.", exp, added)
	}

	// 2 ^ 10 = 1024
	added = Pow(2, 10)
	exp = 1024
	if added != exp {
		t.Errorf("Expected: %v, added: %v.", exp, added)
	}
}

func TestReadeOperand(t *testing.T) {
	// 123
	// ^
	addedOperand, addedCursore := ReadeOperand(0, []int{1, 2, 3}, 0)
	expOperand, expCursore := 123, 3
	if addedOperand != expOperand {
		t.Errorf("Expected: %v, added: %v.", expOperand, addedOperand)
	}
	if addedCursore != expCursore {
		t.Errorf("Expected: %v, added: %v.", expCursore, addedCursore)
	}

	// 1+23 stratege = 1*3^0 = 1
	// ^ *
	// 0 1
	addedOperand, addedCursore = ReadeOperand(1, []int{1, 2, 3}, 0)
	expOperand, expCursore = 1, 1
	if addedOperand != expOperand {
		t.Errorf("Expected: %v, added: %v.", expOperand, addedOperand)
	}
	if addedCursore != expCursore {
		t.Errorf("Expected: %v, added: %v.", expCursore, addedCursore)
	}

	// 12+3 : stratege = 1*3^1 = 3
	// ^  ^
	// 0  2
	addedOperand, addedCursore = ReadeOperand(3, []int{1, 2, 3}, 0)
	expOperand, expCursore = 12, 2
	if addedOperand != expOperand {
		t.Errorf("Expected: %v, added: %v.", expOperand, addedOperand)
	}
	if addedCursore != expCursore {
		t.Errorf("Expected: %v, added: %v.", expCursore, addedCursore)
	}

	// 12+3 : stratege = 1*3^1 = 3
	//  ^ *
	//  1 2
	addedOperand, addedCursore = ReadeOperand(3, []int{1, 2, 3}, 1)
	expOperand, expCursore = 2, 2
	if addedOperand != expOperand {
		t.Errorf("Expected: %v, added: %v.", expOperand, addedOperand)
	}
	if addedCursore != expCursore {
		t.Errorf("Expected: %v, added: %v.", expCursore, addedCursore)
	}

	// 12-346+5 : stratege = 2*3^1 + 1*3^4 = 87
	//     ^  *
	//     3  5
	addedOperand, addedCursore = ReadeOperand(87, []int{1, 2, 3, 4, 6, 5}, 3)
	expOperand, expCursore = 46, 5
	if addedOperand != expOperand {
		t.Errorf("Expected: %v, added: %v.", expOperand, addedOperand)
	}
	if addedCursore != expCursore {
		t.Errorf("Expected: %v, added: %v.", expCursore, addedCursore)
	}

	// 12-346+5   : stratege = 2*3^1 + 1*3^4 = 87
	//        ^ *
	//        5 6
	addedOperand, addedCursore = ReadeOperand(87, []int{1, 2, 3, 4, 6, 5}, 5)
	expOperand, expCursore = 5, 6
	if addedOperand != expOperand {
		t.Errorf("Expected: %v, added: %v.", expOperand, addedOperand)
	}
	if addedCursore != expCursore {
		t.Errorf("Expected: %v, added: %v.", expCursore, addedCursore)
	}
}

func TestReadeOperator(t *testing.T) {
	// 12345673458_9
	//            ^
	//            0
	addedOperator := ReadeOperator(0, 10)
	expOperator := byte(0) // ()
	if addedOperator != expOperator {
		t.Errorf("Expected: %v, added: %v.", expOperator, addedOperator)
	}

	// 1+23 stratege = 1*3^0 = 1
	//  ^
	//  0
	addedOperator = ReadeOperator(1, 0)
	expOperator = byte(1) // (+)
	if addedOperator != expOperator {
		t.Errorf("Expected: %v, added: %v.", expOperator, addedOperator)
	}

	// 1_2+3 : stratege = 1*3^1 = 3
	//  ^
	//  0
	addedOperator = ReadeOperator(3, 0)
	expOperator = byte(0) // ()
	if addedOperator != expOperator {
		t.Errorf("Expected: %v, added: %v.", expOperator, addedOperator)
	}

	// 12+3 : stratege = 1*3^1 = 3
	//   ^
	//   1
	addedOperator = ReadeOperator(3, 1)
	expOperator = byte(1) // (+)
	if addedOperator != expOperator {
		t.Errorf("Expected: %v, added: %v.", expOperator, addedOperator)
	}

	// 12-346+5 : stratege = 2*3^1 + 1*3^4 = 87
	//   ^
	//   1
	addedOperator = ReadeOperator(87, 1)
	expOperator = byte(2) // (-)
	if addedOperator != expOperator {
		t.Errorf("Expected: %v, added: %v.", expOperator, addedOperator)
	}

	// 12-346+5 : stratege = 2*3^1 + 1*3^4 = 87
	//       ^
	//       4
	addedOperator = ReadeOperator(87, 4)
	expOperator = byte(1) //(+)
	if addedOperator != expOperator {
		t.Errorf("Expected: %v, added: %v.", expOperator, addedOperator)
	}
}

func TestSolve(t *testing.T) {
	// #Test 1
	// Создаем выражение 1 2 3 6=18
	s := &expression{
		parameters: []int{1, 2, 3, 6},
		value:      18,
		strateges:  make([]int, 0),
	}

	// Решение: 1+23-6=18
	//           ^  ^
	//           0  2
	// strategy = 1*3^0 + 2*3^2 = 19
	//            +   ^   -   ^
	expSolve := 19

	// Подбор стратегии
	s.Solve()
	addedSolve := s.strateges[0]

	if !s.solvable {
		t.Errorf("Expected: True, added: False.")
	}

	if addedSolve != expSolve {
		t.Errorf("Expected: %v, added: %v.", expSolve, addedSolve)
	}

	// #Test 2
	// Создаем выражение 1 2 3 6=18
	s = &expression{
		parameters: []int{1, 2},
		value:      12,
		strateges:  make([]int, 0),
	}

	// Решение: 12=12
	// strategy = 0
	expSolve = 0

	// Подбор стратегии
	s.Solve()
	addedSolve = s.strateges[0]

	if !s.solvable {
		t.Errorf("Expected: True, added: False.")
	}

	if addedSolve != expSolve {
		t.Errorf("Expected: %v, added: %v.", expSolve, addedSolve)
	}

	// #Test 3
	// Создаем выражение 3 2=1
	s = &expression{
		parameters: []int{3, 2},
		value:      1,
		strateges:  make([]int, 0),
	}

	// Решение: 3-2=1
	//           ^
	//           0
	// strategy = 2*3^0  = 2
	//            -   ^
	expSolve = 2

	// Подбор стратегии
	s.Solve()
	addedSolve = s.strateges[0]

	if !s.solvable {
		t.Errorf("Expected: True, added: False.")
	}

	if addedSolve != expSolve {
		t.Errorf("Expected: %v, added: %v.", expSolve, addedSolve)
	}

	// #Test 4
	// Создаем выражение 3 2=4
	s = &expression{
		parameters: []int{3, 2},
		value:      4,
		strateges:  make([]int, 0),
	}

	// Решения нет
	// strategy = 0 (значение по умолчанию)
	expSolve = 0

	// Подбор стратегии
	s.Solve()

	if s.solvable {
		t.Errorf("Expected: False, added: True.")
	}
}

func TestString(t *testing.T) {
	// #Test 1
	// Создаем выражение 1-23+65=43
	//                    ^  ^
	//                    0  2
	// strategy = 2*3^0 + 1*3^2  = 11
	//            -   ^   +   ^

	s := &expression{
		parameters: []int{1, 2, 3, 6, 5},
		value:      43,
		strateges:  []int{11},
		solvable:   true,
	}

	exp := "1-23+65=43\n"
	added := s.String()

	if added != exp {
		t.Errorf("Expected: %v, added: %v.", exp, added)
	}
}
