// @Description:

package _716_channel

import (
	"fmt"
	"time"
)

// BackOffUntil groutine 起停函数.
func BackOffUntil(stopCh chan struct{}, fn func()) {
	for {
		select {
		// stopCh， close掉了，然后读到值
		case <-stopCh:
			return
		case <-time.After(1 * time.Second):
			fmt.Println("begin running fn()")
			fn()
		}
	}
}
