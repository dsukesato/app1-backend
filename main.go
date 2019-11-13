package main

import (
	"github.com/dsukesato/go13/pbl/app1-backend/controllers"
)

func main() {
	controllers.Serve()

	/*http.HandleFunc("/", indexHandler)

	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}*/
	// [END setting_port]
}
