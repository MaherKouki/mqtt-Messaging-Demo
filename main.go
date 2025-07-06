package main

import (
	"log"

	"github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/listeners"
)

func main() {
	// Create a new MQTT broker instance
	broker := mqtt.New(nil)

	// Create a TCP listener on port 1883
	tcpConfig := listeners.Config{
		ID:   "tcp1",
		Bind: ":1883",
	}

	// Create the TCP listener with new config format
	tcp := listeners.NewTCP(tcpConfig)
	err := broker.AddListener(tcp)
	if err != nil {
		log.Fatalf("Failed to add TCP listener: %v", err)
	}

	log.Println("🟢 MQTT broker started on port 1883")

	// Start the broker (blocking)
	err = broker.Serve()
	if err != nil {
		log.Fatalf("Broker error: %v", err)
	}
}
