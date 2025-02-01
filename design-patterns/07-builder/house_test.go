package _7_builder

import "testing"

func TestMyHouseBuilder(t *testing.T) {

	h := &HouseBuilder{}

	house := h.AddStep(&HouseBuilder{}).Build()

	house.Build()

}
