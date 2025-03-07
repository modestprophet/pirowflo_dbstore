package main

import (
	"bufio"
	"fmt"
	"os"
)

func commandBacklog(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("missing filename argument")
	}

	file, err := os.Open(cmd.args[0])
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		msg, err := parseMessage(scanner.Bytes())
		if err != nil {
			fmt.Printf("error processing message: %v\n", err)
			continue
		}
		saveRowerData(s, msg)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("file scanning error: %w", err)
	}

	return nil
}
