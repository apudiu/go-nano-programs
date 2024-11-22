package main

import "log"

func main() {
	port := 3000
	srv, err := NewServer("localhost", Config{
		Port: &port,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Serving on %s", srv.Addr)

	log.Fatal(srv.ListenAndServe())
}
