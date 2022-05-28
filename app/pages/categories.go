package pages

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kardenvan7/my_expenses_planner_backend_app/models"
)

func CategoriesPage(db *sql.DB, responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("Getting category")

	result, err := db.Query("SELECT * FROM categories;")

	if err != nil {
		log.Fatal(err)
	}

	var category models.Category

	if result.Next() {
		err = result.Scan(
			&category.ID,
			&category.UUID,
			&category.Name,
			&category.Color,
		)

		if err != nil {
			log.Fatal(err)
		}

		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusCreated)
		json.NewEncoder(responseWriter).Encode(category)
	}
}
