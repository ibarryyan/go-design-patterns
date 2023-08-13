package _3_template

import (
	"fmt"
	"testing"
)

type DragonFruit struct {
}

func (d *DragonFruit) CreateYogurt() {
	fmt.Println("用鲜奶和乳酸菌发酵好酸奶")
}

func (d *DragonFruit) CutFruit() {
	fmt.Println("把火龙果切成块块")
}

func (d *DragonFruit) Merge() {
	fmt.Println("放在一起进行搅拌")
}

func (d *DragonFruit) Optimize() {
	fmt.Println("调制出自己喜欢的味道")
}

func (d *DragonFruit) Eating() {
	fmt.Println("开吃")
}

func (d *DragonFruit) Do() {
	d.CreateYogurt()
	d.CutFruit()
	d.Merge()
	d.Optimize()
	d.Eating()
}

func TestDragonFruit(t *testing.T) {
	d := &DragonFruit{}
	d.Do()
}

type Mango struct {
}

func (m *Mango) CreateYogurt() {
	fmt.Println("用鲜奶和乳酸菌发酵好酸奶")
}

func (m *Mango) CutFruit() {
	fmt.Println("把芒果切成块块")
}

func (m *Mango) Merge() {
	fmt.Println("放在一起进行搅拌")
}

func (m *Mango) Optimize() {
	fmt.Println("调制出自己喜欢的味道")
}

func (m *Mango) Eating() {
	fmt.Println("开吃")
}

func (m *Mango) Do() {
	m.CreateYogurt()
	m.CutFruit()
	m.Merge()
	m.Optimize()
	m.Eating()
}

func TestMango(t *testing.T) {
	m := &Mango{}
	m.Do()
}
