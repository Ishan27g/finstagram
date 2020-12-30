package routes

import (
	"appserver/v0/controllers"
	`appserver/v0/models`
	`encoding/json`
	"log"
	`net/http`

	"github.com/gorilla/mux"
)

// AddApproutes will add the routes for the application
func AddApproutes(route *mux.Router) {


	route.HandleFunc("/signin", controllers.SignInUser).Methods("POST")
	route.HandleFunc("/signup", controllers.SignUpUser).Methods("POST", "OPTIONS")
	route.HandleFunc("/userDetails", controllers.GetUserDetails).Methods("GET")
	route.HandleFunc("/up", up).Methods("POST", "OPTIONS")
	log.Println("Routes initialised")
}
 func up(response http.ResponseWriter, request *http.Request){

 	var data models.RegistationParams
 	decoder := json.NewDecoder(request.Body)
 	err := decoder.Decode(&data)

    if err != nil {
 	    panic(err)
    }
    log.Println(data)
 }