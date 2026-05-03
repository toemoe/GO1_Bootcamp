package main

import (
	"flag"
	"fmt"
)

func generator(k, n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := k; i <= n; i++ {
			ch <- i
		}
	}()
	return ch
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for val := range in {
			out <- val * val
		}
	}()
	return out
}

func main() {
	k := flag.Int("K", 1, "Number first")
	n := flag.Int("N", 5, "Number second")
	flag.Parse()
	if *k > *n {
		fmt.Println("error: N must be greater than K")
		return
	}
	gen := generator(*k, *n)

	for val := range square(gen) {
		fmt.Println(val)
	}
}
