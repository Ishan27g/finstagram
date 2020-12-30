package main

import (
	`appserver/v0/controllers`
	`appserver/v0/routes`
	`fmt`
	"github.com/gorilla/mux"
	`log`
	`net/http`
)

// CORS Middleware
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set headers
		w.Header().Set("Access-Control-Allow-Headers", "origin, content-type")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		fmt.Println("server hit!")

		// Next
		next.ServeHTTP(w, r)
		return
	})
}

func main() {
	// Connect to MongoDB
	controllers.ConnectDatabase()
	if controllers.Client == nil {

	}

	router := mux.NewRouter()
	router.Use(CORS)
	// Add routes and handler
	routes.AddApproutes(router)

	log.Println("Server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
	/*log.Fatal(http.ListenAndServe(":8000", handlers.CORS(
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(router)))
	*/

}