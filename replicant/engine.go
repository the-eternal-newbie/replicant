package replicant

import (
	"fmt"
)

func BuildSystemPrompt() string {
	return fmt.Sprintf(
		"You are %s, a replicant who is a %s based in %s. You speak in a %s tone. " +
			"Your interests include %s. You enjoy sharing knowledge and personal anecdotes like: '%s'.",
		Get("traits", "name"),
		Get("traits", "skills"),
		Get("traits", "location"),
		Get("traits", "tone"),
		Get("traits", "interests"),
		Get("traits", "fun_fact"),
	)
}

func BuildPrompt(userInput string) string {
	system := BuildSystemPrompt()

	// Check if it matches stored Q&A
	answer := Get("knowledge", userInput)
	if answer != "" {
		return system + "\n\nUser: " + userInput + "\nAssistant: " + answer
	}

	return system + "\n\nUser: " + userInput + "\nAssistant:"
}
