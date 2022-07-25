// Alipay.com Inc. Copyright (c) 2004-2019 All Rights Reserved.
// @author bofeng.lt
// @version : panic, v 0.1 2022/07/25 5:52 PM bofeng.lt Exp $$
// @Description:

package _725_panic_recover

import "fmt"

// PanicSimulate 模拟 panic
func PanicSimulate(key int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%s\n", err)
		}
	}()

	if key == 2 {
		panic(fmt.Sprintf("panic from key == %v", key))
	}

	go func() {
		//defer func() {
		//	if err := recover(); err != nil {
		//		fmt.Printf("%s\n", err)
		//	}
		//}()
		panic(fmt.Sprintf("panic  again"))
	}()

	fmt.Println("run ok, key == ", key)
}

func runRoutines() {
	for i := 0; i < 3; i++ {
		go PanicSimulate(i)
	}

	for {
		select {
		default:
			fmt.Println("run")
		}
	}
}
