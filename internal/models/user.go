package models

type User struct {
	ID    int    `json:"id" type:"int"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}
