package _6_once

import (
	"fmt"
	"testing"
)

func TestNewMysqlConn(t *testing.T) {
	conn := NewMysqlConn()
	fmt.Println(conn)
}
