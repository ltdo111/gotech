// @Description: the tech of synchronization components in golang - sync.WaitGroup、mutex etc;
// 探讨点:
//      1. 为什么需要并发-数据竞争?
//      2. 如何解决并发?
//         - sync/atomic
//         - sync.Mutex - 类 Java Lock 接口
//         - sync.RMutex - 类 Java ReentrantReadWriteLock
//      3. 多 routine 同步 - sync.WaitGroup
//      4. 多 routine 情况下 -  sync.Once 只执行一次
//      5. sync.Cond - 多 groutine 之间的通信, 类Java Condition, wait/notify 经典范式=加锁+循环+等待
//      6. sync.Map - 线程安全的 Map 容器;
//      录制计划 - 分 6 小节

package _721_syncronized

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// add 单 routine 执行+操作.
func add(w *sync.WaitGroup, num *int) {
	defer w.Done()
	*num = *num + 1
}

// UnsafeAdd 多 go routines 不安全相加.
// 演示  data competition 即 数据竞争.
func UnsafeAdd() {
	var n int = 0
	var wg *sync.WaitGroup = new(sync.WaitGroup)
	wg.Add(1000)
	for i := 0; i < 1000; i = i + 1 {
		go add(wg, &n)
	} // spawn 1000 new goroutines
	wg.Wait()
	println(n)
}

//SafeAddWithAtomic 使用 atomic.AddInt32() 进行原子性加, 类Java CAS.
func SafeAddWithAtomic() {
	var n int32 = 0
	var wg *sync.WaitGroup = new(sync.WaitGroup)
	wg.Add(1000)
	for i := 0; i < 1000; i = i + 1 {
		go func(*sync.WaitGroup, *int32) {
			defer wg.Done()
			atomic.AddInt32(&n, 1)
		}(wg, &n)
	} // spawn 1000 new goroutines
	wg.Wait()
	println(n)
}

// SafeAddAndMinusWithSyncMutex 多 go routines 同时 加、减操作通过使用 Sync.Mutex 实现 多routines 安全.
// the second tech
func SafeAddAndMinusWithSyncMutex() {
	var num int = 0

	var mutex = new(sync.Mutex)
	var wg = new(sync.WaitGroup)
	wg.Add(2)

	// Add
	go func(*sync.Mutex, *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 100000; i = i + 1 {
			mutex.Lock()
			num = num + 1
			mutex.Unlock()
		}
	}(mutex, wg)

	// Minus
	go func(*sync.Mutex, *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 100000; i = i + 1 {
			// 跟Java AbstractQueuedSynchronizer - 抽象队列同步器原理相同;
			mutex.Lock()
			num = num - 1
			mutex.Unlock()
		}
	}(mutex, wg)

	wg.Wait()
	print(num)
}

// ExecutedCodeOnceBySyncOnce - 演示只执行一次动作, 比如项目有个清理动作，防止重复操作.
// the third tech
func ExecutedCodeOnceBySyncOnce() {
	once := new(sync.Once)
	var wg = new(sync.WaitGroup)
	wg.Add(10)

	for i := 0; i < 10; i = i + 1 {
		tmp := i
		go func() {
			defer wg.Done()
			fmt.Println(tmp)
			once.Do(func() {
				fmt.Println("RunOnce")
			})
		}()
	}

	wg.Wait()
}

// MultiRoutinesCommunicateBySyncCond 多协程通过 syncCond 来进行通信交流;
func MultiRoutinesCommunicateBySyncCond() {
	stopCh := make(chan struct{})

	lc := new(sync.Mutex)
	cond := sync.NewCond(lc)

	num := 0

	// consumer
	go func() {
		for {
			cond.L.Lock()
			for num == 0 {
				cond.Wait()
			}
			num -= 1
			fmt.Println("consumer-> ", num)
			cond.Signal()
			cond.L.Unlock()
		}
	}()

	// producer
	go func() {
		for {
			cond.L.Lock()
			for num == 3 {
				cond.Wait()
			}
			num += 1
			fmt.Println("producer -> ", num)
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
