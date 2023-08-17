package db

import (
	"encoding/json"
	"errors"
	"os"
	"path"
)

type Schema struct {
	ColumnNames     []string `json:"columnNames"`
	ColumnTypeNames []string `json:"columnTypeNames"`

	columnTypes []DataType
}

func NewSchema(columnNames []string, columnTypes []DataType) *Schema {
	columnTypeNames := make([]string, 0, len(columnTypes))
	for _, columnType := range columnTypes {
		columnTypeNames = append(columnTypeNames, columnType.String())
	}
	schema := Schema{
		ColumnNames:     columnNames,
		ColumnTypeNames: columnTypeNames,
		columnTypes:     columnTypes,
	}
	return &schema
}

type Table struct {
	Name   string `json:"name"`
	Schema Schema `json:"schema"`
}

func NewTable(name string, schema Schema) *Table {
	table := Table{Name: name, Schema: schema}
	return &table
}

type Database struct {
	DatabasePath string            `json:"databasePath"`
	Tables       map[string]*Table `json:"tables"`
}

func CreateDatabase(databasePath string) (*Database, error) {
	// TODO: return error if a folder already exists there
	db := Database{
		DatabasePath: databasePath,
		Tables:       make(map[string]*Table),
	}
	return &db, nil
}

func LoadDatabase(databasePath string) (*Database, error) {
	metadataPath := path.Join(databasePath, "metadata.json")
	metadata, err := os.ReadFile(metadataPath)
	if err != nil {
		return nil, err
	}
	var db Database
	if err = json.Unmarshal(metadata, &db); err != nil {
		return nil, err
	}
	for _, table := range db.Tables {
		table.Schema.columnTypes = make([]DataType, 0, len(table.Schema.ColumnTypeNames))
		for _, columnTypeName := range table.Schema.ColumnTypeNames {
			columnType, err := ToDataType(columnTypeName)
			if err != nil {
				return nil, err
			}
			table.Schema.columnTypes = append(table.Schema.columnTypes, columnType)
		}
	}
	return &db, nil
}

func (db *Database) AddTable(table *Table) error {
	if _, ok := db.Tables[table.Name]; ok {
		return errors.New("Database alreay contains a table with the same name")
	}
	db.Tables[table.Name] = table
	return nil
}

func (db *Database) Save() error {
	// Create dir if it does not already exist
	if _, err := os.Stat(db.DatabasePath); os.IsNotExist(err) {
		if err = os.Mkdir(db.DatabasePath, 0755); err != nil {
			return err
		}
	}
	// Create dir/metadata.json if it does not already exist
	metadataPath := path.Join(db.DatabasePath, "metadata.json")
	f, err := os.OpenFile(metadataPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	metadata, err := json.Marshal(db)
	if err != nil {
		return err
	}
	if _, err = f.Write(metadata); err != nil {
		return err
	}
	return nil
}
