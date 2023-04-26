package Controller

import (
	"encoding/json"
	"fmt"
	"goapp/myApp/Model"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	course := p["course"]
	_, err := w.Write([]byte("Hello World.\nThis course is" + course))
	if err != nil {
		fmt.Println("error:", err)
	}
}
func AddStudent(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "add student handler")
	// router.HandleFunc("/student", controller.AddStudent)
	var stud Model.Student
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&stud); err != nil {
		w.Write([]byte("Invalid json data"))
		return
	}
	defer r.Body.Close()

	saveErr := stud.Create()
	if saveErr != nil {
		w.Write([]byte("Database error"))
		return
	}

	w.Write([]byte("Response success."))
}

func AddCourse(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	course := p["course"]
	_, err := w.Write([]byte("welcome :" + course))
	_ = err
}
