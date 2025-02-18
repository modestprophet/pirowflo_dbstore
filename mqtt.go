package main

import (
	"fmt"
	"log"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/modestprophet/pirowflo_dbstore/internal/config"
)

type rowerMessage struct {
	StrokeRate        int       `json:"stroke_rate"`
	TotalStrokes      int       `json:"total_strokes"`
	TotalDistanceM    int       `json:"total_distance_m"`
	InstantaneousPace float64   `json:"instantaneous pace"`
	Speed             int       `json:"speed"`
	Watts             int       `json:"watts"`
	TotalKcal         float64   `json:"total_kcal"`
	TotalKcalHour     int       `json:"total_kcal_hour"`
	TotalKcalMin      int       `json:"total_kcal_min"`
	HeartRate         int       `json:"heart_rate"`
	Elapsedtime       int       `json:"elapsedtime"`
	Timestamp         time.Time `json:"timestamp"`
	WorkoutID         string    `json:"workout_id"`
}

type MQTTClient struct {
	client    MQTT.Client
	MessageCh chan []byte
	topic     string
}

func NewMQTTClient(cfg *config.Config) (*MQTTClient, error) {
	// Set MQTT client paramters
	opts := MQTT.NewClientOptions()
	opts.AddBroker(cfg.MqServerURL)
	opts.SetClientID(cfg.MqClientID)
	opts.SetUsername(cfg.MqUser)
	opts.SetPassword(cfg.MqPassword)

	// Set up reconnection logic
	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(time.Second * 30)

	// Set up connection lost handler
	opts.SetConnectionLostHandler(func(client MQTT.Client, err error) {
		log.Printf("Connection lost: %v", err)
	})

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, fmt.Errorf("error connecting to MQTT broker: %v", token.Error())
	}

	return &MQTTClient{
		client:    client,
		MessageCh: make(chan []byte, 100), // Buffered channel
		topic:     cfg.MqTopic,
	}, nil
}

func (m *MQTTClient) Subscribe() error {
	handler := func(client MQTT.Client, msg MQTT.Message) {
		m.MessageCh <- msg.Payload()
		msg.Ack()
	}

	if token := m.client.Subscribe(m.topic, 1, handler); token.Wait() && token.Error() != nil {
		return fmt.Errorf("error subscribing to topic: %v", token.Error())
	}

	log.Printf("Subscribed to topic: %s", m.topic)
	return nil
}

func (m *MQTTClient) Disconnect() {
	m.client.Disconnect(250)
	close(m.MessageCh)
}
