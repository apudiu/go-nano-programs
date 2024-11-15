package main

import "fmt"

func main() {
	logging := false

	var client *string
	if logging {
		client, err := createClientWithLogging()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(client)
	} else {
		client, err := createClient()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(client)
	}

	fmt.Printf("client is; %#v \n", client)
}

func createClientWithLogging() (*string, error) {
	s := "client with logging"
	return &s, nil
}

func createClient() (*string, error) {
	s := "client without logging"
	return &s, nil
}
