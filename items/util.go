package items

import (
	"kar/engine/mathutil"
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
	index := mathutil.RandRangeInt(0, len(BlockIDs)-1)
	return BlockIDs[index]
}
