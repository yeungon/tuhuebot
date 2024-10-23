package logging

import (
	"bufio"
	"fmt"
	"os"
)

func Read() {
	file, err := os.Open("tuhuebot.json")
	if err != nil {
		fmt.Printf("Failed to open log file: %v\n", err)
		return
	}
	defer file.Close()

	fmt.Println("[")
	scanner := bufio.NewScanner(file)
	first := true
	for scanner.Scan() {
		if !first {
			fmt.Println(",")
		}
		fmt.Print(scanner.Text())
		first = false
	}
	fmt.Println("\n]")

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
}
