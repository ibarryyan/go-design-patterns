package _8_observer

import (
	"testing"
	"time"
)

func TestName(t *testing.T) {
	subject := Subject{}
	observer1 := ConcreteObserver1{}
	observer2 := ConcreteObserver2{}

	subject.Register(observer1) // 注册观察者1
	subject.Register(observer2) // 注册观察者2

	subject.SetState("new state") // 设置状态并通知观察者
	time.Sleep(5*time.Second)
}