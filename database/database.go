package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Database interface {
	Open() error
	Close() error
	GetDb() *sql.DB
}

type DatabaseStruct struct {
	db *sql.DB
}

func (dbs *DatabaseStruct) Open() error {
	log.Println("Open connection to DB")
	tempDb, err := sql.Open("mysql", connectionStr)

	if err != nil {
		log.Println(err)
		return err
	}

	dbs.db = tempDb
	log.Println("Connected to DB")
	return err
}

func (dbs *DatabaseStruct) Close() error {
	err := dbs.db.Close()

	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Close connection to db")
	return err
}

func (dbs *DatabaseStruct) GetDb() *sql.DB {
	return dbs.db
}
