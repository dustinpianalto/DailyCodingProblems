package main

import "log"

func cons(a, b int) func(func(int, int) (int, int)) (int, int) {
	return func(f func(int, int) (int, int)) (int, int) {
		return f(a, b)
	}
}

func test(a, b int) (int, int) {
	return a, b
}

func cdr(f func(func(int, int) (int, int)) (int, int)) int {
	_, b := f(test)
	return b
}

func car(f func(func(int, int) (int, int)) (int, int)) int {
	a, _ := f(test)
	return a
}

func main() {
	log.Println(car(cons(3, 4)))
	log.Println(cdr(cons(3, 4)))
}
