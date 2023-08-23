package _6_once

import "sync"

type MysqlConn struct {
	Addr string
}

var mysqlConn *MysqlConn

var once = sync.Once{}

func NewMysqlConn() *MysqlConn {
	once.Do(func() {
		mysqlConn = &MysqlConn{Addr: "127.0.0.1"}
	})
	return mysqlConn
}
