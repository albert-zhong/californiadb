package californiadb

import "fmt"

type DataType interface {
	Name() string
	Size() int
}

type StringDataType interface {
	ReadValue(data []byte) string
}

type IntegerDataType interface {
	ReadValue(data []byte) int
}

type VarCharType struct {
	size int
}

func NewVarCharType(size int) VarCharType {
	return VarCharType{size: size}
}

func (v VarCharType) Name() string {
	return fmt.Sprintf("VARCHAR(%d)", v.size)
}

func (v VarCharType) Size() int {
	return 4 + v.size
}

func (v VarCharType) ReadValue(data []byte) int {

}

type Schema struct {
	columnNames []string
	columnTypes []DataType
}

func NewSchema(columnNames []string, columnTypes []DataType) *Schema {
	schema := Schema{columnNames, columnTypes}
	return &schema
}

type Table struct {
	name   string
	schema Schema
}

type Database struct {
	name   string
	tables map[string]*Table
}
