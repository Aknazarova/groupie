package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// A Response struct to map the Entire Response
type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

// A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
	Name string `json:"name"`
}

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	b, err := json.Marshal(group.Name)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

	// responseData, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var responseObject Response
	// json.Unmarshal(responseData, &responseObject)

	// // fmt.Println(responseObject.Name)
	// // fmt.Println(len(responseObject.Pokemon))

	// for i := 0; i < len(responseObject.Pokemon); i++ {
	// 	fmt.Println(responseObject.Pokemon[i].Species.Name)
	// }

}

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	type ColorGroup struct {
// 		ID     int
// 		Name   string
// 		Colors []string
// 	}
// 	group := ColorGroup{
// 		ID:     1,
// 		Name:   "Reds",
// 		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
// 	}
// 	b, err := json.Marshal(group)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	os.Stdout.Write(b)
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// )

// func main() {

// 	slcD := []string{"apple", "peach", "pear"}
// 	slcB, _ := json.Marshal(slcD)

// 	http.HandleFunc("/artists", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, "About Page")
// 	})

// 	fmt.Println(string(slcB))
// 	fmt.Println("Server is listening...")
// 	http.ListenAndServe(":8080", nil)
// }
