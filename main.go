package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	booksAPI := app.Party("/book")
	booksAPI.Use(iris.Compression)
	booksAPI.Get("/", list)
	booksAPI.Post("/", create)

	hj := app.Party("/highjack")
	hj.Use(iris.Compression)
	hj.Get("/", highjacker.jd)
	app.Listen(":8081")
}

//Book example
type Book struct {
	Title string `json:"title"`
}

func list(ctx iris.Context) {
	books := []Book{
		{"Mastering Concurrency in Go."},
		{"Go design patterns."},
		{"Black Hat go"},
	}
	ctx.JSON(books)
}

func create(ctx iris.Context) {
	var b Book
	err := ctx.ReadJSON(&b)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title("Book creation failure").DetailErr(err))
		return
	}
	fmt.Println("Received book: " + b.Title)
	ctx.StatusCode(iris.StatusCreated)
}
