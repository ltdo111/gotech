// Alipay.com Inc. Copyright (c) 2004-2019 All Rights Reserved.
// @author bofeng.lt
// @version : error, v 0.1 2022/07/24 5:59 PM bofeng.lt Exp $$
// @Description: 复杂错误- 自定义错误
// 1. 复杂错误 - 自定义error structs + 使用类型检查
// 3. 错误处理的时机

package complex

import (
	"fmt"
	"gotech/0724_error_handling/normal"
)

type errType string

const (
	errNotFound     errType = "item not found"
	errMissingParam errType = "missing param"
	errUnKnown      errType = "unKnownErr"
)

type BusinessError struct {
	errType errType
	msg     string
}

func NewBusinessError(errType errType, msg string) *BusinessError {
	return &BusinessError{errType, msg}
}

func (e *BusinessError) Error() string {
	return fmt.Sprintf("%v_%v", e.errType, e.msg)
}

func getItem(key string) (any, error) {
	cn := normal.NewMap()
	val, ok := cn[key]
	if !ok {
		return nil, NewBusinessError(errNotFound, key)
	}
	return val, nil
}

// switchErrByTypeChecking 通过错误类型判断，来区分错误;
func switchErrByTypeChecking(err error) {
	if err != nil {
		switch err.(type) {
		case *BusinessError:
			fmt.Println(err.Error())
		default:
			fmt.Println(errUnKnown)
		}
	}
}

// HandlingErrorWithCustomError 复杂场景下自定义Error
func HandlingErrorWithCustomError(key string) {
	val, err := getItem(key)
	if err != nil {
		switchErrByTypeChecking(err)
		return
	}
	// 模拟最终的业务处理逻辑
	fmt.Println("未发生异常--", val)
	return
}
