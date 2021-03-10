# Daily Coding Problem: Problem #677 [Easy]

This problem was asked by Square.

The Sieve of Eratosthenes is an algorithm used to generate all prime
numbers smaller than N.
The method is to take increasingly larger prime numbers,
and mark their multiples as composite.

For example,
to find all primes less than 100,
we would first mark [4, 6, 8, ...] (multiples of two),
then [6, 9, 12, ...] (multiples of three), and so on.
Once we have done this for all primes less than N,
the unmarked numbers that remain will be prime.

Implement this algorithm.

Bonus: Create a generator that produces primes indefinitely
(that is, without taking N as an input).

## Build and Run

[Code](a1.go)

```sh
$ go build a1.go
$ ./a1 | more
```

The program produces an endless stream of text representations
of prime numbers on stdout.

## Analysis

I wrote a program in Go that runs one goroutine per prime number
that it finds.
It creates a `chan int`, the input channel,
and a goroutine that produces integers from 2 to 9223372036854775807.
The main goroutine reads ints from the channel.
Each of these ints should be a prime number.
For each int it reads,
it starts a goroutine
with an integer read from the channel,
an the input channel as arguments.
Each of these goroutines reads integers from their input channels,
only passing along those integers that aren't divisible
by the input argument.
That is, each goroutine so created sieves out
numbers that aren't divisible by some particular integer.
The starting number from the input channel is 2,
the smallest prime number.
The main goroutine creates a goroutine that only passes input
numbers that aren't divisible by 2.
It weeds out every other number.
The next number is 3, which gets passed by the divisible-by-2
goroutine, so the main goroutine creates a divisible-by-3
sieving goroutine,
and so forth.

I don't think this was easy:
passing channels into goroutines,
and reading from channels returned when starting
the groutines,
was tricky.
Naturally, I had data races
in setting the variable containing the channel
to read primes from.
I think I should have passed in a new channel
from which to read primes,
rather than having the start-a-filter func
create it.

On second thought,
I'm not sure this meets the problem statement.
The statement implies that an array of integers exists,
the all the multiples of 2 get struck,
then all the multiples of 3,
then all the multiples of 5,
etc etc.
This program generates a stream of integers,
and independent worker routines throw out multiples
of a single number.
The program adds to the sieve as it finds previously-unknow
primes.

I found [another version](a2.go) of this I did a few years ago
while watching a
[Rob Pike youtube video](https://www.youtube.com/watch?v=hB05UFqOtFA&feature=youtu.be).
Pike talked about a research language he wrote called
[Newsqueak](https://en.wikipedia.org/wiki/Newsqueak).
I transposed his Newsqueak code into Go to get this,
then forgot about it until now.

## Interview analysis

How in the heck is this "easy"?

This would be a decent problem for a mid- to senior-level
job applicant.
There's an algorithm (Sieve of Eratosthenes)
that can be implemented by careful threading.
There's some finicky keeping track of channels,
and in its Golang form,
the easy potential for creating data races.
