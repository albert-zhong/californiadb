package db

import "fmt"

const PAGE_SIZE = 4096

type PageId struct {
	path   string
	offset int64
}

type Page struct {
	PageId
	data []byte

	dirty    bool
	capacity int
}

type BoolPage Page

const BOOL_PAGE_CAPACITY = 3640
const BOOL_PAGE_BITMAP_SIZE = 455

// Slots returns the capacity of values on this page
func (p *BoolPage) Slots() int {
	return BOOL_PAGE_CAPACITY
}

func (p *BoolPage) slotToIndex(slot int) int {
	return BOOL_PAGE_BITMAP_SIZE + slot
}

func (p *BoolPage) SlotIsEmpty(slot int) bool {
	c := p.data[1+(slot/8)]
	i := (slot % 8)
	isEmpty := (c >> i) & 1
	return isEmpty == 0
}

func (p *BoolPage) Insert(value bool) error {
	for i := 0; i < p.Slots(); i++ {
		if p.SlotIsEmpty(i) {
			if err := p.Update(i, value); err != nil {
				return nil
			}
		}
	}
	return fmt.Errorf("cannot insert value into page")
}

func (p *BoolPage) Update(slot int, value bool) error {
	p.data[p.slotToIndex(slot)] = byte(value)
}

func (p *BooleanPage) Delete() {

}
