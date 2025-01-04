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

func (d DbConfig) getConnection() (*pgx.Conn, error) {
	connstring := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?%v", d.Db_user, d.Db_pass, d.Db_addr, d.Db_port, d.Db_name, d.Db_args)
	Log.Trace(connstring)
	conn, err := pgx.Connect(context.Background(), connstring)
	if err != nil {
		return nil, fmt.Errorf("cannot get dbconnection object: %v", err)
	}
	return conn, err
}

func fetchData(db *pgx.Conn, id string) ([]Book, error) {
	books := []Book{}

	query := "select id,name,author,price,year from books"
	if id != "" {
		query += " where id = " + id
	}

	rows, err := db.Query(context.Background(), query)
	if err != nil {
		Log.Error(err)
		return nil, fmt.Errorf("fetch query execution failed for \"%v\" due to %v", query, err)
	}
	defer rows.Close()
	for rows.Next() {
		b := Book{}
		rows.Scan(&b.Id, &b.Name, &b.Author, &b.Price, &b.Year)
		books = append(books, b)
	}
	if len(books) == 0 {
		Log.Warnf("No Matching Data was Found for id = %v", id)
		return nil, err_NoMatch
	}
	Log.Debugf("Retrieved data -> %v", books)
	return books, nil
}

func postData(db *pgx.Conn, b Book) error {
	query := fmt.Sprintf("insert into books (id, author, \"name\", price, \"year\") VALUES(%v, '%v', '%v', %v, %v)", b.Id, b.Author, b.Name, b.Price, b.Year)
	rows, err := db.Exec(context.Background(), query)
	if err != nil {
		Log.Errorf("POST query execution failed for [%v] due to %v", query, err)
		return err
	}
	if rows.RowsAffected() == 0 {
		Log.Errorf("INSERT query succeeded but no new entry was created for [%v]", query)
		return err_NoEffect
	}
	Log.Debugf("%v row was inserted: %#v", rows, b)
	return nil
}

func putData(db *pgx.Conn, b Book) error {
	query := fmt.Sprintf("update books set author='%v', \"name\"='%v', price=%v, \"year\"=%v WHERE id=%v", b.Author, b.Name, b.Price, b.Year, b.Id)
	rows, err := db.Exec(context.Background(), query)
	if err != nil {
		Log.Errorf("UPDATE query execution failed for [%v] due to %v", query, err)
		return err
	}
	if rows.RowsAffected() == 0 {
		Log.Errorf("No match found, no updates were made for id=%v, query=[%v]", b.Id, query)
		return err_NoMatch
	}
	Log.Debugf("%v row was updated: %#v", rows, b)
	return nil
}

func deleteData(db *pgx.Conn, id string) error {
	if id == "" {
		return err_NoId
	}
	query := fmt.Sprintf("delete from books where id=%v", id)
	rows, err := db.Exec(context.Background(), query)
	if err != nil {
		Log.Errorf("DELETE query execution failed for [%v] due to %v", query, err)
		return err
	}
	if rows.RowsAffected() == 0 {
		Log.Errorf("No match found, no deletions were made for id=%v for [%v]", id, query)
		return err_NoEffect
	}
	Log.Debugf("%v row was deleted", rows)
	return nil
}
