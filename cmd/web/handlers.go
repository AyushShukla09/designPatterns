package main

import (
	"fmt"
	"net/http"
)

func (app *Application) HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}
