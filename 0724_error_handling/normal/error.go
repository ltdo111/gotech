// Alipay.com Inc. Copyright (c) 2004-2019 All Rights Reserved.
// @author bofeng.lt
// @version : error, v 0.1 2022/07/24 4:35 PM bofeng.lt Exp $$
// @Description: go 中 error 处理 - 三种项目场景下的最佳实践
// 对应博客 - https://medium.com/@sebdah/go-best-practices-error-handling-2d15e1f0c5ee
// 1. 简单错误 - 原生errors package + 直接比较

package normal

import (
	"errors"
	"fmt"
)

//type error interface {
//	Error() string
//}

var (
	errNotFound     = errors.New("item not found")
	errMissingParam = errors.New("missing param")
	errUnKnown      = "unKnownErr"
)

func NewMap() map[string]string {
	resMap := make(map[string]string)
	resMap["name"] = "bofeng"
	resMap["age"] = "28"
	return resMap
}

func getItem(key string) (any, error) {
	cn := NewMap()
	val, ok := cn[key]
	if !ok {
		return nil, errNotFound
	}
	return val, nil
}

func switchErrByComparison(error error) {
	if error == nil {
		return
	}
	switch error {
	case errNotFound:
		// 可以记录、上报错误日志
		fmt.Println(errNotFound.Error())
	case errMissingParam:
		fmt.Println(errMissingParam.Error())
	default:
		fmt.Println(errUnKnown)
	}
}

// HandlingErrorInSimpleCase 在简单场景下处理错误
func HandlingErrorInSimpleCase(key string) {
	val, err := getItem(key)
	if err != nil {
		switchErrByComparison(err)
		return
	}
	// 模拟最终的业务处理逻辑
	fmt.Println("未发生异常--", val)
	return
}

// 留一个问题，如果后续有根据 errType 分组的需求,即error中需要 记录errorType, msg 那 如何来实现这个需求呢? 具体最佳实践，我们下期见
