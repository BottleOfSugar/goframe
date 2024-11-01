package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const version string = "1.0"
const codename string = "Star Ink"

var lines_of_commands int = 0

func main() {
	reader := bufio.NewReader(os.Stdin) // Wczytywanie danych od użytkownika

	fmt.Printf("Welcome to %s Shell (version %s)\n", codename, version)

	for {
		fmt.Printf("%d.> ", lines_of_commands)
		lines_of_commands++

		// Wczytuje komendę od użytkownika
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) // Usuwa końcowe znaki nowej linii i spacje

		// Jeśli użytkownik wpisze "exit", kończy działanie powłoki
		if input == "exit" {
			fmt.Println("Exiting shell...")
			break
		}

		// Rozdziela komendę i jej argumenty
		args := strings.Split(input, " ")
		cmd := exec.Command(args[0], args[1:]...)

		// Ustawia wyjście i błąd na stdout i stderr
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Wykonuje komendę
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	}
}
