package main

import "fmt"

func generator(nums []int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, num := range nums {
			ch <- num
		}
	}()

	return ch
}

func sqr(ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for c := range ch {
			out <- c * c
		}
	}()

	return out
}

func sumFunc(ch <-chan int) int {
	sum := 0
	for c := range ch {
		sum += c
	}

	return sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	gen := generator(nums)
	sqrCh := sqr(gen)
	sum := sumFunc(sqrCh)
	fmt.Println(sum)
}
