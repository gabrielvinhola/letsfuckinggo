package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadUserInput() []string {
	var userInputs []string

	fmt.Println("Enter countries (one per line, empty line to stop):")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			break
		}
		userInputs = append(userInputs, input)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading standard input:", err)
	}

	return userInputs
}
