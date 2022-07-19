package _718_function

import (
	"testing"
)

func TestOp(t *testing.T) {
	msg := "[caller_once]"

	// 解码
	decodeOp := decode(msg)
	decodedMsg, err := decodeOp(msg)
	if err != nil {
		return
	}

	// 执行业务逻辑
	actionOp := opAction(decodedMsg)
	opRes, err := actionOp(decodedMsg)
	if err != nil {
		return
	}

	// 编码
	encodeOp := encode(opRes)
	_, err = encodeOp(opRes)
	if err != nil {
		return
	}

}
