package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/modestprophet/pirowflo_dbstore/internal/config"
	"github.com/modestprophet/pirowflo_dbstore/internal/database"
)

func startDataStorage(s *state, cmd command) error {
	mqttClient, err := setupMQTT(s.cfg)
	if err != nil {
		return fmt.Errorf("failed to setup MQTT: %v", err)
	}
	defer mqttClient.Disconnect()

	sigChan := setupSignalHandling()

	go processMessages(s, mqttClient.MessageCh)

	<-sigChan
	fmt.Println("Shutting down...")
	return nil
}

func setupMQTT(cfg *config.Config) (*MQTTClient, error) {
	client, err := NewMQTTClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create MQTT client: %v", err)
	}
	if err := client.Subscribe(); err != nil {
		return nil, fmt.Errorf("failed to subscribe: %v", err)
	}
	return client, nil
}

func setupSignalHandling() chan os.Signal {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	return sigChan
}

func processMessages(s *state, messageCh <-chan []byte) {
	for payload := range messageCh {
		msg, err := parseMessage(payload)
		if err != nil {
			continue
		}
		saveRowerData(s, msg)
	}
}

func parseMessage(payload []byte) (rowerMessage, error) {
	var msg rowerMessage
	if err := json.Unmarshal(payload, &msg); err != nil {
		return rowerMessage{}, fmt.Errorf("error unmarshaling message: %w", err)
	}
	return msg, nil
}

func saveRowerData(s *state, msg rowerMessage) {
	id := uuid.New()
	now := time.Now()
	rower_data := database.SaveRowerDataParams{
		ID:                id,
		CreatedAt:         now,
		StrokeRate:        int32(msg.StrokeRate),
		TotalStrokes:      int32(msg.TotalStrokes),
		TotalDistanceM:    int32(msg.TotalDistanceM),
		InstantaneousPace: float32(msg.InstantaneousPace),
		Speed:             int32(msg.Speed),
		Watts:             int32(msg.Watts),
		TotalKcal:         float32(msg.TotalKcal),
		TotalKcalHour:     int32(msg.TotalKcalHour),
		TotalKcalMin:      int32(msg.TotalKcalMin),
		HeartRate:         int32(msg.HeartRate),
		Elapsedtime:       int32(msg.Elapsedtime),
	}

	if _, err := s.db.SaveRowerData(context.Background(), rower_data); err != nil {
		fmt.Println("Error writing message to database: %w", err)
	}
}

// func startDataStorageOld(s *state, cmd command) error {
// 	// Initialize MQTT Client
// 	mqttClient, err := NewMQTTClient(s.cfg)
// 	if err != nil {
// 		log.Fatalf("Failed to create MQTT client: %v", err)
// 	}
// 	defer mqttClient.Disconnect()

// 	// Start MQTT subscription
// 	if err := mqttClient.Subscribe(); err != nil {
// 		log.Fatalf("Failed to subscribe: %v", err)
// 	}

// 	// Set up signal handling for graceful shutdown
// 	sigChan := make(chan os.Signal, 1)
// 	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

// 	// Process messages
// 	go func() {
// 		for payload := range mqttClient.MessageCh {
// 			var msg rowerMessage
// 			if err := json.Unmarshal(payload, &msg); err != nil {
// 				log.Printf("Error unmarshaling message: %v", err)
// 				continue
// 			}

// 			id := uuid.New()
// 			now := time.Now()
// 			rower_data := database.SaveRowerDataParams{
// 				ID:                id,
// 				CreatedAt:         now,
// 				StrokeRate:        int32(msg.StrokeRate),
// 				TotalStrokes:      int32(msg.TotalStrokes),
// 				TotalDistanceM:    int32(msg.TotalDistanceM),
// 				InstantaneousPace: float32(msg.InstantaneousPace),
// 				Speed:             int32(msg.Speed),
// 				Watts:             int32(msg.Watts),
// 				TotalKcal:         float32(msg.TotalKcal),
// 				TotalKcalHour:     int32(msg.TotalKcalHour),
// 				TotalKcalMin:      int32(msg.TotalKcalMin),
// 				HeartRate:         int32(msg.HeartRate),
// 				Elapsedtime:       int32(msg.Elapsedtime),
// 			}

// 			if _, err := s.db.SaveRowerData(context.Background(), rower_data); err != nil {
// 				log.Printf("Error writing message to database: %v", err)
// 			}
// 		}
// 	}()

// 	// Wait for shutdown signal
// 	<-sigChan
// 	log.Println("Shutting down...")
// 	return nil
// }
