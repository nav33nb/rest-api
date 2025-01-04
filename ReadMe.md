# REST API implementation in Go 

## Item representation
```go
type Book struct {
	Id     int		//primarykey
	Name   string
	Author string
	Price  float64
	Year   int
}
```

## Endpoints
| Endpoint             | Description                 |
| -------------------- | --------------------------- |
| `GET /`              | home, welcome msg           |
| `GET /books`         | fetch all books             |
| `POST /books`        | add a book                  |
| `PUT /books`         | update an existing book     |
| `DELETE /books/{id}` | delete an existing book     |
| `GET /books/{id}`    | fetch a specific book by id |
