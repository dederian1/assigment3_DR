package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

// Status struct merepresentasikan struktur JSON status.
type Status struct {
	Water  int    `json:"water"`
	Wind   int    `json:"wind"`
	Status string `json:"status"`
}

func main() {
	// Handle request untuk /update
	http.HandleFunc("/update", updateHandler)
	// Mulai server pada port 8080
	http.ListenAndServe(":8080", nil)
}

// updateHandler menangani permintaan update dari client.
func updateHandler(w http.ResponseWriter, r *http.Request) {
	// Generate nilai acak untuk air dan angin
	water := rand.Intn(100) + 1
	wind := rand.Intn(100) + 1

	// Tentukan status air dan angin
	waterStatus := determineWaterStatus(water)
	windStatus := determineWindStatus(wind)

	// Buat struktur status
	status := Status{
		Water:  water,
		Wind:   wind,
		Status: fmt.Sprintf("Air: %s, Angin: %s", waterStatus, windStatus),
	}

	// Marshal status ke JSON
	jsonData, err := json.Marshal(status)
	if err != nil {
		// Jika terjadi kesalahan, kirim respons error ke client
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Error marshalling JSON:", err) // Cetak error ke konsol
		return
	}

	// Tulis respons JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// determineWaterStatus menentukan status air berdasarkan nilainya.
func determineWaterStatus(water int) string {
	switch {
	case water < 5:
		return "Aman"
	case water >= 6 && water <= 8:
		return "Siaga"
	default:
		return "Bahaya"
	}
}

// determineWindStatus menentukan status angin berdasarkan nilainya.
func determineWindStatus(wind int) string {
	switch {
	case wind < 6:
		return "Aman"
	case wind >= 7 && wind <= 15:
		return "Siaga"
	default:
		return "Bahaya"
	}
}
