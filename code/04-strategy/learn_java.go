package _4_strategy

import "fmt"

type LikeJava struct {
}

func (g *LikeJava) PrepareData() {
	fmt.Println("准备Java资料")
}

func (g *LikeJava) DoLearn() {
	fmt.Println("学习Java")
}

func (g *LikeJava) Play() {
	fmt.Println("玩Java")
}
