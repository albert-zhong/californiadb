package db

type ColumnFile struct {
	path       string
	columnType DataType
}

func NewColumnFile(path string, columnType DataType) *ColumnFile {
	columnFile := ColumnFile{path: path, columnType: columnType}
	return &columnFile
}

func (cf *ColumnFile) InsertBooleans(values []bool) error {
	return nil
}

func (cf *ColumnFile) InsertIntegers(values []int32) error {
	return nil
}

func (cf *ColumnFile) InsertVarChars(values []string, n int) error {
	return nil
}
