# mqtt-Messaging-Demo
This project demonstrates a simple one-way messaging system using the MQTT protocol. A Spring Boot application publishes messages to a public MQTT broker (HiveMQ), and a Go client subscribes to the same topic and prints received messages. This example is useful for learning how to integrate Spring Boot and Go using lightweight IoT messaging protocols like MQTT.
# MQTT Messaging Demo with Spring Boot and Go

This project demonstrates a one-way MQTT messaging setup using:
- ☕ Spring Boot (as MQTT Publisher)
- 🐹 Go (as MQTT Subscriber)
- 🌐 MQTT Broker: [HiveMQ Public Broker](https://www.hivemq.com/public-mqtt-broker/)

## How it works

- Spring Boot app exposes a REST API to publish messages to a topic.
- Go app subscribes to the topic and prints messages it receives.

## Usage

1. Run Go subscriber:
```bash
cd go-subscriber
go run goSubscriber.go

