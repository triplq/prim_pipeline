package main

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

func main() {

}
