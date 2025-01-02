package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Book struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
	Year   int     `json:"year"`
}

type DbConfig struct {
	Db_user string
	Db_pass string
	Db_addr string
	Db_port string
	Db_name string
	Db_args string
}

func (d DbConfig) getConnection() *pgx.Conn {
	connstring := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?%v", d.Db_user, d.Db_pass, d.Db_addr, d.Db_port, d.Db_name, d.Db_args)
	Log.Trace(connstring)
	conn, err := pgx.Connect(context.Background(), connstring)
	if err != nil {
		Log.Errorf("Unable to connect %v\n", err)
		return nil
	}
	return conn
}

func (app *App) fetchData(id string) ([]Book, error) {
	books := []Book{}

	query := "select id,name,author,price,year from books"
	if id != "all" {
		query += " where id = " + id
	}

	rows, err := app.db.Query(context.Background(), query)
	if err != nil {
		Log.Errorf("Query execution FAILED for \"%v\"\n", query)
		Log.Error(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		b := Book{}
		rows.Scan(&b.Id, &b.Name, &b.Author, &b.Price, &b.Year)
		books = append(books, b)
	}
	Log.Debugf("Retrieved data -> %v", books)
	if len(books) == 0 {
		Log.Warn("No Matching Data was Found for this id=" + id)
	}
	return books, err
}
