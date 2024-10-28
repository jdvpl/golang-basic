package webserver

import "net/http"

func StartWebServer() {
	// Start the web server
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	// Display the home page
	http.ServeFile(w, r, "./public/index.html")
}
