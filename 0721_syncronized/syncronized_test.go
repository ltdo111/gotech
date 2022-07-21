// @Description: tech golang sync、mutex 等同步源语功能;

package _721_syncronized

import "testing"

func Test_UnsafeAdd(t *testing.T) {
	UnsafeAdd()
}

func Test_SafeAddWithAtomic(t *testing.T) {
	SafeAddWithAtomic()
}

func Test_SafeAddAndMinusWithSyncMutex(t *testing.T) {
	SafeAddAndMinusWithSyncMutex()
}

func Test_ExecutedCodeOnceBySyncOnce(t *testing.T) {
	ExecutedCodeOnceBySyncOnce()
}

func TestMultiRoutinesCommunicateBySyncCond(t *testing.T) {
	MultiRoutinesCommunicateBySyncCond()
}
