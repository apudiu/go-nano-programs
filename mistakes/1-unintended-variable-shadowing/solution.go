package main

import "fmt"

func main() {
	logging := false

	var client *string
	var err error

	if logging {
		client, err = createClientWithLogging2()
	} else {
		client, err = createClient2()
	}

	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Printf("client is; %#v \n", client)
}

func createClientWithLogging2() (*string, error) {
	s := "client with logging"
	return &s, nil
}

func createClient2() (*string, error) {
	s := "client without logging"
	return &s, nil
}
