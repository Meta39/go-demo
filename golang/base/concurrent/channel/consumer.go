package channel

// Consumer 从通道中接收数据进行计算，接收参数是：chan int
func Consumer(ch chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

// Consumer2 参数为接收通道，接收参数是：<-chan int【多了个<-，表示接收通道】
func Consumer2(ch <-chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}
