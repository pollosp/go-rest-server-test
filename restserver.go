package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var tpl *template.Template

type Payload struct {
	Stuff Data
}

type Data struct {
	Fruit  Fruits
	Veggie Vegetables
}

type Fruits map[string]int
type Vegetables map[string]int

//Init is the first function called even before main
func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func index(w http.ResponseWriter, r *http.Request) {
	useragent := r.UserAgent()
	log.Println(useragent)
	response, err := getJSONResponse()
	if err != nil {
		log.Fatal("Error Fatal", err)
		panic(err)
	}
	fmt.Fprintf(w, string(response))
}

func parameterTest(w http.ResponseWriter, r *http.Request) {
	useragent := r.UserAgent()
	params := mux.Vars(r)
	response := params["id"]
	log.Println(useragent)
	fmt.Fprintf(w, string(response))
}

func templateTest(w http.ResponseWriter, r *http.Request) {
	useragent := r.UserAgent()
	params := mux.Vars(r)
	text := params["text"]
	log.Println(useragent)

	tpl.ExecuteTemplate(w, "test.gohtml", text)

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/params/{id}", parameterTest).Methods("GET")
	router.HandleFunc("/template/{text}", templateTest).Methods("GET")
	err := http.ListenAndServe("0.0.0.0:2525", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Println("Running the server")
}

func getJSONResponse() ([]byte, error) {
	fruits := make(map[string]int)
	fruits["Apples"] = 25
	fruits["Melon"] = 30
	fruits["Orange"] = 100

	vegetables := make(map[string]int)
	vegetables["Cucumber"] = 10
	vegetables["Pepper"] = 35
	d := Data{fruits, vegetables}
	p := Payload{d}
	return json.MarshalIndent(p, "", "  ")
}
