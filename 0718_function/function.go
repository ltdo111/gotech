// @Description: 模拟服务端Rpc处理.

package _718_function

import "fmt"

// Op 定义的对msg处理的函数集合.
type Op func(msg any) (any, error)

// decode 解码远程消息.
func decode(msg any) Op {
	return func(any) (any, error) {
		// decode msg
		fmt.Println("decoding ... ", msg)
		decodedRes := fmt.Sprintf("decoded_%v", msg)
		fmt.Println("decoded to parameter-> ", decodedRes)
		return decodedRes, nil
	}
}

// opAction 模拟服务提供方处理业务逻辑.
func opAction(parameter any) Op {
	return func(any) (any, error) {
		// decode msg
		fmt.Println("do opAction ... ", parameter)
		opRes := fmt.Sprintf("opAction_%v", parameter)
		fmt.Println("after opAction result -> ", opRes)

		return opRes, nil
	}
}

// encode 模拟服务提供方将处理结果
func encode(result any) Op {
	return func(any) (any, error) {
		// decode msg
		fmt.Println("encoding ... ", result)
		encodedRes := fmt.Sprintf("encoded_%v", result)
		fmt.Println("after encoded result -> ", encodedRes)
		return encodedRes, nil
	}
}
