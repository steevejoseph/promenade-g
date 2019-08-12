package models

// User struct(model)
type User struct {
	ID string `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Age int `json:"age"`
	Username string `json:"username"`
	Email  string `json:"email"`
	Password string `json:"password"`
	Spouse Person `json:"spouse"`
}

// Person struct(model)
type Person struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}