package main

import ( 
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/steevejoseph/promenade-g/api/routes"
)

func main() {
	r := mux.NewRouter()
	routes.SetupHandlers(r);
	// @steeve: find out how to conv to process.env.PORT
	log.Println("Starting server on port 8080");
	log.Fatal(http.ListenAndServe(":8080", r));
}