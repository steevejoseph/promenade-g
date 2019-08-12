package controllers

import ( 
	"net/http"
	"encoding/json"
	"math/rand"
	"github.com/steevejoseph/promenade-g/models"
	"strconv"
	"github.com/gorilla/mux"
	"log"
)

// User model
type User = models.User

// Person model
type Person = models.Person

var users []User

// CreateUser creates a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = strconv.Itoa(rand.Intn(1e9));
	// users = append(users, user)
	json.NewEncoder(w).Encode(user);
}


// DeleteUser deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	log.Println(params)

	for index, item := range users {
		if item.ID == params["id"]{
			// @steeve: lookup this syntax later on
			users = append(users[:index], users[index+1:]...)
			break
		}
	}	

	json.NewEncoder(w).Encode(users)
}

// GetUser gets users
func GetUsers(w http.ResponseWriter, r *http.Request) {

	users = append(users, 
		User {
			ID: "1",
			FirstName: "John",
			LastName: "Smith", 
			Age: 23,
			Username: "johnsmith",
			Email: "johnsmith@gmail.com",
			Password: "insecure",
			Spouse: 
				Person {
					FirstName: "Jane",
					LastName: "Doe",
				},
					
		});

	log.Println(users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users);
}

// GetUser gets specific user
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	log.Println(params)

	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	// essentially returning empty user
	json.NewEncoder(w).Encode(&User{});
}
