package main

/*
https://www.youtube.com/watch?v=hB05UFqOtFA&feature=youtu.be

"Advanced Topics in Programming Languages: Concurrency/message passing
Newsqueak", 2012 Rob Pike talk.
*/

import "fmt"

func main() {
	prime := make(chan int64)
	go sieve(prime)
	for i := 0; true; i++ {
		n := <-prime
		fmt.Printf("%d\t%d\n", i, n)
	}
}

func counter(c chan int64) {
	var i int64 = 1
	for {
		i++
		c <- i
	}
}

func filter(prime int64, recv chan int64, send chan int64) {
	for {
		i := <-recv
		if i%prime != 0 {
			// fmt.Printf("filter for %d passing along %d\n", prime, i)
			send <- i
		}
	}
}

func sieve(prime chan int64) {
	c := make(chan int64)
	go counter(c)
	for {
		p := <-c
		prime <- p
		newc := make(chan int64)
		go filter(p, c, newc)
		c = newc
	}
}
