// Start an HTTP GraphQL API server, which is loaded multiple Databases for serving
package main

import (
	"log"
	"net/http"
	"time"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/tonyghita/graphql-go-example/handler"
	"loader"
)

func main(){
	// Tweakable Arugments
	var (
		port = ":8000"
		readHeaderTimeout = 1 * time.Second
		writeTimeOut = 10 * time.Second
		idleTimeout = 90 * time.Second
		maxHeaderBytes = http.DefaultMaxHeaderBytes
	)

	log.SetFlags(log.Lshortfile | log. LstdFlags)

	// Register handlers to routes.
	mux := http.NewServeMux()
	mux.Handle("/", handler.GraphiQL{})
	mux.Handle("/graphql/", h)
	mux.Handle("/graphql", h) // Register without a trailing slash to avoid redirect.

	// Configure the HTTP server
	
	server := &http.Server{
		Addr: port,
		Handler: mux,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout: writeTimeOut,
		IdleTimeout: idleTimeout,
		MaxHeaderBytes: maxHeaderBytes
	}

	// Begin listing for requests.
	log.Printf("Listing for requests on %s", server.Addr)

	if err = server.ListenAndServe(); err != nil{
		log.Println("server.ListenAndServer:", err)
	}

	log.Println("Shut down.")
}
