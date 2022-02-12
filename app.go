package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/wd30gsrc/bonus-time-server/customMiddleware"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("環境変数読み込み失敗: %v", err)
	}
}

type Message struct {
	SendMessage string `json:"message"`
}

func main() {
	loadEnv()

	public := http.HandlerFunc(publicHandler)
	private := http.HandlerFunc(privateHandler)
	http.Handle("/public", customMiddleware.Cors(public))
	http.Handle("/private", customMiddleware.Cors(customMiddleware.Auth(private)))

	http.ListenAndServe(":5000", nil)
}

func publicHandler(w http.ResponseWriter, r *http.Request) {
	message := &Message{
		SendMessage: "Public Message!",
	}
	data, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func privateHandler(w http.ResponseWriter, r *http.Request) {
	message := &Message{
		SendMessage: "Private Message!",
	}
	data, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// func initializeAppWithServiceAccount() *firebase.App {
// 	// [START initialize_app_service_account_golang]
// 	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
// 	app, err := firebase.NewApp(context.Background(), nil, opt)
// 	if err != nil {
// 		log.Fatalf("error initializing app: %v\n", err)
// 	}
// 	// [END initialize_app_service_account_golang]

// 	return app
// }
