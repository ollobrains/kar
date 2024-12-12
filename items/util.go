package items

import (
	"math/rand/v2"
)

var BlockIDs []uint16

func init() {
	for id := range Property {
		if IsBlock(id) {
			BlockIDs = append(BlockIDs, id)
		}
	}
}

func IsBreakable(id uint16) bool {
	return Property[id].Category&Unbreakable == 0
}

func IsHarvestable(id uint16) bool {
	return Property[id].Category&Harvestable != 0
}
func IsBlock(id uint16) bool {
	return Property[id].Category&Block != 0
}

func RandomBlock() uint16 {
	return BlockIDs[rand.IntN(len(BlockIDs))]
}
