package _6_once

type DB struct {
	Addr, UserName, Passwd string
}

func NewMysqlConn(addr, userName, passwd string) *DB {
	return &DB{
		Addr:     addr,
		UserName: userName,
		Passwd:   passwd,
	}
}
