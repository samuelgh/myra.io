package main

import (
"database/sql"
_ "github.com/mattn/go-sqlite3"
"log"
"fmt"
)

var db *sql.DB

type Item struct {
	Id      string
	Name    string
	Field1  string
	Field2  string
	Field3  string
	Field4  string
	Field5  string
	Field6  string
	Field7  string
	Field8  string
	Field9  string
	Field10 string
	Timestamp string
}

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./db.db")
	if ( err != nil) {
		log.Fatal(err)
	}
	defer db.Close()
	createTable()
}

func createTable() {
	sql_table := `
	CREATE TABLE IF NOT EXISTS items(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		Name	TEXT NOT NULL,
		Field1	TEXT,
		Field2	TEXT,
		Field3	TEXT,
		Field4	TEXT,
		Field5	TEXT,
		Field6	TEXT,
		Field7	TEXT,
		Field8	TEXT,
		Field9	TEXT,
		Field10	TEXT,
		InsertDate DATETIME
	);
	`
	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}

func Store(item Item) {
	sql_additem := `
	INSERT OR REPLACE INTO items(
		Name,
		Field1,
		Field2,
		Field3,
		Field4,
		Field5,
		Field6,
		Field7,
		Field8,
		Field9,
		Field10,
		InsertDate
	) values(?,?,?,?,?,?,?,?,?,?,?,CURRENT_TIMESTAMP)
	`
	sqlStatement, err := db.Prepare(sql_additem)
	if err != nil {
		panic(err)
	}
	defer sqlStatement.Close()

	_, err = sqlStatement.Exec(item.Name, item.Field1, item.Field2, item.Field3, item.Field4,
		item.Field5, item.Field6, item.Field7, item.Field8, item.Field9, item.Field10)
	if err != nil {
		panic(err)
	}
}

func ReadByName(id string) {
	readAllQuery := `SELECT * FROM items WHERE name = ? ORDER BY InsertDate`
	rows, err := db.Query(readAllQuery, id)
	if err != nil { panic(err) }
	defer rows.Close()
	items := make([]Item, 0)
	for rows.Next() {
		item := parseItem(rows)
		items = append(items, item)
		fmt.Println(item.Id)
	}
}

func parseItem(rows *sql.Rows) Item {
	item := Item{}
	err2 := rows.Scan(&item.Id, &item.Name, &item.Field1, &item.Field2, &item.Field3,
		&item.Field4, &item.Field5, &item.Field6,
		&item.Field7, &item.Field8, &item.Field9,
		&item.Field10, &item.Timestamp);
	if err2 != nil { panic(err2) }
	return item;
}

