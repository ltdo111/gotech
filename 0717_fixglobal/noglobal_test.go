// Alipay.com Inc. Copyright (c) 2004-2019 All Rights Reserved.
// @author bofeng.lt
// @version : noglobal_test, v 0.1 2022/07/17 2:09 PM bofeng.lt Exp $$
// @Description:

package _717_fixglobal

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestDemoModel_All(t *testing.T) {
	dm := DemoModel{
		DB: initDB(),
	}

	dms, err := dm.All()
	if err != nil {
		return
	}

	fmt.Println(dms)
}

// initDB.
func initDB() *sql.DB {
	return nil
}
