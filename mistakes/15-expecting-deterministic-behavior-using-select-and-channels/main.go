package main

import "fmt"

func main() {

	msgChan := make(chan string, 4) // this can be guaranteed with 'noMsgDeliveryGuarantee' when unbuffered
	disconnectChan := make(chan string)
	defer close(msgChan)
	defer close(disconnectChan)

	go func() {
		//noMsgDeliveryGuarantee(msgChan, disconnectChan)
		msgDeliveryGuarantee(msgChan, disconnectChan)
	}()

	for i := 0; i < 10; i++ {
		msgChan <- fmt.Sprintf("Message #%d", i)
	}

	disconnectChan <- ""
}

func noMsgDeliveryGuarantee(msgChan <-chan string, disconnCh <-chan string) {
	for {
		select {
		case msg := <-msgChan:
			fmt.Printf("%#v \n", msg)
		case <-disconnCh:
			fmt.Println("disconnected! channel closed")
			return
		}
	}
}

func msgDeliveryGuarantee(msgChan <-chan string, disconnCh <-chan string) {
	for {
		select {
		case msg := <-msgChan:
			fmt.Printf("%#v \n", msg)
		case <-disconnCh:

			// when disconnecting, check for remaining messages
			for {
				select {
				case msg := <-msgChan:
					fmt.Printf("-> %#v \n", msg)
				default:
					fmt.Println("disconnected! channel closed")
					return
				}
			}

		}
	}
}
