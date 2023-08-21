package _4_strategy

import "fmt"

type LikeGo struct {
}

func (g *LikeGo) PrepareData() {
	fmt.Println("准备Go资料")
}

func (g *LikeGo) DoLearn() {
	fmt.Println("学习Go")
}

func (g *LikeGo) Play() {
	fmt.Println("玩转Go语言")
}
