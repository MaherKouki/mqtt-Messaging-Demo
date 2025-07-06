package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {

	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("goSubscriber")

	messageHandler := func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message on topic %s: %s\n", msg.Topic(), msg.Payload())
	}

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	topic := "chat/maher/test"
	if token := client.Subscribe(topic, 0, messageHandler); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	fmt.Println("Listening for MQTT messages on topic:", topic)
	select {}
}
