package database

import (
	"database/sql"
	"fmt"
	"log"
)

type Options struct {
	Name       string
	Create     bool
	Drop       bool
	Initialise bool
	URI        string
}

type DB struct {
	conn    *sql.DB
	options Options
}

func New(options Options) *DB {
	_db, err := sql.Open("postgres", options.URI)
	if err != nil {
		panic(err)
	}
	log.Println("connected to database:", options.URI)

	return &DB{
		conn:    _db,
		options: options,
	}
}

func (db *DB) Create() {
	if !db.options.Create {
		return
	}
	_, err := db.conn.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", db.options.Name))
	if err != nil {
		panic(err)
	}
}

func (db *DB) Drop() {
	if !db.options.Drop {
		return
	}
	_, err := db.conn.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", db.options.Name))
	if err != nil {
		panic(err)
	}
}

func (db *DB) Initialise() {
	if !db.options.Initialise {
		return
	}
	for _, schema := range schemas {
		_, err := db.conn.Exec(schema.Schema)
		if err != nil {
			panic(fmt.Sprintf("Failed to create %s: %v", schema.Name, err))
		}
	}
}
