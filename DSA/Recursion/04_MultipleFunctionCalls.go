package main

func fibonacchi(n int) int {

	if n == 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	return n + fibonacchi(n-1)
}
