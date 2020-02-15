package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		for k := range v {
			if k != "Name" {
				delete(v, k)
			}
		}
		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
	}
}

// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {

// 	http.HandleFunc("/artists", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, "About Page")
// 	})
// 	http.HandleFunc("/locations", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, "Contact Page")
// 	})
// 	http.HandleFunc("/dates", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, "Index Page")
// 	})
// 	fmt.Println("Server is listening...")
// 	http.ListenAndServe(":8080", nil)
// }
