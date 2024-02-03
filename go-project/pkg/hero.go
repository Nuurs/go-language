

package hero

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Nuurs/go-language/api"
	"fmt"
)

func GetHeroesListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(api.GetHeroesList()) 
}


func GetHeroHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	heroName := params["name"]

	if hero := api.GetHeroByName(heroName); hero != nil {
		json.NewEncoder(w).Encode(hero)
	} else {
		http.Error(w, "Hero not found", http.StatusNotFound)
	}
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "Super Hero, Author is Nursultan")
}
