package main

import (
	"fmt"

	"github.com/logrusorgru/aurora"
	"github.com/nats-io/nats.go"
)

func subscribe(nc *nats.Conn) error {
	fmt.Println(aurora.Sprintf(aurora.Cyan("Subscribing to nats channel...")))
	
	// Defining nats channel
	var requestSubject = _RequestSubject
	var responseSubject = _ResponseSubject

	// Subscribe to nats channel
	ch := make(chan *nats.Msg, _ChanSize)
	sub, err := nc.ChanSubscribe(requestSubject, ch)
	if err != nil {
		return err
	}

	// Drain subscription
	defer func() {
		if err := sub.Unsubscribe(); err != nil {
			fmt.Println(aurora.Sprintf(aurora.Red("Error unsubscribing: %s"), err))
			panic(err)
		}
		if err := sub.Drain(); err != nil {
			fmt.Println(aurora.Sprintf(aurora.Red("Error draining sub: %s"), err))
			panic(err)
		}
	}()

	// Get process request from the app
	for msg := range ch {
		fmt.Println(aurora.Sprintf(aurora.Green("Received message: %s"), string(msg.Data)))
		data := Data{
			Type:    processRequest,
			Payload: msg.Data,
		}

		// Start processing the request
		var processResult, err = process(&data)
		if err != nil {
			return err
		}

		// Publish the result of the process
		if err := publish(nc, responseSubject, processResult); err != nil {
			return err
		}
		
	}

	return nil
}
