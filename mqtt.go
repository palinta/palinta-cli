package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	mqttBroker    = "tcp://iot.eclipse.org:1883"
	clientID      = "palinta-go-cli-client"
	mqttRelayPath = "palinta/relay"
)

func sendMqttMessage(path string, message string) {
	opts := mqtt.NewClientOptions().AddBroker(mqttBroker)
	opts.SetClientID(clientID)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	token := c.Publish(path, 0, false, message)
	token.Wait()

	c.Disconnect(250)
}
