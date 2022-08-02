// Alipay.com Inc. Copyright (c) 2004-2019 All Rights Reserved.
// @author bofeng.lt
// @version : default, v 0.1 2022/08/02 5:30 PM bofeng.lt Exp $$
// @Description:

package _804_select_default

import (
	"fmt"
	"time"
)

func sleep(seconds int, endSignal chan<- bool) {
	time.Sleep(time.Duration(seconds) * time.Second)
	endSignal <- true
}

func run() {
	endSignal := make(chan bool, 1)
	go sleep(3, endSignal)

	timeout := time.After(5 * time.Second)
	for {
		select {
		case <-endSignal:
			fmt.Println("The end!")
		case <-timeout:
			fmt.Println("There's no more time to this. Exiting!")
		default:
		}
	}
}
