package db

import (
	"os"
)

const MAX_PAGES = 512

type BufferPool struct {
	pages map[PageId]Page
}

func (p Page) flush() error {
	// no need to write if the page was not written to
	if !p.dirty {
		return nil
	}

	f, err := os.OpenFile(p.path, os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteAt(p.data, p.offset*PAGE_SIZE)
	if err != nil {
		return err
	}
	return nil
}

func NewBufferPool() *BufferPool {
	bufferPool := BufferPool{}
	return &bufferPool
}

func (bp BufferPool) GetPage(path string, offset int64) (Page, error) {
	// Return the page if it already exists in the buffer pool
	id := PageId{path, offset}
	if page, ok := bp.pages[id]; ok {
		return page, nil
	}

	if len(bp.pages) >= MAX_PAGES {
		// TODO: implement cache eviction policy
	}

	f, err := os.OpenFile(path, os.O_RDWR, 0755)
	if err != nil {
		return Page{}, err
	}
	defer f.Close()

	data := make([]byte, PAGE_SIZE)
	_, err = f.ReadAt(data, offset*PAGE_SIZE)
	if err != nil {
		return Page{}, err
	}

	page := Page{
		PageId: PageId{
			path:   path,
			offset: offset,
		},
		dirty: false,
		data:  data,
	}

	return page, nil
}
