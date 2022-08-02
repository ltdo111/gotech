// Alipay.com Inc. Copyright (c) 2004-2019 All Rights Reserved.
// @author bofeng.lt
// @version : channel, v 0.1 2022/08/02 7:23 PM bofeng.lt Exp $$
// @Description:

package _805_channel_basic

import "fmt"

func blockCh() {
	msg := make(chan string)
	msg <- "Hey There"
	go func() {
		fmt.Println(<-msg)
	}()
}

func noBlockCh() {
	msg := make(chan string, 1)
	msg <- "Hey There"
	go func() {
		fmt.Println(<-msg)
	}()
}
