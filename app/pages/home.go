package pages

import (
	"fmt"
	"net/http"
)

func HomePage(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "<p>Home Page</p>")
}
