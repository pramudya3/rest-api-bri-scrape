package model

type GetSaldo struct {
	NomorRekening string `json:"nomor_rekening"`
	JenisProduk   string `json:"jenis_produk"`
	Nama          string `json:"nama"`
	MataUang      string `json:"mata_uang"`
	Saldo         string `json:"saldo"`
}
