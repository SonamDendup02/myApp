package main

import (
	"goapp/myApp/routes"
)

// func demoHandler(w http.ResponseWriter, r *http.Request) {
// 	p := mux.Vars(r)
// 	course := p["course"]
// 	student:=p["student"]
// 	_, err := w.Write([]byte("hello world.\nThis course is" + course))
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// }\

// func demoHandler(w http.ResponseWriter, r *http.Request) {
// 	p := mux.Vars(r)
// 	course := p["course"]
// 	_, err := w.Write([]byte("hello world.\nThis course is" + course))
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// }

// func main() {
// 	router := mux.NewRouter()
// 	router.HandleFunc("/home/{course}", demoHandler)
// 	err := http.ListenAndServe(":8080", router)
// 	if err != nil {
// 		return
// 	}
// }

func main() {
	// router := mux.NewRouter()
	// router.HandleFunc("/home", tryHandler)
	// // Home page
	// router.HandleFunc("/home/{course}", demoHandler)
	// // course
	// router.HandleFunc("/home/{course}/{student}", studentHandler)
	// // Student
	// err := http.ListenAndServe(":8080", router)
	// if err != nil {
	// 	return
	// }
	// routes.IntializeRoutes()
	routes.IntializeRoutes()
}
