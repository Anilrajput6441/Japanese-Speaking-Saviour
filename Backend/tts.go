package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	texttospeechpb "cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"google.golang.org/api/option"
)

func GenerateAudio(text string) (string, error) {
	ctx := context.Background()

	// Initialize Google TTS client
	client, err := texttospeech.NewClient(ctx, option.WithCredentialsFile("credentials.json"))
	if err != nil {
		return "", err
	}
	defer client.Close()

	// Create the TTS request
	req := &texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: text},
		},
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "ja-JP",
			SsmlGender:   texttospeechpb.SsmlVoiceGender_NEUTRAL,
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	// Make the TTS request
	resp, err := client.SynthesizeSpeech(ctx, req)
	if err != nil {
		return "", err
	}

	// Save the audio file into ./static/ folder
	filename := fmt.Sprintf("audio_%d.mp3", time.Now().UnixNano())
	filePath := "./static/" + filename

	err = ioutil.WriteFile(filePath, resp.AudioContent, 0644)
	if err != nil {
		return "", err
	}

	// Return the relative URL for frontend to access
	return "/static/" + filename, nil
}
