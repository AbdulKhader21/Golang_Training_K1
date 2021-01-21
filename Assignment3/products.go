package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	UID   string  `json:"UID"`
	Name  string  `json:"Name"`
	Desc  string  `json:"Desc"`
	Price float64 `json:"Price"`
}

var invertory []Item

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint called: homepage()")
}

func getInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Function Called: getInventory()")

	json.NewEncoder(w).Encode(invertory)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	invertory = append(invertory, item)

	json.NewEncoder(w).Encode(item)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	_deleteItemAtUid(params["uid"])

	json.NewEncoder(w).Encode(invertory)
}

func _deleteItemAtUid(uid string) {
	for index, item := range invertory {
		if item.UID == uid {
			// Delete item from slice
			invertory = append(invertory[:index], invertory[index+1:]...)
			break
		}
	}
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	params := mux.Vars(r)

	// Delete the item at UID
	_deleteItemAtUid(params["uid"])
	// Create it with new data
	invertory = append(invertory, item)

	json.NewEncoder(w).Encode(invertory)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/inventory", getInventory).Methods("GET")
	router.HandleFunc("/inventory", createItem).Methods("POST")
	router.HandleFunc("/inventory/{uid}", deleteItem).Methods("DELETE")
	router.HandleFunc("/inventory/{uid}", updateItem).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	invertory = append(invertory, Item{
		UID:   "0",
		Name:  "Cheese",
		Desc:  "A fine block of cheese",
		Price: 4.99,
	})
	invertory = append(invertory, Item{
		UID:   "1",
		Name:  "Milk",
		Desc:  "A pure bottle of Milk",
		Price: 3.25,
	})
	// TO POST
	// {
	//     "UID": "0",
	//     "Name": "Eggs",
	//     "Desc": "A fresh box of Eggs",
	//     "Price": 2.55
	// }

	handleRequests()
}
