// Alipay.com Inc. Copyright (c) 2004-2019 All Rights Reserved.
// @author bofeng.lt
// @version : error, v 0.1 2022/07/24 4:35 PM bofeng.lt Exp $$
// @Description: go 中 error 处理 - 三种项目场景下的最佳实践
// 对应博客 - https://medium.com/@sebdah/go-best-practices-error-handling-2d15e1f0c5ee
// 1. 简单错误 - 原生errors package + 直接比较
// 2. 复杂错误 - 自定义error structs + 使用类型检查
// 3. 错误处理的时机

package _724_error_handling

import "errors"

//type error interface {
//	Error() string
//}

var (
	errNotFound     = errors.New("item not found")
	errMissingParam = errors.New("missing param")
)

// HandlingErrorInSimpleCase 在简单场景下处理错误
func HandlingErrorInSimpleCase() {

}
