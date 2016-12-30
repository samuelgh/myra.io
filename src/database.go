package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"fmt"
)

var db *sql.DB

type Item struct {
	Id        string `json:"id"`
	Field1    string `json:"field1,omitempty"`
	Field2    string `json:"field2,omitempty"`
	Field3    string `json:"field3,omitempty"`
	Field4    string `json:"field4,omitempty"`
	Field5    string `json:"field5,omitempty"`
	Field6    string `json:"field6,omitempty"`
	Field7    string `json:"field7,omitempty"`
	Field8    string `json:"field8,omitempty"`
	Field9    string `json:"field9,omitempty"`
	Field10   string `json:"field10,omitempty"`
	Timestamp string `json:"timestamp"`
}

type ItemMeta struct {
	Id      string `json:"id"`
	Field1  string `json:"field1,omitempty"`
	Field2  string `json:"field2,omitempty"`
	Field3  string `json:"field3,omitempty"`
	Field4  string `json:"field4,omitempty"`
	Field5  string `json:"field5,omitempty"`
	Field6  string `json:"field6,omitempty"`
	Field7  string `json:"field7,omitempty"`
	Field8  string `json:"field8,omitempty"`
	Field9  string `json:"field9,omitempty"`
	Field10 string `json:"field10,omitempty"`
	Timestamp string `json:"created"`
}

type ResultData struct {
	Meta ItemMeta `json:"meta"`
	Data []Item  `json:"data"`
}

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./db.db")
	if ( err != nil) {
		log.Fatal(err)
	}
	createMetaTable();
}

func StoreMetaData(itemMeta ItemMeta) int64 {
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
	) values(?, ?,?,?,?,?,?,?,?,?,?,CURRENT_TIMESTAMP)`
	sqlStatement, err := db.Prepare(sql_additem)
	if err != nil {
		panic(err)
	}
	defer sqlStatement.Close()

	result, err := sqlStatement.Exec(itemMeta.Id, itemMeta.Field1, itemMeta.Field2, itemMeta.Field3, itemMeta.Field4,
		itemMeta.Field5, itemMeta.Field6, itemMeta.Field7, itemMeta.Field8, itemMeta.Field9, itemMeta.Field10)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	return id
}

func CreateTableWithName(name string) {
	table := "CREATE TABLE IF NOT EXISTS " + name +
		`(Field1	TEXT,
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
	);`
	_, err := db.Exec(table)
	if err != nil {
		panic(err)
	}
}

func createMetaTable() {
	table := `CREATE TABLE IF NOT EXISTS items(
	Name TEXT,
	Field1 TEXT,
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
	);`
	_, err := db.Exec(table)
	if err != nil {
		panic(err)
	}
}

func Store(item Item, name string) int64 {
	sql_additem := `
	INSERT OR REPLACE INTO ` + name + `(
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
	) values(?,?,?,?,?,?,?,?,?,?,CURRENT_TIMESTAMP)`
	sqlStatement, err := db.Prepare(sql_additem)
	if err != nil {
		panic(err)
	}
	defer sqlStatement.Close()

	result, err := sqlStatement.Exec(item.Field1, item.Field2, item.Field3, item.Field4,
		item.Field5, item.Field6, item.Field7, item.Field8, item.Field9, item.Field10)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	return id
}
func ReadMetaByName(id string) ItemMeta {
	query := `SELECT * FROM items WHERE name = ?`
	row := db.QueryRow(query, id)
	result := parseItemMeta(row);
	return result
}
func ReadByName(id string) []Item {
	readAllQuery := `SELECT _ROWID_,* FROM ` + id + ` ORDER BY InsertDate`
	rows, err := db.Query(readAllQuery, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	items := make([]Item, 0)
	for rows.Next() {
		item := parseItem(rows)
		items = append(items, item)
		fmt.Println(item.Id)
	}
	return items
}

func parseItemMeta(row *sql.Row) ItemMeta {
	item := ItemMeta{}
	err2 := row.Scan(&item.Id, &item.Field1, &item.Field2, &item.Field3,
		&item.Field4, &item.Field5, &item.Field6,
		&item.Field7, &item.Field8, &item.Field9,
		&item.Field10, &item.Timestamp);
	if err2 != nil {
		panic(err2)
	}
	return item;
}

func parseItem(rows *sql.Rows) Item {
	item := Item{}
	err2 := rows.Scan(&item.Id, &item.Field1, &item.Field2, &item.Field3,
		&item.Field4, &item.Field5, &item.Field6,
		&item.Field7, &item.Field8, &item.Field9,
		&item.Field10, &item.Timestamp);
	if err2 != nil {
		panic(err2)
	}
	return item;
}

