// Alipay.com Inc. Copyright (c) 2004-2019 All Rights Reserved.
// @author bofeng.lt
// @version : syncronized_cond, v 0.1 2022/07/23 8:55 PM bofeng.lt Exp $$
// @Description: 演示使用 sync.Cond 解决 生产者消费者 问题
// 本质上是协作， 协作本质就是同步信息; 所以 生产者消费者的问题本质上就是多 routine 之间的 同步问题;

// 举个🌰: 夏天喝咖啡加冰块, 需要制冰机, 加入将制冰机想像成 生产者 routine, 取冰的人是consumer routine,
//        那么 制冰机 在机箱满的时候需要停止生产， 机箱空的时候，取冰的人需要停止取冰，
//        这样一种场景我们就叫生产者消费者;

package _721_syncronized

import (
	"fmt"
	"sync"
)

// 最多冰块数, 代表杯子已经满了, 杯子不能再装新的冰块了;
const maxCnt = 3

// 最少冰块数, 代表杯子中没有冰块了.
const minCnt = 0

// iceCubes 冰块.
type iceCube int

// cup 冰箱.
type cup struct {
	iceCubes []iceCube
}

// ProduceAndConsumerSimulationWithSyncCond 多协程通过 syncCond 来进行生产者消费者的模拟;
// 类比Java 加锁 + 循环&&等待 唤醒 在 Go 中也是 加锁 + 循环&&等待 经典范式
// A: 为什么加锁:
// Q:
//    1. 加锁 获得程序执行权;
//    2. 不加锁情况下， 如果生产冰块同时时还能从杯子中拿出冰块, 万一生产速率 < 拿取速率， 杯子就空了, 反之，杯子溢出冰块，都是
//       非预期的情况
// A: 思考为什么需要cond.wait(). cond.signal()?
// Q:
//    1. 因为要通知阻塞的 routine 重新获得执行权(获取锁);
//
func ProduceAndConsumerSimulationWithSyncCond() {
	stopCh := make(chan struct{})

	lc := new(sync.Mutex)
	cond := sync.NewCond(lc)

	cup := cup{
		iceCubes: make([]iceCube, 3, 3),
	}

	// consumer
	go func() {
		for {
			cond.L.Lock()
			for len(cup.iceCubes) == minCnt {
				cond.Wait()
			}
			// 删除头部的冰块
			cup.iceCubes = cup.iceCubes[1:]
			fmt.Println("consume 1 iceCube, left iceCubes ->  ", len(cup.iceCubes))
			cond.Signal()
			cond.L.Unlock()
		}
	}()

	// producer
	go func() {
		for {
			cond.L.Lock()
			for len(cup.iceCubes) == maxCnt {
				cond.Wait()
			}
			// 杯子中新添加进一个冰块.
			cup.iceCubes = append(cup.iceCubes, 1)
			fmt.Println("producer 1 iceCube, left iceCubes ", len(cup.iceCubes))
			cond.Signal()
			cond.L.Unlock()
		}

	}()

	for {
		select {
		case <-stopCh:
			return
		default:
		}
	}
}
