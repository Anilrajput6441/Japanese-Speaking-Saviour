package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type TranslateRequest struct {
	Text string `json:"text"`
}

type TranslateResponse struct {
	Translated string `json:"translated"`
	AudioUrl   string `json:"audioUrl"`
}

func DeleteAudioFile(audioFile string) error {
	// Delete the audio file after a fixed time (1 hour for example)
	err := os.Remove(audioFile)
	if err != nil {
		log.Println("Error deleting audio file:", err)
		return err
	}
	log.Println("Audio file deleted:", audioFile)
	return nil
}

func main() {
	r := gin.Default()

	// Enable CORS middleware to allow cross-origin requests
	r.Use(cors.Default())

	// Serve static files (audio) from the /static folder
	r.Static("/static", "./static")

	// WebSocket endpoint
	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("WebSocket Upgrade failed:", err)
			return
		}
		defer conn.Close()

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("WebSocket Read failed:", err)
				break
			}

			var req TranslateRequest
			err = json.Unmarshal(msg, &req)
			if err != nil {
				log.Println("Error unmarshalling message:", err)
				continue
			}

			// Translate the text and generate the audio
			translatedText, _ := TranslateToJapanese(req.Text) // Implement this function as needed
			audioURL, _ := GenerateAudio(translatedText)

			// Respond with translation and audio URL
			response := TranslateResponse{
				Translated: translatedText,
				AudioUrl:   "http://localhost:8080" + audioURL, // Ensure correct URL is returned
			}

			// Send the response back
			conn.WriteJSON(response)

			// Delete the audio file after a fixed expiration time (e.g., 1 hour)
			go func() {
				// Wait for 1 hour (3600 seconds) before deleting all audio files in the static directory
				time.Sleep(1 * time.Hour)
				files, err := os.ReadDir("./static")
				if err != nil {
					log.Println("Error reading static directory:", err)
					return
				}
				for _, file := range files {
					if !file.IsDir() {
						err := os.Remove("./static/" + file.Name())
						if err != nil {
							log.Println("Error deleting file:", file.Name(), err)
						} else {
							log.Println("Deleted file:", file.Name())
						}
					}
				}
			}()
		}
	})

	log.Println("Server started at http://localhost:8080")
	r.Run(":8080")
}
