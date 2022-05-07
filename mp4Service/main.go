package main

import (
	"fmt"

	"github.com/logrusorgru/aurora"
	"github.com/nats-io/nats.go"
)

func main() {
	fmt.Println(aurora.Sprintf(aurora.Green("Starting mp4Service...")))

	// Defining nats server
	var natsServer = _Url

	// Connect to the server
	nc, err := nats.Connect(natsServer)
	if err != nil {
		fmt.Println(aurora.Sprintf(aurora.Red("Could not connect to nats server: %s"), err))
		return
	}

	// Drain connection
	defer func() {
		if err := nc.Drain(); err != nil {
			fmt.Println(aurora.Sprintf(aurora.Red("Error draining nats connection: %s"), err))
			panic(err)
		}
	}()

	// Subscribe to nats channel
	if err := subscribe(nc); err != nil {
		fmt.Println(aurora.Sprintf(aurora.Red("Could not subscribe to nats channel: %s"), err))
		return
	}
}
