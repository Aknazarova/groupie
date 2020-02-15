package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type code struct {
	Code    int
	Descrip string
}

func main() {
	var data []code

	rcvd := `[{"Code":200,"Descrip":"StatusOK"},{"Code":301,"Descrip":"StatusMovedPermanently"},{"Code":302,"Descrip":"StatusFound"},{"Code":303,"Descrip":"StatusSeeOther"},{"Code":307,"Descrip":"StatusTemporaryRedirect"},{"Code":400,"Descrip":"StatusBadRequest"},{"Code":401,"Descrip":"StatusUnauthorized"},{"Code":402,"Descrip":"StatusPaymentRequired"},{"Code":403,"Descrip":"StatusForbidden"},{"Code":404,"Descrip":"StatusNotFound"},{"Code":405,"Descrip":"StatusMethodNotAllowed"},{"Code":418,"Descrip":"StatusTeapot"},{"Code":500,"Descrip":"StatusInternalServerError"}]`

	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range data {
		fmt.Println(v.Code, "-", v.Descrip)
	}
}

// package main

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// )

// type person struct {
// 	Fname string
// 	Lname string
// 	Items []string
// }

// func main() {
// 	http.HandleFunc("/", foo)
// 	http.HandleFunc("/mshl", mshl)
// 	http.HandleFunc("/encd", encd)
// 	http.ListenAndServe(":8888", nil)
// }

// func foo(w http.ResponseWriter, req *http.Request) {
// 	s := `<!DOCTYPE html>
// 		<html lang="en">
// 		<head>
// 		<meta charset="UTF-8">
// 		<title>FOO</title>
// 		</head>
// 		<body>
// 		You are at foo
// 		</body>
// 		</html>`
// 	w.Write([]byte(s))
// }

// func mshl(w http.ResponseWriter, req *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	p1 := person{
// 		Fname: "James",
// 		Lname: "Bond",
// 		Items: []string{"Suit", "Gun", "Wry sense of humor"},
// 	}
// 	j, err := json.Marshal(p1)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	w.Write(j)
// }

// func encd(w http.ResponseWriter, req *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	p1 := person{
// 		Fname: "James",
// 		Lname: "Bond",
// 		Items: []string{"Suit", "Gun", "Wry sense of humor"},
// 	}
// 	err := json.NewEncoder(w).Encode(p1)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }
