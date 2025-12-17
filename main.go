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

func main() {

}
