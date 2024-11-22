package main

import "log"

func main() {

	// build config
	srv, err := NewServer("localhost", WithPort(4000)) // can add more options
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Serving on %s", srv.Addr) // dsaf

	log.Fatal(srv.ListenAndServe())
}
