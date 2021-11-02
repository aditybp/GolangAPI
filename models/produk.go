package models

type Produk struct {
	Id uint
	Nama_Barang string
	Kode_Barang string `gorm:"unique"`
	Harga_Barang string
}
