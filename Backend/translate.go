package main

import (
	"context"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

func TranslateToJapanese(text string) (string, error) {
	ctx := context.Background()

	client, err := translate.NewClient(ctx, option.WithCredentialsFile("credentials.json"))
	if err != nil {
		return "", err
	}
	defer client.Close()

	target, _ := language.Parse("ja")

	resp, err := client.Translate(ctx, []string{text}, target, nil)
	if err != nil {
		return "", err
	}

	return resp[0].Text, nil
}
