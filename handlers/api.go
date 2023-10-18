package handlers

import (
	"api/model"
	"encoding/json"
	"net/http"
	"time"
)

func hitungAngsuran(plafon float64, sukuBunga float64, lamaPinjaman int, tanggalMulai time.Time) []model.Angsuran {
	var angsuran []model.Angsuran
	sukuBungaBulanan := (sukuBunga / 100) / 12
	totalAngsuran := (plafon * sukuBungaBulanan) / (1 - (1 / (1 + sukuBungaBulanan)))
	sisaAngsuran := plafon

	for i := 1; i <= lamaPinjaman; i++ {
		angBunga := (sukuBunga / 360) * 30 * sisaAngsuran
		angPokok := totalAngsuran - angBunga

		ang := model.Angsuran{
			AngsuranKe:        i,
			Tanggal:           tanggalMulai.Format("2006-01-02"),
			TotalAngsuran:     totalAngsuran,
			AngsuranPokok:     angPokok,
			AngsuranBunga:     angBunga,
			SisaAngsuranPokok: sisaAngsuran - angPokok,
		}

		angsuran = append(angsuran, ang)

		tanggalMulai = tanggalMulai.AddDate(0, 1, 0)
		sisaAngsuran -= angPokok
	}

	return angsuran
}

func KalkulatorPinjaman(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Plafon       float64 `json:"plafon"`
		LamaPinjaman int     `json:"lama_pinjaman"`
		SukuBunga    float64 `json:"suku_bunga"`
		TanggalMulai string  `json:"tanggal_mulai"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	tanggalMulai, err := time.Parse("2006-01-02", input.TanggalMulai)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	angsuran := hitungAngsuran(input.Plafon, input.SukuBunga, input.LamaPinjaman, tanggalMulai)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(angsuran)
}
