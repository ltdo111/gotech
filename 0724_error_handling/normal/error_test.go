// Alipay.com Inc. Copyright (c) 2004-2019 All Rights Reserved.
// @author bofeng.lt
// @version : error_test.go, v 0.1 2022/07/24 5:36 PM bofeng.lt Exp $$
// @Description:

package normal

import "testing"

func TestHandlingErrorInSimpleCase(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HandlingErrorInSimpleCase(tt.args.key)
		})
	}
}
