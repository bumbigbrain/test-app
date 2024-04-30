package main


import (
	"log"
	"fmt"
	"net/http"
	"encoding/json"
)


type Product struct {
	Name string
}


func main() {
	http.HandleFunc("/getNameFromDatabase", handlerGetNameFromDatabase)
	http.HandleFunc("/createBoxToDatabase", handlerCreateBoxToDatabase)
	log.Fatal(http.ListenAndServe(":8081", nil))
}


func handlerGetNameFromDatabase(w http.ResponseWriter, r *http.Request) {	

	p := Product{
		Name: "kyu",
	}

	jsonData, _ := json.Marshal(p)
	fmt.Println("handlerGetNameFromDatabase was invoked")
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
		
}




func handlerCreateBoxToDatabase(w http.ResponseWriter, r *http.Request) {	

	fmt.Println("handlerCreateBoxFromDatabase was invoked")
		
}