// @Description:

package _717_fixglobal

import "database/sql"

// Demo 演示示例.
type Demo struct {
	Name        string
	Description string
}

// DemoModel 自定义 DemoModel 内聚了DB pool.
type DemoModel struct {
	// type embedding  垂直组合 sql.DB 的能力
	*sql.DB
}

// All 使用 DemoModel 跑这个SQL query.
func (m DemoModel) All() ([]Demo, error) {
	rows, err := m.Query("SELECT name, description FROM demo")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var demos []Demo
	for rows.Next() {
		var dm Demo

		err := rows.Scan(&dm.Name, &dm.Description)
		if err != nil {
			return nil, err
		}

		demos = append(demos, dm)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return demos, nil
}
