package californiadb

import (
	"errors"
	"fmt"
	"os"
)

const (
	PageSize = 4048
	Header   = 0xBEEFFACE
)

type Page struct {
	fileName string
	offset   int
	schema   *Schema
	// in-memory buffer of page data
	buffer []byte
	// logical number of tuples
	size int
}

func NewPage(fileName string, offset int, schema *Schema) (*Page, error) {
	if offset < 0 {
		return nil, errors.New("offset must be non-negative")
	}
	page := Page{
		fileName: fileName,
		offset:   offset,
		schema:   schema,
		buffer:   make([]byte, PageSize),
	}
	return &page, nil
}

// Read the page from disk.
func (p Page) Read() error {
	pageFile, err := os.Open(p.fileName)
	defer pageFile.Close()
	if err != nil {
		return fmt.Errorf("could not open page: %w", err)
	}
	_, err = pageFile.ReadAt(p.buffer, int64(p.offset*PageSize))
	if err != nil {
		return fmt.Errorf("could not read page: %w", err)
	}
	return nil
}

// Flush the page to disk.
func (p Page) Flush() error {
	pageFile, err := os.Open(p.fileName)
	defer pageFile.Close()
	if err != nil {
		return fmt.Errorf("could not open page: %w", err)
	}
	_, err = pageFile.WriteAt(p.buffer, int64(p.offset*PageSize))
	if err != nil {
		return fmt.Errorf("could not write page: %w", err)
	}
	return nil
}

// Add a tuple to the page.
func (p Page) Add() error {

}
