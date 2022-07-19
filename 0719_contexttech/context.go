// @Description: 学习 go 中 context 的使用
// 1. 了解context
// 2. 了解context 重要性
// 3. context基本用法
// 4. context使用场景 withCancel()

package _719_contexttech

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func NewCtx() {
	bg := context.Background()
	fmt.Println(bg)

	todo := context.TODO()
	fmt.Println(todo)
}

// CtxWithValue ctx中存储值.
// 突出2点:
//        1. ctx 不可变
//        2. ctx WithValue 后会基于Parent Context 生成 child Context
func CtxWithValue() {
	ctxA := context.Background()
	ctxB := context.WithValue(ctxA, "b", "B")
	ctxC := context.WithValue(ctxA, "c", "C")
	ctxD := context.WithValue(ctxB, "d", "D")
	ctxE := context.WithValue(ctxB, "e", "E")
	ctxF := context.WithValue(ctxC, "f", "F")

	fmt.Println(ctxA)
	fmt.Println(ctxB)
	fmt.Println(ctxC)
	fmt.Println(ctxD)
	fmt.Println(ctxE)
	fmt.Println(ctxF)
}

// AccessParentContextValue 访问父context的值.
// 程序中 context 之间的关系如下
//         ctxB
//       /
// ctxA
//       \
//         ctxC
//           \
//             ctxF
// 突出3点:
//       1. child context 可以访问 parent context
//       2. 分支路径上的 context 之间不能互相访问
//       3. parent context 不能访问 child context 中的值
func AccessParentContextValue() {
	ctxA := context.Background()
	ctxB := context.WithValue(ctxA, "b", "B")
	ctxC := context.WithValue(ctxA, "c", "C")
	ctxF := context.WithValue(ctxC, "f", "F")

	// 防止报错
	fmt.Println(ctxB)

	// 当前 context 访问 self context 内的值
	fmt.Println(ctxF.Value("f"))
	// 1.child context 可以访问 parent context
	fmt.Println(ctxF.Value("c"))
	// 2.分支路径上的 context 之间不能互相访问
	fmt.Println(ctxF.Value("b"))
	// 3. parent context 不能访问 child context 中的值
	fmt.Println(ctxA.Value("b"))
}

// WithCancelContext - 演示withCancel
func WithCancelContext() {
	bg := context.Background()
	ctx, cancel := context.WithCancel(bg)
	fmt.Println(ctx)

	wg := &sync.WaitGroup{}

	// monitor1 主动关闭 ctx
	wg.Add(1)
	go func(context.Context, *sync.WaitGroup) {
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				fmt.Println("monitor1 listen ctx done")
				return
			case <-time.After(4 * time.Second):
				fmt.Println("monitor1 being do cancel()")
				// cancel 之后 ctx.Done() 返回的是 close 后的 closeCh
				cancel()
			}
		}
	}(ctx, wg)

	// monitor2 监听ctx关闭动作
	wg.Add(1)
	go func(context.Context, *sync.WaitGroup) {
		cnt := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("monitor2 listen ctx done")
				wg.Done()
				return
			case <-time.After(2 * time.Second):
				cnt += 1
				fmt.Println("cnt-> ", cnt)
			}
		}
	}(ctx, wg)

	// 主 groutine 阻塞
	wg.Wait()

	// context.WithDeadline()
	// context.WithTimeout()
}
