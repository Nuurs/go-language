
package api


type Hero struct {
	Name  string `json:"name"`
	Alias string `json:"alias"`
	Age   int    `json:"age"`
	Power string `json:"power"`
}

var heroesList = []Hero{
	{"Captain America", "Steve Rogers", 100, "Super Strength"},
	{"Iron Man", "Tony Stark", 45, "Genius, Billionaire, Playboy"},
	{"Wonder Woman", "Diana Prince", 3000, "Superhuman Strength"},
	{"Spider-Man", "Peter Parker", 25, "Wall-Crawling"},
}

func GetHeroesList() []Hero {
	return heroesList
}

func GetHeroByName(name string) *Hero {
	for _, hero := range heroesList {
		if hero.Name == name {
			return &hero
		}
	}
	return nil
}
