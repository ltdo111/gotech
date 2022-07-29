// Alipay.com Inc. Copyright (c) 2004-2019 All Rights Reserved.
// @author bofeng.lt
// @version : panic, v 0.1 2022/07/25 5:52 PM bofeng.lt Exp $$
// @Description:
// panic - 程序出现的异常，而非业务错误;
// 重要3点:
//       1. panic(errMsg) - 显式抛出异常 5/0
//       2. 关于groutine 的异常 recover()
//          编程范式:
//          defer func() {
//              // stops the panicking sequence, and return the panic(error)
//		        if err := recover(); err != nil {
//			        fmt.Printf("%s\n", err)
//		        }
//	        }()
//         将发生异常的groutine 恢复到正常情况（）， 并继续运行
//       3. recover 生效条件:
//            defer 中运行
//            只对当前 groutine 生效;

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

	fmt.Println("run ok, key == ", key)
}

func runRoutines() {
	for i := 0; i < 10; i++ {
		go PanicSimulate(i)
	}

	for {
		select {
		default:
			//fmt.Println("run")
		}
	}
}
