package main

import "fmt"

func main() {
	in := make(chan int)
	startcounter(in)
	found := 0
OUT:
	for {
		var m int
		select {
		case m = <-in:
			fmt.Printf("%d\n", m)
			in = start(m, in)
			found++
			if found >= 10000 {
				break OUT
			}
		}
	}
}

func startcounter(o chan int) {
	go func(o chan int) {
		for n := 2; true; n++ {
			o <- n
		}
	}(o)
}

func start(m int, in chan int) (out chan int) {
	out = make(chan int)

	go func(in, out chan int, m int) {
		for n := range in {
			x := n % m
			if x == 0 {
				continue
			}
			out <- n
		}
	}(in, out, m)

	return
}
