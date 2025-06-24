package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, map[string]string{
			"status":  "ok",
			"service": "1",
		})
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("➡️  /hello hit from %s", r.RemoteAddr)
		jsonResponse(w, map[string]string{
			"message": "Hello from Service 1",
		})
	})

	log.Println("Service 1 listening on port 8001...")
	if err := http.ListenAndServe(":8001", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// jsonResponse writes the given data as JSON to the http.ResponseWriter.
func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

