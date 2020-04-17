package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// get the config & build the server

	// build the router
	r := mux.NewRouter()

	// setup report server...
	// should serve the file contents inside a
	// specific <pre> wrapper based on config
	setupFileHandler("./report", "/report/", r)
	// setup email server
	// should server only the .htm files from
	// folder and ignore the others (esp. .eml)
	setupFileHandler("./email", "/email/", r)

	// setup the application server
	// wraps it all together
	r.HandleFunc("/", handleApp).Methods("GET", "POST", "PUT")

	portNo := 3000

	portString := fmt.Sprintf(":%v", portNo)
	log.Printf("listening on %v...\n", portString)
	err := http.ListenAndServe(portString, r)
	if err != nil {
		log.Fatal(err)
	}
}

func setupFileHandler(path string, prefix string, router *mux.Router) {
	// check prefix for / at start & end?
	//webRoute := fmt.Sprintf("%v/", prefix)
	/*
		fs := http.FileServer(http.Dir(path))
		//http.Handle(webRoute, fs)
		mux.Handle(webRoute, http.StripPrefix(prefix, fileRouter(fs)))
	*/
	router.PathPrefix(prefix).Handler(
		http.StripPrefix(prefix, http.FileServer(http.Dir(path))))
}

/*
func fileRouter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
*/
func handleApp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Building the application!")
}
