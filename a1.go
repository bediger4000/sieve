package main

import "fmt"

func main() {
	in := make(chan int)
	found := 0
	go startcounter(in)
OUT:
	for {
		m := <-in
		fmt.Printf("%d\n", m)
		found++
		if found >= 10000 {
			break OUT
		}
		in = startfilter(m, in)
	}
}

func startcounter(o chan int) {
	for n := 2; true; n++ {
		o <- n
	}
}

func startfilter(m int, in chan int) chan int {
	out := make(chan int)
	// anonymous function, read ints from channel in,
	// write ints to channel out if they aren't divisible
	// by prime, which should have the value of a prime number.
	go func(in, out chan int, prime int) {
		for n := range in {
			if (n % prime) != 0 {
				out <- n
			}
		}
	}(in, out, m)
	return out
}
