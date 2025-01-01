# REST API implmentation in Go 

## Item representation
```go
type Book struct {
	Id     int
	Name   string
	Author string
	Price  float64
	Year   int
}
```

## Endpoints
- `GET /books` list of all books
- `GET /book/id` list single book
- `POST /book` add a book
- `PUT /book/id` update(overwrite) a book
- `DELETE /book/id` delete a book