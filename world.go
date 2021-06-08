package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

var db *gorm.DB
var err error
var	dsn = "host=localhost user=gorm password=localgorm dbname=storyworlds port=5432 sslmode=disable TimeZone=America/Chicago"

type Storyworld struct {
	gorm.Model
	Name string
}

func InitialMigration() {
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	checkErrors(err)
	err = db.AutoMigrate(&Storyworld{})
}

func HandleWorldOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}

func HandleGetAllWorlds(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	checkErrors(err)

	var storyworlds []Storyworld

	db.Find(&storyworlds)
	json.NewEncoder(w).Encode(storyworlds)
}

func HandleGetWorld(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, "get world %s", id)
}

func HandleCreateWorld(w http.ResponseWriter, r *http.Request) {
	s := Storyworld{}
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	checkErrors(err)
	db.Create(&s)
	json.NewEncoder(w).Encode(s)
}

func HandleUpdateWorld(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "update world %s", id)
}

func HandleDeleteWorld(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "delete world %s", id)
}

func checkErrors(e error) {
	if err != nil {
		fmt.Printf(err.Error())
		panic("Could not connect to the db")
	}
}
