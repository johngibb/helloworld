package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("handleAPI", "path", r.RequestURI)
		w.Header().Set("Content-Type", "application/json")
		resp := map[string]any{
			"msg":      "Hello, world! Sincerely, us",
			"path":     r.RequestURI,
			"location": "NYC",
		}
		b, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			slog.Error("marshaling json", "error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		_, err = w.Write(b)
		if err != nil {
			slog.Error("writing response", "error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})
	slog.Info("listening", "port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		slog.Error("couldn't listen on port", "error", err)
	}
}
