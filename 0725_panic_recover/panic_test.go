// Alipay.com Inc. Copyright (c) 2004-2019 All Rights Reserved.
// @author bofeng.lt
// @version : panic_test.go, v 0.1 2022/07/25 5:57 PM bofeng.lt Exp $$
// @Description:

package _725_panic_recover

import "testing"

func Test_runRoutines(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "test_1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runRoutines()
		})
	}
}
