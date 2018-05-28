package main

import "fmt"

//解决chan 阻塞方法 1 ：使用buf
//buf == 0 || 未设置 同步，阻塞
//buf > 0 || 可能不阻塞 ，取决于元素个数
func main() {
	buf := 2
	ch := make(chan string, buf)
	sendMsg(ch)
	printMsg(ch)
}

func sendMsg(ch chan string) {
	ch <- "s"
	ch <- "sd"
}

func printMsg(ch chan string) {
	fmt.Println(cap(ch))
	fmt.Println(<-ch)

}
