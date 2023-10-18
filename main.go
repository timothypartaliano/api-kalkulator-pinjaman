package main

import (
	"api/handlers"
	"net/http"
)

func main() {
    http.HandleFunc("/kalkulator_pinjaman", handlers.KalkulatorPinjaman)
    http.ListenAndServe(":8080", nil)
}