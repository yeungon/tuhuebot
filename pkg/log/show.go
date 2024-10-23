package logging

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	tele "gopkg.in/telebot.v3"
)

func Show(c tele.Context, file string) {
	logs, err := ReadLastNLines(file, 20)
	if err != nil {
		fmt.Printf("Failed to read log file: %v\n", err)
		return
	}
	// Format the logs into a single message
	message := strings.Join(logs, "\n")
	if len(message) > 4096 { // Telegram's message limit
		message = message[:4096] + "\n...(truncated)"
	}

	c.Send(message)
}

// Function to read the last N lines from a file
func ReadLastNLines(filename string, n int) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}

	// Read the file line by line and store in a slice
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	// Get the last N lines (or all if there are fewer than N)
	if len(lines) > n {
		return lines[len(lines)-n:], nil
	}
	return lines, nil
}
