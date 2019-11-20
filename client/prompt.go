package client

import (
	"fmt"
	"os"
	"strings"

	"github.com/arpitbbhayani/dqueue/models"
	"github.com/c-bata/go-prompt"
)

func executor(in string) {
	tokens := tokenize(normalize(in))
	switch tokens[0] {
	case "exit":
		os.Exit(0)
	case "put":
		fmt.Println(dqclient.PutMessage(&models.PutMessageRequest{
			Message: strings.Join(tokens[1:], " "),
		}).ToString())
	}
}

func completer(d prompt.Document) []prompt.Suggest {
	commandSuggests := []prompt.Suggest{
		{Text: "put", Description: "puts a message to the queue"},
		{Text: "peek", Description: "peeks the message from the front of the queue"},
		{Text: "exit", Description: "exits the client"},
	}

	tokens := tokenize(normalize(d.Text))
	if len(tokens) == 0 {
		return commandSuggests
	}

	if len(tokens) == 1 && d.GetWordBeforeCursor() != "" {
		return prompt.FilterHasPrefix(commandSuggests, d.GetWordBeforeCursor(), true)
	}

	return []prompt.Suggest{}
}
