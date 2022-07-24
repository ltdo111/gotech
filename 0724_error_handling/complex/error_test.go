// Alipay.com Inc. Copyright (c) 2004-2019 All Rights Reserved.
// @author bofeng.lt
// @version : error_test.go, v 0.1 2022/07/24 6:13 PM bofeng.lt Exp $$
// @Description:

package complex

import "testing"

func TestHandlingErrorWithCustomError(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "find by ", args: args{"age"}},
		{name: "find by ", args: args{"name"}},
		{name: "find by ", args: args{"sex"}},
		{name: "find by ", args: args{"addr"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HandlingErrorWithCustomError(tt.args.key)
		})
	}
}
