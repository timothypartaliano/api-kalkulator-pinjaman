1. go run main.go
2. Buka Postman
3. Pilih metode HTTP "POST"
4. Masukkan URL Endpoint API : http://localhost:8080/kalkulator_pinjaman
5. Isi body dengan format JSON :
{
    "plafon": 5000000,
    "lama_pinjaman": 6,
    "suku_bunga": 8,
    "tanggal_mulai": "2023-01-08"
}
6. Klik Send dan hasil angsuran akan muncul