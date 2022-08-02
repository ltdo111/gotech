// Alipay.com Inc. Copyright (c) 2004-2019 All Rights Reserved.
// @author bofeng.lt
// @version : 0803_Broadcaster, v 0.1 2022/08/01 7:40 PM bofeng.lt Exp $$
// @Description:

package _803_broadcaster

// Master 主节点 leadership.
// env
// hook slaver-master -> 新 slaver master
// switch true.
type Master struct {
	host      string
	otherMeta any
	slaver    []Slaver
}

// Slaver 定义Slaver.
type Slaver struct {
	host      string
	otherMeta any
}

// 写同步逻辑。 就是一个rpc请求， slaver 开接口，接收来自master 的信息,并存储;
func (s Slaver) sync(info any) {
	// rpc_caller
}

// AllSlavers 返回所有slaver.
func AllSlavers() []*Slaver {
	slavers := make([]*Slaver, 10)
	return slavers
}

type Dispatcher interface {
	Dispatch(info any) error
}

type StaticDispatcher struct {
}

func (sd StaticDispatcher) Dispatch(info any) error {
	slavers := AllSlavers()
	for _, slaver := range slavers {
		slaver.sync(info)
	}
	return nil
}
