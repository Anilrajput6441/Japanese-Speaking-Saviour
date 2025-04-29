# 🇯🇵🌸 Arigatou Gozaimasu !

# 🇯🇵 Live Japanese Translator (Go + React)

A full-stack real-time Japanese translation and audio playback tool — built with **React**, **Go (Gin)**, and **Google Cloud Text-to-Speech**.  
This project was born out of my personal dream of working with Japanese technology and culture, and showcases real-time communication, frontend UX, and backend audio processing.

## 🔗 Demo



https://github.com/user-attachments/assets/b38fe8b1-e5db-4389-b02d-a09dd19d18a6



> ⚠️ This project is not yet hosted live. You can run it locally by following the instructions below.

---

## 🚀 Features

- **Real-time translation** of English to Japanese
- **Live audio generation** using Google Cloud TTS API
- **WebSocket-based** communication for low-latency streaming
- **Auto-deletion** of audio files after 1 hour (Go routine cleanup)
- **Responsive frontend** using React + TailwindCSS
- **Dynamic audio playback** using HTML5 `<audio>` tag

---

## 🛠 Tech Stack

| Layer      | Tech Used                         |
|------------|----------------------------------|
| Frontend   | React, Tailwind CSS, Vite        |
| Backend    | Go, Gin Framework, Gorilla WebSocket |
| Cloud API  | Google Cloud Text-to-Speech API  |
| Dev Tools  | UUID, GCP service credentials, Go routines |

---

## 📦 Project Structure
    /frontend # React App (Vite + Tailwind) 
    /backend ├── main.go # Gin server with WebSocket handler 
             ├── tts.go # TTS audio generation logic using GCP 
             ├── translate.go # Text translation logic (can be stubbed or extended) 
             └── static/ # Audio files generated temporarily



---

## 🧑‍💻 Setup Instructions

### 🔹 Backend (Go)

1. Make sure you have **Go** installed (v1.18+ recommended)
2. Create a Google Cloud Project & enable the **Text-to-Speech API**
3. Download your `credentials.json` from GCP and place it in `/backend/`
4. Run the backend:

```bash
backend(Go)
cd backend
go mod tidy
go run main.go tts.go translate.go

          Server will run at http://localhost:8080

 Frontend (React)
cd frontend
npm install
npm run dev

          App will be available at http://localhost:5173


Extras
Files are saved to /static/ and served via Gin

Audio auto-deletes after 1 hour to prevent bloat

WebSocket optimizations with debounced input from frontend

CORS properly configured for localhost:5173 ⇄ 8080


🧳 Author
Anil Behera
[Portfolio](https://github.com/Anilrajput6441)
[LinkedIn](https://www.linkedin.com/in/anil-behera-ai)






