package main

func main() {
	arr := []int{1, 2, 3}
	ch := make(chan int)
	go sum(arr, ch)
	<-ch
}

func sum(arr []int, ch chan int) {
	var s int = 0
	for _, v := range arr {
		s += v
	}
	ch <- s
}
