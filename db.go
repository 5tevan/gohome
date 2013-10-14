package main

import(
	"fmt"
)

type DB struct {
	username string
	password string
	database string
	host string
	queries chan string
	//hat about  results channel for communication the other way
}

func NewDB() *DB {
	d := &DB{
		username:"foo",
		password:"bar",
		database:"jah",
		host:"brony",
	}

	go d.queryLoop()
	return d
}

// infinite loop to read queries out of a channel and execute them
func (d *DB) queryLoop() {
	for {
		q := <-d.queries
		fmt.Println(q) // replace with an actual query
	}
}