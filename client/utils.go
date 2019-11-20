package client

import "strings"

func normalize(text string) string {
	return strings.TrimSpace(text)
}

func tokenize(text string) []string {
	return strings.Split(text, " ")
}
