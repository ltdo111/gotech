// Alipay.com Inc. Copyright (c) 2004-2019 All Rights Reserved.
// @author bofeng.lt
// @version : channel_test.go, v 0.1 2022/08/02 7:24 PM bofeng.lt Exp $$
// @Description:

package main

import "testing"

func Test_blockCh(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "test_blockCh"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			blockCh()
		})
	}
}

func Test_noBlockCh(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "test_noBlockCh"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			noBlockCh()
		})
	}
}
