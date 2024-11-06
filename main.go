package main

import (
	"bufio"
	"fmt"
	"os"
	"parking-app/handlers"
	"strconv"
	"strings"
)

func main() {

	// Process for input file .txt commands
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command file.")
		return
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		processCommand(line)
	}
}

func processCommand(line string) {
	args := strings.Fields(line)
	if len(args) == 0 {
		return
	}

	command := args[0]
	// Checking for valid commands
	switch command {
	case "create_parking_lot":
		capacity, _ := strconv.Atoi(args[1])
		handlers.CreateParkingLot(capacity)
	case "park":
		carNumber := args[1]
		handlers.Park(carNumber)
	case "leave":
		carNumber := args[1]
		hours, _ := strconv.Atoi(args[2])
		handlers.Leave(carNumber, hours)
	case "status":
		handlers.Status()
	default:
		fmt.Println("Invalid command:", command)
	}
}
