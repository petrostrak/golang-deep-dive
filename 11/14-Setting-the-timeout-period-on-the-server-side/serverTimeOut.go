package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>\n", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

// You need the timeout period on the server side because there are
// times when clients take much longer than expected to end an HTTP
// connection. This is usually happens for two reasons. The first is
// bugs in the client software and the second is when a server process
// is experiencing an attack!
func main() {

	PORT := ":8080"
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Printf("Listening on http://0.0.0.0%s\n", PORT)
	} else {
		PORT = ":" + arguments[1]
		fmt.Printf("Listening on http://0.0.0.0%s\n", PORT)
	}

	mux := http.NewServeMux()

	// In this case we are using an http.Server{} structure that supports
	// two kinds of timeouts with its fields. The value of the ReadTimeout
	// specifies the maximum duration allowed to read the entire request,
	// including the body. The value of WriteTimeout specifies the maximum
	// duration before timing out the writing of the response. Put simply
	// this is the time from the end of the request header read to the
	// end of the response write.
	srv := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	mux.HandleFunc("/time", timeHandler)
	mux.HandleFunc("/", myHandler)

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err)
		return
	}

}
