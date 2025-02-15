package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/brigadecore/brigade/sdk/v2"
	"github.com/brigadecore/brigade/sdk/v2/core"
	"github.com/brigadecore/brigade/sdk/v2/restmachinery"
)

func main() {
	ctx := context.Background()

	// Get the Brigade API server address and token from the environment
	apiServerAddress := os.Getenv("APISERVER_ADDRESS")
	if apiServerAddress == "" {
		log.Fatalf("Required environment variable APISERVER_ADDRESS not found.\n")
	}
	gatewayToken := os.Getenv("API_TOKEN")
	if gatewayToken == "" {
		log.Fatalf("Required environment variable API_TOKEN not found.\n")
	}

	// The default Brigade deployment mode uses self-signed certs
	// Hence, we allow insecure connections in our APIClientOptions
	// This can be changed to false if insecure connections should not be allowed
	apiClientOpts := &restmachinery.APIClientOptions{
		AllowInsecureConnections: true,
	}

	// Create an API client with the gateway token value
	client := sdk.NewAPIClient(
		apiServerAddress,
		gatewayToken,
		apiClientOpts,
	)

	// Construct a Brigade Event
	event := core.Event{
		// This is the source value for this event
		Source: "example.org/example-gateway",
		// This is the event's type
		Type: "hello",
		// This is the event's payload
		Payload: "Dolly",
	}

	// Create the Brigade Event
	events, err := client.Core().Events().Create(ctx, event)
	if err != nil {
		log.Fatal(err)
	}

	// If the returned events list has no items, no event was created, most
	// likely because there is no project subscribing to an event of this type
	if len(events.Items) != 1 {
		fmt.Println("No event was created.")
		return
	}

	// The Brigade event was successfully created!
	fmt.Printf("Event created with ID %s\n", events.Items[0].ID)
}
