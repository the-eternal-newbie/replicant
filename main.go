package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"replicant/replicant"
)

func main() {
	replicant.InitDB("data/memory.db")
	defer replicant.CloseDB()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to Replicant.\nChoose mode: [train / chat]")
	mode, _ := reader.ReadString('\n')
	mode = strings.TrimSpace(mode)

	if mode == "train" {
		replicant.Train()
		fmt.Println("Training complete. Restart in chat mode to talk with your Replicant.")
		return
	}

	fmt.Println("Start chatting with your Replicant. Type 'exit' to quit.")
	for {
		fmt.Print("You: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		prompt := replicant.BuildPrompt(input)
		response, err := replicant.QueryLLM(prompt)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println("Replicant:", response)
	}
}
