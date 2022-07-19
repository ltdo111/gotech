// @Description:

package _716_channel

import (
	"fmt"
	"testing"
	"time"
)

func Test_Stop(t *testing.T) {
	stopCh := make(chan struct{})

	fn := func() {
		fmt.Println("fn run")
	}

	// 模拟3s之后停止
	go func(chan struct{}) {
		select {
		case <-time.After(3 * time.Second):
			close(stopCh)
			fmt.Println("stopCh closed")
		}
	}(stopCh)

	BackOffUntil(stopCh, fn)
}
