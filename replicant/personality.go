package replicant

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var setupQuestions = map[string]string{
	"name":      "What's your name?",
	"location":  "Where are you from?",
	"skills":    "List your main skills (comma-separated):",
	"interests": "What are your interests?",
	"tone":      "What tone should I use when replying? (e.g., sarcastic, friendly, professional)",
	"fun_fact":  "Tell me a fun fact about yourself.",
}

func Train() {
	fmt.Println("Welcome to Replicant training mode.\nAnswer the following to define your bot's personality.")

	reader := bufio.NewReader(os.Stdin)
	for key, question := range setupQuestions {
		fmt.Print(question + " ")
		answer, _ := reader.ReadString('\n')
		Save("traits", key, strings.TrimSpace(answer))
	}

	// Add some known facts
	fmt.Println("\nAdd a few facts or common questions and answers your bot should know (type 'done' to finish):")
	for {
		fmt.Print("Q: ")
		question, _ := reader.ReadString('\n')
		if strings.TrimSpace(question) == "done" {
			break
		}
		fmt.Print("A: ")
		answer, _ := reader.ReadString('\n')
		Save("knowledge", strings.TrimSpace(question), strings.TrimSpace(answer))
	}
}
