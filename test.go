package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	ch := make(chan bool, 2)
//
//	ch <- true
//	ch <- true
//
//	fmt.Println(len(ch))
//
//	func(ch chan bool) {
//		try1:
//		for {
//			select {
//
//			case <-ch:
//				fmt.Println("ok")
//				fmt.Println("ok")
//
//				ch = make(chan bool, 2)
//
//			default:
//				fmt.Println("channel is full !")
//				break try1
//			}
//		}
//	}(ch)
//
//	fmt.Println(len(ch))
//}
//
//
//
//
//
