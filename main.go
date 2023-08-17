package main

import (
	"fmt"

	"albertzhong.com/californiadb/db"
)

func prog1() {
	z, err := db.CreateDatabase("./test_db")

	s := db.NewSchema([]string{"age", "isAdult"}, []db.DataType{db.INTEGER_TYPE, db.BOOLEAN_TYPE})
	table := db.NewTable("persons", *s)
	z.AddTable(table)

	if err != nil {
		panic(err)
	}
	if err = z.Save(); err != nil {
		panic(err)
	}
}

func prog2() {
	z, err := db.LoadDatabase("./test_db")
	if err != nil {
		panic(err)
	}
	fmt.Println(z.DatabasePath)
}

func main() {
	prog1()
	prog2()
}
