

package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Nuurs/go-language/pkg"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/heroes", hero.GetHeroesListHandler).Methods("GET")
	r.HandleFunc("/heroes/{name}", hero.GetHeroHandler).Methods("GET")
	r.HandleFunc("/health", hero.HealthCheckHandler).Methods("GET")

	http.Handle("/", r)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
