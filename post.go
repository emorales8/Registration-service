package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Car struct {
	ID         int    `json:"id"`
	Brand      string `json:"brand"`
	Model      string `json:"model"`
	HorsePower string `json:"horsepower"`
}

var db *sql.DB
var err error

func connect() {
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Unable to load env file! %v", err)
	}
	fmt.Println(os.Getenv("DB_HOST"))
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err = sql.Open("mysql", DBURL)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully Connected to mymariadb database")
}

func insertDB(w http.ResponseWriter, r *http.Request) {
	connect()

	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("INSERT INTO cars (id, brand, model, horse_power) VALUES (?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfully Inserted into my DB")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	var car Car

	json.Unmarshal(body, &car)
	car.ID = (rand.Intn(10000))

	//car.HorsePower, err = strconv.Atoi(keyVal["horsepower"])  cuando era un int

	_, err = stmt.Exec(car.ID, car.Brand, car.Model, car.HorsePower)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New post was created")
	fmt.Println("Successfully Inserted to myMariadb database")

	json.NewEncoder(w).Encode(car)

	defer db.Close()
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/service/v1/cars", insertDB).Methods("POST")

	//pass in our newly created router as the second argument
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("You are using my-post-go-app!")
	handleRequests()
}
