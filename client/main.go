package client

import (
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
)

func normalize(text string) string {
	return strings.TrimSpace(text)
}

func tokenize(text string) []string {
	return strings.Split(text, " ")
}

func executor(in string) {
	tokens := tokenize(normalize(in))
	switch tokens[0] {
	case "exit":
		os.Exit(0)
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

func Run() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("dqueue> "),
		prompt.OptionTitle("dqueue"),
	)
	p.Run()
}
