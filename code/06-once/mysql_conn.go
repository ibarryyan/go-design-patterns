package _6_once

import "sync"

type MysqlConn struct {
	Addr string
}

var (
	mysqlConn *MysqlConn
	once = sync.Once{}
)


func GetMySQLConn() *MysqlConn {
	once.Do(func() {
		mysqlConn = &MysqlConn{Addr: "127.0.0.1"}
	})
	return mysqlConn
}

func NewMysqlConn() *MysqlConn {
	once.Do(func() {
		mysqlConn = &MysqlConn{Addr: "127.0.0.1"}
	})
	return mysqlConn
}
