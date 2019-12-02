package client

import (
	"os"
	"strings"

	"github.com/arpitbbhayani/dqueue/models"
	"github.com/c-bata/go-prompt"
	"github.com/sirupsen/logrus"
)

func executor(in string) {
	tokens := tokenize(normalize(in))
	switch tokens[0] {
	case "exit":
		os.Exit(0)
	case "put":
		r, err := dqclient.PutMessage(&models.PutMessageRequest{
			Message: strings.Join(tokens[1:], " "),
		})
		if err != nil {
			logrus.Errorf("error: %s", err)
			return
		}
		logrus.Info(r.ToString())
	case "get":
		r, err := dqclient.GetMessage(&models.GetMessageRequest{})
		if err != nil {
			logrus.Errorf("error: %s", err)
			return
		}
		logrus.Info(r.ToString())
	}
}

func completer(d prompt.Document) []prompt.Suggest {
	commandSuggests := []prompt.Suggest{
		{Text: "put", Description: "puts a message to the queue"},
		{Text: "get", Description: "gets a message from the queue"},
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
