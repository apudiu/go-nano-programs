package main

import "log"

func main() {

	// build config
	builder := ConfigBuilder{}
	builder.Port(8000)
	cfg, err := builder.Build()
	if err != nil {
		log.Fatal(err)
	}

	srv, err := NewServer("localhost", cfg)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Serving on %s", srv.Addr)

	log.Fatal(srv.ListenAndServe())
}
