package main

type Book struct {
	Id     int
	Name   string
	Author string
	Price  float64
	Year   int
}

var Books = []Book{
	{Id: 1, Name: "The Great Gatsby", Author: "F. Scott Fitzgerald", Price: 10.99, Year: 1925},
	{Id: 2, Name: "1984", Author: "George Orwell", Price: 8.99, Year: 1949},
	{Id: 3, Name: "To Kill a Mockingbird", Author: "Harper Lee", Price: 12.50, Year: 1960},
	{Id: 4, Name: "The Catcher in the Rye", Author: "J.D. Salinger", Price: 9.75, Year: 1951},
	{Id: 5, Name: "Moby-Dick", Author: "Herman Melville", Price: 15.20, Year: 1851},
	{Id: 6, Name: "Pride and Prejudice", Author: "Jane Austen", Price: 7.99, Year: 1813},
	{Id: 7, Name: "The Hobbit", Author: "J.R.R. Tolkien", Price: 14.99, Year: 1937},
	{Id: 8, Name: "War and Peace", Author: "Leo Tolstoy", Price: 19.99, Year: 1869},
	{Id: 9, Name: "The Alchemist", Author: "Paulo Coelho", Price: 11.45, Year: 1988},
	{Id: 10, Name: "The Road", Author: "Cormac McCarthy", Price: 13.60, Year: 2006},
}
