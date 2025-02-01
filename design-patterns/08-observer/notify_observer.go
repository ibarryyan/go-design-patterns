package _8_observer

import (
	"fmt"
)

// Subject 是被观察者
type Subject struct {
	observers []Observer
	state     string
}

// Register 注册观察者
func (s *Subject) Register(observer Observer) {
	s.observers = append(s.observers, observer)
}

// SetState 设置状态并通知观察者
func (s *Subject) SetState(state string) {
	s.state = state
	s.Notify()
}

// Notify 通知所有观察者
func (s *Subject) Notify() {
	for _, observer := range s.observers {
		go observer.Update(s.state) // 使用协程避免阻塞
	}
}

// Observer 观察者接口
type Observer interface {
	Update(string)
}

// ConcreteObserver1 具体观察者1
type ConcreteObserver1 struct{}

func (o ConcreteObserver1) Update(state string) {
	fmt.Printf("ConcreteObserver1 received state: %s\n", state)
}

// ConcreteObserver2 具体观察者2
type ConcreteObserver2 struct{}

func (o ConcreteObserver2) Update(state string) {
	fmt.Printf("ConcreteObserver2 received state: %s\n", state)
}
