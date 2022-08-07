// Alipay.com Inc. Copyright (c) 2004-2019 All Rights Reserved.
// @author bofeng.lt
// @version : channel, v 0.1 2022/08/02 7:23 PM bofeng.lt Exp $$
// @Description:

package main

import (
	"fmt"
	//_ "net/http/pprof"
)

func blockCh() {
	msg := make(chan string)
	// for select goroutine, 导致死锁检测器实效;
	go func() {
		for {
			select {
			default:

			}
		}
	}()
	msg <- "Hey There"

	//go func() {
	//fmt.Println(<-msg)
	//}()
}

func noBlockCh() {
	msg := make(chan string, 1)
	msg <- "Hey There"
	go func() {
		fmt.Println(<-msg)
	}()
}

//func main() {
//	stopped := make(chan bool)
//	// step 4:  managerFactory (车间)持续运行.
//	//for {
//	//	select {
//	//	// chan struct{} 类型, 是用来控制 goroutine 是否持续运行的信号量.
//	//	case <-mf.Stopped():
//	//		return
//	//	default:
//	//	}
//	//}
//	//go func() {
//	//	fmt.Println("")
//	//}()
//	//
//	//time.Sleep(time.Second)
//
//	<-stopped
//}
