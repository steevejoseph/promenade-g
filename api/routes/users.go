package routes

import (
	"github.com/gorilla/mux"
	"github.com/steevejoseph/promenade-g/models"
	"github.com/steevejoseph/promenade-g/api/controllers"
)

// User model
type User = models.User 
// Person model
type Person = models.Person

func SetupHandlers(r *mux.Router){
	r.HandleFunc("/get-users", controllers.GetUsers).Methods("GET");
	r.HandleFunc("/get-user/{id}", controllers.GetUser).Methods("GET");
	r.HandleFunc("/create-user/", controllers.CreateUser).Methods("POST");
	r.HandleFunc("/delete-user/{id}", controllers.DeleteUser).Methods("DELETE");
}


