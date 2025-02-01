package _7_builder

import "fmt"

type House interface {
	Build()
}

type MyHouse struct {
	name string
}

func (p *MyHouse) Build() {
	fmt.Println("Building myHouse :", p.name)
}

type Builder interface {
	AddStep(step Builder) Builder
	Build() House
}

type HouseBuilder struct {
	product *MyHouse
	step    Builder
}

func (b *HouseBuilder) AddStep(step Builder) Builder {
	if step == nil {
		return nil
	}
	return &HouseBuilder{product: b.product, step: step}
}

func (b *HouseBuilder) Build() House {
	b.product = &MyHouse{name: "Barry House"}
	return b.product
}
