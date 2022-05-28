package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kardenvan7/my_expenses_planner_backend_app/database"
	"github.com/kardenvan7/my_expenses_planner_backend_app/pages"
)

func main() {
	db, err := database.ConnectToDb()

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", pages.HomePage)
	http.HandleFunc(
		"/categories",
		func(responseWriter http.ResponseWriter, request *http.Request) {
			pages.CategoriesPage(db, responseWriter, request)
		},
	)
	http.ListenAndServe(":8080", nil)
}
