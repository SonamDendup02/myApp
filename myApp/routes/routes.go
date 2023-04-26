package routes

import (
	"log"
	"myApp/Controller"
	"net/http"

	"github.com/gorilla/mux"
)

func IntializeRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/", Controller.HomeHandler)

	router.HandleFunc("/Students", Controller.AddStudent)
	router.HandleFunc("/Student/{Course}", Controller.AddCourse)

	log.Println("Application running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
