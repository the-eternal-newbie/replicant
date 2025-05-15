package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"replicant/replicant"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No .env file found or couldn't load it")
	}

	dbPath := os.Getenv("REPLICANT_DB")
	if dbPath == "" {
		dbPath = "data/memory.db"
	}

	replicant.InitDB(dbPath)
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
