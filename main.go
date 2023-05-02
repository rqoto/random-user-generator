package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		U, err := json.Marshal(CreateUser())
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Write(U)
	})
	http.ListenAndServe(":3000", r)
}
