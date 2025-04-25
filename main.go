package main

import "fmt"

func collatz(n int) int {
	if n%2 == 0 {
		n = n / 2
		fmt.Println(n)
		return n
	} else {
		n = (n * 3) + 1
		fmt.Println(n)
		return n
	}
}

func main() {
	fmt.Println("Collatz Conjecture")

	var input, steps int

	fmt.Print("Input any positive number:")
	fmt.Scan(&input)

	//main collatz loop
	for i := 0; input != 1; i++ {
		input = collatz(input)
		steps = steps + 1
		if input == 1 {
			fmt.Println("steps count:", steps)
		}
	}

}

