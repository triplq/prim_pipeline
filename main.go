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

func sumFunc(ch <-chan int) int {
	sum := 0
	for c := range ch {
		sum += c
	}

	return sum
}

func main() {

}
