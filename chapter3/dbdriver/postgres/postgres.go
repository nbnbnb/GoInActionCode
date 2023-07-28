package postgres

import (
	"database/sql"
	"database/sql/driver"
	"errors"
)

// 自定义一个 PostgresDriver
type PostgresDriver struct{}

// PostgresDriver 实现 driver.Driver 的 Open 方法
func (dr PostgresDriver) Open(string) (driver.Conn, error) {
	return nil, errors.New("Unimplemented")
}

var d *PostgresDriver

// 初始化时调用
func init() {
	d = new(PostgresDriver)
	sql.Register("postgres", d)
}
