package arc

import (
	"kar/items"
)

type ItemStack struct {
	ID       uint16
	Quantity uint8
}

type Inventory struct {
	SelectedSlot int
	Slots        [9]ItemStack
	HandSlot     ItemStack
}

func NewInventory() *Inventory {
	inv := &Inventory{}
	inv.HandSlot = ItemStack{}
	for i := range inv.Slots {
		inv.Slots[i] = ItemStack{}
	}
	return inv
}

// AddItemIfEmpty adds item to inventory if empty
func (i *Inventory) AddItemIfEmpty(id uint16) bool {
	idx, ok1 := i.HasItemStackSpace(id)
	if ok1 {
		i.Slots[idx].Quantity++
		return true
	} else {
		i2, ok2 := i.HasEmptySlot()
		if ok2 {
			i.Slots[i2].Quantity++
			i.Slots[i2].ID = id
			return true
		}
	}
	return false
}

func (i *Inventory) SetSlot(slotIndex int, id uint16, quantity uint8) {
	i.Slots[slotIndex] = ItemStack{
		ID:       id,
		Quantity: quantity,
	}
}

func (i *Inventory) SelectNextSlot() {
	if i.SelectedSlot+1 < len(i.Slots) {
		i.SelectedSlot++
	} else {
		i.SelectedSlot = 0
	}
}
func (i *Inventory) RemoveHandItem(id uint16) bool {
	ok := i.HasHandItem(id)
	if ok {
		i.HandSlot.Quantity--
		return true
	} else {
		i.HandSlot.ID = items.Air
	}
	return false
}

func (i *Inventory) RemoveItem(id uint16) bool {
	idx, ok := i.HasItem(id)
	if ok {
		i.Slots[idx].Quantity--
		return true
	}
	return false
}
func (i *Inventory) RemoveItemFromSelectedSlot() {
	if i.Slots[i.SelectedSlot].Quantity > 0 {
		i.Slots[i.SelectedSlot].Quantity--
	}
}
func (i *Inventory) SelectedSlotID() uint16 {
	return i.Slots[i.SelectedSlot].ID
}
func (i *Inventory) SelectedSlotQuantity() uint8 {
	return i.Slots[i.SelectedSlot].Quantity
}

func (i *Inventory) ClearSlot(index int) {
	i.Slots[index] = ItemStack{}
}

func (i *Inventory) ClearAllSlots() {
	for idx := range i.Slots {
		i.Slots[idx] = ItemStack{}
	}
}
func (i *Inventory) RandomFillAllSlots() {
	for idx := range i.Slots {
		i.SetSlot(idx, items.RandomBlock(), 10)
	}
}

func (i *Inventory) ClearSelectedSlot() {
	i.ClearSlot(i.SelectedSlot)
}

func (i *Inventory) HasEmptySlot() (index int, ok bool) {
	for idx, v := range i.Slots {
		if v.Quantity == 0 {
			return idx, true
		}
	}
	return -1, false
}

func (i *Inventory) HasItemStackSpace(id uint16) (index int, ok bool) {
	for idx, v := range i.Slots {
		if v.ID == id && v.Quantity < 64 && v.Quantity > 0 {
			return idx, true
		}
	}
	return -1, false
}

func (i *Inventory) HasItem(id uint16) (index int, ok bool) {
	for idx, v := range i.Slots {
		if v.ID == id && v.Quantity > 0 {
			return idx, true
		}
	}
	return -1, false
}
func (i *Inventory) HasHandItem(id uint16) bool {
	if i.HandSlot.ID == id && i.HandSlot.Quantity > 0 {
		return true
	}
	return false
}
