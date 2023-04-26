package Model

import "fmt"

type Student struct {
	StdId     int64  `json:"stdid"`
	FirstName string `json:"fname"`
	Lastname  string `json:"lname"`
	Email     string `json:"email"`
}

func (s *Student) Create() error {
	// var stud model.S
	fmt.Println("hi")
	var err error
	return err
}
