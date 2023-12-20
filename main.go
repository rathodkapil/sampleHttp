package main

import (
	"encoding/json"
	"net/http"
	"sort"

	"github.com/gorilla/mux"
)

type Person struct {
	Name  string
	Age   int
	Phone string
}

var Persons []Person

func myHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("say Hello"))
}

func myMuxHandler(w http.ResponseWriter, req *http.Request) {
	p1 := Person{"kapil", 20, "12221"}
	p2 := Person{"DarshitA", 20, "12221"}
	p3 := Person{"pAVIT", 20, "12221"}
	p4 := Person{"Mohit", 20, "12221"}
	p5 := Person{"Mummy", 20, "12221"}

	Persons = append(Persons, p1)
	Persons = append(Persons, p2)
	Persons = append(Persons, p3)
	Persons = append(Persons, p4)
	Persons = append(Persons, p5)

	sort.Slice(Persons, func(i, j int) bool {
		return Persons[i].Name < Persons[j].Name
	})

	b, _ := json.Marshal(Persons)
	w.Write(b)
	//w.Write([]byte("myMux Handler"))
}

func main() {
	// http.HandleFunc("/greetings/", myHandler)
	// http.ListenAndServe(":8080", nil)

	myMux := mux.NewRouter()
	myMux.HandleFunc("/myMux", myMuxHandler)
	http.ListenAndServe(":8080", myMux)

}
