package main

import (
	"net/http"
	"fmt"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
    js := `{"status": "available", "environment": %q}`
    js = fmt.Sprintf(js, app.env)
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(js))
}
